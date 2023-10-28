package k8s

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var setupToolsCmd = &cobra.Command{
	Use:   "setup",
	Short: "Install essential Kubernetes tools",
	Long:  `This command will install essential Kubernetes tools on macOS.`,
	Run: func(cmd *cobra.Command, args []string) {
		installTools()
	},
}

func init() {
	KubeCmd.AddCommand(setupToolsCmd)
}

func installTools() {
	tools := []string{
		"Azure/kubelogin/kubelogin",
		"k9s",
		"kubectl",
		"kubectx",
		"kubernetes-cli",
	}

	for _, tool := range tools {
		// Check if tool is already installed
		checkCmd := exec.Command("brew", "list", "--formula", tool)
		if err := checkCmd.Run(); err == nil {
			fmt.Printf("%s is already installed, skipping.\n", tool)
			continue
		}

		// Attempt to install the tool
		fmt.Printf("Installing %s...\n", tool)
		installCmd := exec.Command("brew", "install", tool)
		if err := installCmd.Run(); err != nil {
			fmt.Printf("Failed to install %s: %s\n", tool, err)
			fmt.Printf("Please run it manually: brew install %s\n", tool)
		} else {
			fmt.Printf("Successfully installed %s.\n", tool)
		}
	}
}
