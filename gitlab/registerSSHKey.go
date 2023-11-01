package gitlab

import (
	"dx-cli/config" // replace with your actual config package path
	"dx-cli/utils"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var registerSSHKeyCmd = &cobra.Command{
	Use:   "register",
	Short: "Register an SSH key with a GitLab account",
	RunE: func(cmd *cobra.Command, args []string) error {
		currentContext, err := utils.GetCurrentContext(config.ConfigFilePath, false)
		if err != nil {
			utils.LogError("Error fetching current context: %s", err)
			return nil
		}

		// Validate current context
		if currentContext == nil {
			utils.LogWarning("No current context defined.")
			return nil
		}

		if len(currentContext.GitLabContexts) == 0 {
			utils.LogWarning("No GitLab definitions available.")
			return nil
		}

		// Select GitLab instance
		selectedContext, err := utils.SelectGitlabDefinition()
		if err != nil {
			utils.LogError("Error fetching gitlab context: %s", err)
			return nil
		}

		// Validate current context
		if selectedContext == nil {
			utils.LogWarning("No gitlab context defined.")
			return nil
		}

		// Check if key exists
		privateKeyPath := filepath.Join(os.Getenv("HOME"), ".ssh", fmt.Sprintf("id_%s_%s", currentContext.Name, selectedContext.Name))
		if _, err := os.Stat(privateKeyPath); os.IsNotExist(err) {
			utils.LogError("SSH key does not exist. Please create one first.", err)
			return err
		}

		// Open GitLab SSH key page
		cleanHost := strings.TrimRight(selectedContext.Host, "/")
		gitLabURL := fmt.Sprintf("%s/-/profile/keys", cleanHost)
		err = exec.Command("open", gitLabURL).Run()
		if err != nil {
			utils.LogError("Could not open GitLab URL: %s", err)
			return err
		}

		// Display public key content to paste
		publicKeyPath := privateKeyPath + ".pub"
		publicKeyContent, err := os.ReadFile(publicKeyPath)
		if err != nil {
			utils.LogError("Could not read public key: %s", err)
			return err
		}
		clipboard.WriteAll(string(publicKeyContent))
		utils.LogInfo("Paste the following public key into GitLab:")
		fmt.Println(string(publicKeyContent))

		return nil
	},
}

func init() {
	GitlabCmd.AddCommand(registerSSHKeyCmd)
}
