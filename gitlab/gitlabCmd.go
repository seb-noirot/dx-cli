package gitlab

import (
	"dx-cli/gitlabstack"
	"github.com/spf13/cobra"
)

var GitlabCmd = &cobra.Command{
	Use:   "gitlab",
	Short: "Manage gitlab",
}

func init() {
	GitlabCmd.AddCommand(gitlabstack.StackCmd)
}
