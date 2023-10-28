package k8s

import (
	"github.com/spf13/cobra"
)

var KubeCmd = &cobra.Command{
	Use:   "k8s",
	Short: "Manage k8s",
}
