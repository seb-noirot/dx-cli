package gitlabstack

import (
	"github.com/spf13/cobra"
)

var StackCmd = &cobra.Command{
	Use:   "stack",
	Short: "📚 Manage gitlabstack from GitLab",
	Long:  "📚 The `stack` command allows you to manage your GitLab stack configurations directly from the GitLab repository. Perfect for keeping your stack setup in sync and organized. 🔄",
}
