package gitlab

import (
	"dx-cli/config"
	"dx-cli/utils"
	"fmt"
	"github.com/spf13/cobra"
)

// addGitLabCmd represents the add command
var addGitLabCmd = &cobra.Command{
	Use:   "add",
	Short: "➕ Add GitLab Definition",
	Long:  "📚 Add a new GitLab definition to the current context by specifying its name and host.",
	Run: func(cmd *cobra.Command, args []string) {
		// Fetch the current context
		currentContext, err := utils.GetCurrentContext(config.ConfigFilePath, false)
		if err != nil {
			utils.LogError("Fetching current context", err)
			return
		}

		if currentContext == nil {
			utils.LogWarning("No current context defined.")
			return
		}

		// Prompt for name and host
		name := utils.PromptUser("Enter the name of the GitLab definition", nil)
		host := utils.PromptUser("Enter the host of the GitLab definition", nil)

		// Validate inputs
		if name == "" || host == "" {
			utils.LogWarning("Name and host cannot be empty.")
			return
		}

		// Create new GitLabContext and append it
		newGitLabContext := config.GitLabContext{Name: name, Host: host}
		currentContext.GitLabContexts = append(currentContext.GitLabContexts, newGitLabContext)

		// Update the current context
		err = config.UpdateCurrentContext(currentContext)
		if err != nil {
			utils.LogError("Updating current context", err)
			return
		}

		utils.LogInfo(fmt.Sprintf("New GitLab definition '%s' added.", name))
	},
}

func init() {
	GitlabCmd.AddCommand(addGitLabCmd)
}
