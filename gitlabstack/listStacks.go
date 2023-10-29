package gitlabstack

import (
	"dx-cli/config" // Replace with the actual import path
	"dx-cli/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var listStacksCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all GitLab stacks for the current context",
	Run: func(cmd *cobra.Command, args []string) {
		currentContext, _ := utils.GetCurrentContext(config.ConfigFilePath, false)
		for _, gitlabContext := range currentContext.GitLabContexts {
			fmt.Println("GitLab Context:", gitlabContext.Name)
			for _, stack := range gitlabContext.GitlabStacks {
				fmt.Println("  Stack:", stack.Name)
				fmt.Println("  Path:", stack.Path)
				fmt.Println("  Projects:", stack.Projects)
			}
		}
	},
}

func init() {
	StackCmd.AddCommand(listStacksCmd)
}
