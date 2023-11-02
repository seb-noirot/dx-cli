package gitlabstack

import (
	"dx-cli/config" // Replace with the actual import path
	"dx-cli/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var listStacksCmd = &cobra.Command{
	Use:   "list",
	Short: "📋👀 Lists all GitLab stacks for the current context",
	Long:  "📋🔍 Use this command to display all the GitLab stacks defined in your current context. Ideal for a quick overview and management. 🕵️‍♂️👌",
	Run: func(cmd *cobra.Command, args []string) {
		path, err := config.GetConfigFilePath()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		currentContext, err := utils.GetCurrentContext(path, false)
		if err != nil {
			utils.LogError("❌ Error fetching current context: %s", err)
			return
		}

		if currentContext == nil {
			utils.LogInfo("⚠️ No current context defined.")
			return
		}

		if len(currentContext.GitLabContexts) == 0 {
			utils.LogInfo("⚠️ No GitLab definitions available.")
			return
		}

		utils.LogInfo("👇 Available GitLab Stacks:")
		for _, gitlabContext := range currentContext.GitLabContexts {
			utils.LogInfo(fmt.Sprintf("🔹 GitLab Context: %s", gitlabContext.Name))
			for _, stack := range gitlabContext.GitlabStacks {
				utils.LogInfo(fmt.Sprintf("  📦 Stack: %s", stack.Name))
				utils.LogInfo(fmt.Sprintf("  📍 Path: %s", stack.Path))

				projectList := "  📚 Projects:"
				for _, project := range stack.Projects {
					projectList += fmt.Sprintf("\n    • %s", project)
				}
				utils.LogInfo(projectList)
			}
		}

	},
}

func init() {
	StackCmd.AddCommand(listStacksCmd)
}
