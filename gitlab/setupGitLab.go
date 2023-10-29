package gitlab

import (
	"dx-cli/utils"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"dx-cli/config"
	"github.com/spf13/cobra"
)

// setupGitLabCmd represents the setup command
var setupGitLabCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup a GitLab definition",
	Run: func(cmd *cobra.Command, args []string) {
		// Fetch the current context
		currentContext, err := utils.GetCurrentContext(config.ConfigFilePath, false)
		if err != nil {
			fmt.Printf("Error fetching current context: %s\n", err)
			return
		}

		if currentContext == nil {
			fmt.Println("No current context defined.")
			return
		}

		if len(currentContext.GitLabContexts) == 0 {
			fmt.Println("No GitLab definitions available.")
			return
		}

		fmt.Println("Select a GitLab definition:")
		for i, glContext := range currentContext.GitLabContexts {
			fmt.Printf("[%d] %s (%s)\n", i+1, glContext.Name, glContext.Host)
		}

		var choice int
		fmt.Scanln(&choice)

		if choice < 1 || choice > len(currentContext.GitLabContexts) {
			fmt.Println("Invalid choice.")
			return
		}

		selectedGitLab := currentContext.GitLabContexts[choice-1]

		keyName := fmt.Sprintf("id_%s_%s", currentContext.Name, selectedGitLab.Name)
		// Create SSH Key
		createSSHKeys(keyName)

		// Update known_hosts
		setupSSHConfig(selectedGitLab, keyName)
	},
}

func createSSHKeys(keyName string) error {
	privateKeyPath := filepath.Join(os.Getenv("HOME"), ".ssh", keyName)
	if _, err := os.Stat(privateKeyPath); err == nil {
		fmt.Println("Key already exists. Skipping key generation.")
		return nil
	} else if os.IsNotExist(err) {
		cmd := exec.Command("ssh-keygen", "-t", "rsa", "-b", "4096", "-f", privateKeyPath, "-N", "")
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("could not create SSH keys: %s", err)
		}
		fmt.Println("SSH keys created successfully.")
	} else {
		return fmt.Errorf("could not check for existing SSH key: %s", err)
	}

	return nil
}

func setupSSHConfig(selectedGitLab config.GitLabContext, privateKeyPath string) error {
	// Define the SSH config file path; you may want to make this configurable
	sshConfigPath := filepath.Join(os.Getenv("HOME"), ".ssh", "config")

	cleanHost := strings.TrimSuffix(strings.TrimPrefix(selectedGitLab.Host, "https://"), "/")

	// Create the configuration string to be appended
	configString := fmt.Sprintf("\nHost %s\n  PreferredAuthentications publickey\n  IdentityFile %s\n",
		cleanHost,
		privateKeyPath)

	// Open the file in append mode, creating it if it doesn't exist
	f, err := os.OpenFile(sshConfigPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("could not open SSH config file: %s", err)
	}
	defer f.Close()

	// Append the configuration string
	if _, err = f.WriteString(configString); err != nil {
		return fmt.Errorf("could not write to SSH config file: %s", err)
	}

	fmt.Println("SSH configuration completed successfully.")
	return nil
}

func init() {
	GitlabCmd.AddCommand(setupGitLabCmd)
}
