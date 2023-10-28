package maven

import (
	"github.com/spf13/cobra"
)

var MavenCmd = &cobra.Command{
	Use:   "mvn",
	Short: "Manage maven",
}
