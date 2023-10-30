package k8s

import (
	"dx-cli/utils"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var setupToolsCmd = &cobra.Command{
	Use:   "setup",
	Short: "🛠️ Gear Up Your K8s Environment!",
	Long:  `🎉 This command installs the essential Kubernetes tools on macOS, making your Kubernetes experience hassle-free.`,
	Run: func(cmd *cobra.Command, args []string) {
		installTools()
	},
}

func init() {
	KubeCmd.AddCommand(setupToolsCmd)
}

func installTool(tool string) error {
	utils.Printf(true, "🔨 Installing %s...\n", tool)
	installCmd := exec.Command("brew", "install", tool)
	if err := installCmd.Run(); err != nil {
		return fmt.Errorf("failed to install %s: %s", tool, err)
	}
	utils.Printf(true, "🎉 Successfully installed %s.\n", tool)
	return nil
}

func installTools() {
	tools := []string{
		"Azure/kubelogin/kubelogin",
		"k9s",
		"derailed/k9s/k9s",
		"kubectl",
		"kubectx",
		"kubernetes-cli",
	}

	for _, tool := range tools {
		utils.Printf(true, "🔍 Checking %s...\n", tool)

		// Check if tool is already installed
		checkCmd := exec.Command("brew", "list", "--formula", tool)
		if err := checkCmd.Run(); err == nil {
			utils.Printf(true, "✅ %s is already installed\n", tool)
			continue
		}

		if utils.PromptUser(fmt.Sprintf("🛠️ Install %s?", tool), []string{"Yes", "No"}) == "Yes" {
			if err := installTool(tool); err != nil {
				utils.Printf(true, "🚨 %s\n", err.Error())
				utils.Printf(true, "🚀 Please run it manually: brew install %s\n", tool)
			}
		} else {
			utils.Printf(true, "🛑 Skipping installation of %s.\n", tool)
		}
	}
}
