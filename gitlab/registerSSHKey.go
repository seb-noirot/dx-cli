package gitlab

import (
	"dx-cli/config"
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

var testGitLabConnectionCmd = &cobra.Command{
	Use:   "test",
	Short: "Test SSH connection to GitLab",
	Run: func(cmd *cobra.Command, args []string) {
		// Fetch the current context
		currentContext, err := config.GetCurrentContext()
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
		cleanHost := strings.TrimSuffix(strings.TrimPrefix(selectedGitLab.Host, "https://"), "/")
		TestSSHConnection(cleanHost)
	},
}

func TestSSHConnection(host string) error {
	cmd := exec.Command("ssh", "-T", fmt.Sprintf("git@%s", host))
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("SSH connection failed: %s", err)
	}

	fmt.Println("SSH connection successful. Output:", string(output))
	return nil
}

func init() {
	// Assuming you have a GitLab root command like 'gitlabCmd'
	GitlabCmd.AddCommand(testGitLabConnectionCmd)
}
