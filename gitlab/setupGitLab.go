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
	Short: "Initialize and configure a specific GitLab instance",
	Long:  "This command sets up a selected GitLab instance by generating necessary SSH keys and updating SSH configurations.",
	Run:   runSetupGitLab,
}

func runSetupGitLab(cmd *cobra.Command, args []string) {
	// Fetch current context
	currentContext, err := utils.GetCurrentContext(config.ConfigFilePath, false)
	if err != nil {
		utils.LogError("Error fetching current context: %s", err)
		return
	}

	// Validate current context
	if currentContext == nil {
		utils.LogWarning("No current context defined.")
		return
	}

	if len(currentContext.GitLabContexts) == 0 {
		utils.LogWarning("No GitLab definitions available.")
		return
	}

	// Select GitLab instance
	selectedContext, err := utils.SelectGitlabDefinition()
	if err != nil {
		utils.LogError("Error fetching gitlab context: %s", err)
		return
	}

	// Validate current context
	if selectedContext == nil {
		utils.LogWarning("No gitlab context defined.")
		return
	}

	// Generate and manage SSH keys
	keyName := fmt.Sprintf("id_%s_%s", currentContext.Name, selectedContext.Name)
	err = createSSHKeys(keyName)
	if err != nil {
		utils.LogError("SSH Key creation failed: %s", err)
		return
	}

	// Update SSH config
	err = setupSSHConfig(selectedContext, keyName)
	if err != nil {
		utils.LogError("SSH config setup failed: %s", err)
		return
	}

	utils.LogInfo("GitLab setup completed.")
}

func createSSHKeys(keyName string) error {
	privateKeyPath := filepath.Join(os.Getenv("HOME"), ".ssh", keyName)
	_, err := os.Stat(privateKeyPath)
	if err == nil {
		utils.LogInfo("SSH keys already exist. Skipping key generation.")
		return nil
	}

	if os.IsNotExist(err) {
		cmd := exec.Command("ssh-keygen", "-t", "rsa", "-b", "4096", "-f", privateKeyPath, "-N", "")
		err = cmd.Run()
		if err != nil {
			return fmt.Errorf("SSH key generation failed: %s", err)
		}
		utils.LogInfo("SSH keys created.")
	} else {
		return fmt.Errorf("Existing SSH key check failed: %s", err)
	}
	return nil
}

func setupSSHConfig(selectedGitLab *config.GitLabContext, privateKeyPath string) error {
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
