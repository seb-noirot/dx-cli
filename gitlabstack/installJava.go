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
	Run: func(cmd *cobra.Command, args []string) {
		selectedGitLab, err := utils.SelectGitlabDefinition()
		if err != nil {
			return
		}
		if selectedGitLab == nil {
			return
		}

		// List stacks and ask for selection
		if len(selectedGitLab.GitlabStacks) == 0 {
			fmt.Println("No stacks defined.")
			return
		}

		fmt.Println("Select a gitlabstack to install:")
		for i, stack := range selectedGitLab.GitlabStacks {
			fmt.Printf("[%d] %s\n", i+1, stack.Name)
		}

		var choice int
		fmt.Scanln(&choice)
		if choice < 1 || choice > len(selectedGitLab.GitlabStacks) {
			fmt.Println("Invalid choice.")
			return
		}

		selectedStack := selectedGitLab.GitlabStacks[choice-1]

		// Check if a Java version is defined for the selected stack
		if len(selectedStack.Javas) == 0 {
			fmt.Println("No Java versions defined for this stack.")
			return
		}

		for _, javaVersion := range selectedStack.Javas {
			// Install Java using SDKMAN
			fmt.Printf("Installing Java version %s using SDKMAN...\n", javaVersion)

			cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("source $HOME/.sdkman/bin/sdkman-init.sh && sdk install java %s", javaVersion))
			err = cmd.Run()
			if err != nil {
				fmt.Printf("Failed to install Java version %s: %s\n", javaVersion, err)
				return
			}

			fmt.Printf("Successfully installed Java version %s.\n", javaVersion)
		}

		return
	},
}

func init() {
	StackCmd.AddCommand(installJavaCmd)
}
