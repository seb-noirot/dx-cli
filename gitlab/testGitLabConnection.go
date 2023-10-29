package gitlab

import (
	"dx-cli/utils"
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

var testGitLabConnectionCmd = &cobra.Command{
	Use:   "test",
	Short: "Test SSH connection to GitLab",
	Run: func(cmd *cobra.Command, args []string) {
		selectedGitLab, err := utils.SelectGitlabDefinition()
		if err != nil {
			utils.LogError("Error fetching GitLab context", err)
			return
		}

		if selectedGitLab == nil {
			utils.LogWarning("No valid GitLab definition selected.")
			return
		}

		cleanHost := strings.TrimSuffix(strings.TrimPrefix(selectedGitLab.Host, "https://"), "/")
		err = TestSSHConnection(cleanHost)
		if err != nil {
			utils.LogError("SSH connection test failed", err)
		} else {
			utils.LogInfo("SSH connection test successful.")
		}
	},
}

func TestSSHConnection(host string) error {
	cmd := exec.Command("ssh", "-T", fmt.Sprintf("git@%s", host))
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("SSH connection failed: %s", string(output))
	}

	utils.LogInfo(fmt.Sprintf("SSH connection successful. Output: %s", string(output)))
	return nil
}

func init() {
	// Assuming you have a GitLab root command like 'gitlabCmd'
	GitlabCmd.AddCommand(testGitLabConnectionCmd)
}
