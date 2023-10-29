package k8s

import (
	"github.com/spf13/cobra"
)

var KubeCmd = &cobra.Command{
	Use:   "k8s",
	Short: "🚀 Elevate Your K8s Game!",
	Long:  `🎉 Unleash the full potential of Kubernetes with a single CLI. From resource management to cluster orchestration, make everything a breeze.`,
}
