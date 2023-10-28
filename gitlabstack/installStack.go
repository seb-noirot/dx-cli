package gitlabstack

import (
	"dx-cli/utils"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var installStackCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Checkout all GitLab project for a stack",
	Run: func(cmd *cobra.Command, args []string) {
		// Select GitLab definition
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

		// Create directory if not exists
		if _, err := os.Stat(selectedStack.Path); os.IsNotExist(err) {
			os.MkdirAll(selectedStack.Path, os.ModePerm)
		}

		// Loop through projects and clone if they do not exist
		for _, project := range selectedStack.Projects {
			projectDir := strings.TrimSuffix(filepath.Join(os.Getenv("HOME"), selectedStack.Path, filepath.Base(project)), ".git")
			trimmedHost := strings.TrimSuffix(selectedGitLab.Host, "/") // Remove trailing slash
			cloneURL := fmt.Sprintf("git@%s:%s", strings.Replace(trimmedHost, "https://", "", 1), project)

			if _, err := os.Stat(projectDir); os.IsNotExist(err) {
				cmd := exec.Command("git", "clone", cloneURL, projectDir)
				err := cmd.Run()
				if err != nil {
					fmt.Printf("Failed to clone project %s: %s\n", cloneURL, err)
				} else {
					fmt.Printf("Successfully cloned project %s in folder %s.\n", cloneURL, projectDir)
				}
			} else {
				fmt.Printf("Project %s already exists in %s, skipping.\n", cloneURL, projectDir)
			}
		}
	},
}

func init() {
	StackCmd.AddCommand(installStackCmd)
}
