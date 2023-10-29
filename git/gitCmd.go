package git

import (
	"github.com/spf13/cobra"
)

var GitCmd = &cobra.Command{
	Use:   "git",
	Short: "Manage Git operations",
	Long:  "The 'git' command provides a set of sub-commands to manage various Git operations such as testing SSH connections, registering SSH keys, and more.",
}
