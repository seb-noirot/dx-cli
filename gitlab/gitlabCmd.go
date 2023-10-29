package gitlab

import (
	"dx-cli/gitlabstack"
	"github.com/spf13/cobra"
)

var GitlabCmd = &cobra.Command{
	Use:   "gitlab",
	Short: "🦊 Manage GitLab Resources",
	Long:  `🎉 This command provides a suite of sub-commands to interact with and manage various GitLab resources such as repositories, pipelines, and more.`,
}

func init() {
	GitlabCmd.AddCommand(gitlabstack.StackCmd)
}
