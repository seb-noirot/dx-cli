package gitlab

import (
	"dx-cli/config"
	"fmt"

	"github.com/spf13/cobra"
)

// listGitLabCmd represents the list command
var listGitLabCmd = &cobra.Command{
	Use:   "list",
	Short: "List all GitLab definitions",
	Run: func(cmd *cobra.Command, args []string) {
		// Fetch the current context
		currentContext, err := config.GetCurrentContext()
		if err != nil {
			fmt.Printf("Error fetching current context: %s\n", err)
			return
		}

		// Handle null or empty cases
		if currentContext == nil {
			fmt.Println("No current context defined.")
			return
		}

		if len(currentContext.GitLabContexts) == 0 {
			fmt.Println("No GitLab definitions in the current context.")
			return
		}

		// List the GitLab definitions
		fmt.Println("GitLab definitions in current context:")
		for i, gitlabContext := range currentContext.GitLabContexts {
			fmt.Printf("%d. Name: %s, Host: %s\n", i+1, gitlabContext.Name, gitlabContext.Host)
		}
	},
}

// This function adds the listGitLabCmd to the rootCmd
func init() {
	GitlabCmd.AddCommand(listGitLabCmd)
}
