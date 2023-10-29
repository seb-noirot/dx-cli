package gitlab

import (
	"dx-cli/config"
	"dx-cli/utils"
	"github.com/spf13/cobra"
)

// listGitLabCmd represents the list command
var listGitLabCmd = &cobra.Command{
	Use:   "list",
	Short: "📝 List GitLab Definitions",
	Long:  "📚 Enumerate all available GitLab definitions within the current context, showing the name and host for each.",
	Run: func(cmd *cobra.Command, args []string) {
		// Fetch the current context
		currentContext, err := utils.GetCurrentContext(config.ConfigFilePath, false)
		if err != nil {
			utils.Printf(true, "🚨 Error: Could not fetch current context: %s\n", err)
			return
		}

		// Handle null or empty cases
		if currentContext == nil {
			utils.Println(true, "⚠️  Warning: No current context defined.")
			return
		}

		if len(currentContext.GitLabContexts) == 0 {
			utils.Println(true, "⚠️  Warning: No GitLab definitions found in the current context.")
			return
		}

		// List the GitLab definitions
		utils.Printf(true, "📋 GitLab Definitions in Current Context: %s\n", currentContext.Name)
		for i, gitlabContext := range currentContext.GitLabContexts {
			utils.Printf(true, "🔹 %d. Name: %s, Host: %s\n", i+1, gitlabContext.Name, gitlabContext.Host)
		}
	},
}

// This function adds the listGitLabCmd to the rootCmd
func init() {
	GitlabCmd.AddCommand(listGitLabCmd)
}
