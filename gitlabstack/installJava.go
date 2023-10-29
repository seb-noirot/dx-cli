package gitlabstack

import (
	"dx-cli/utils"
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

var installJavaCmd = &cobra.Command{
	Use:   "java",
	Short: "Install java for the stack",
	RunE: func(cmd *cobra.Command, args []string) error {
		selectedGitLab, err := utils.SelectGitlabDefinition()
		if err != nil {
			utils.LogInfo("Error selecting GitLab definition.")
			return err
		}
		if selectedGitLab == nil {
			utils.LogInfo("No GitLab definition selected.")
			return nil
		}

		if len(selectedGitLab.GitlabStacks) == 0 {
			utils.LogInfo("⚠️  No stacks defined.")
			return nil
		}

		selectedStack, err := utils.SelectGitlabStack(selectedGitLab)
		if err != nil {
			return err
		}
		if selectedStack == nil {
			return nil
		}

		if len(selectedStack.Javas) == 0 {
			utils.LogInfo("⚠️  No Java versions defined for this stack.")
			return nil
		}

		for _, javaVersion := range selectedStack.Javas {
			utils.LogInfo(fmt.Sprintf("☕ Installing Java version %s using SDKMAN...", javaVersion))

			cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("source $HOME/.sdkman/bin/sdkman-init.sh && sdk install java %s", javaVersion))
			err = cmd.Run()
			if err != nil {
				utils.LogInfo(fmt.Sprintf("❌ Failed to install Java version %s: %s", javaVersion, err))
				return err
			}

			utils.LogInfo(fmt.Sprintf("✅ Successfully installed Java version %s.", javaVersion))
		}
		return nil
	},
}

func init() {
	StackCmd.AddCommand(installJavaCmd)
}
