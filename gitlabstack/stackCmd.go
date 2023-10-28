package gitlabstack

import (
	"github.com/spf13/cobra"
)

var StackCmd = &cobra.Command{
	Use:   "stack",
	Short: "Manage gitlabstack from gitlab",
}
