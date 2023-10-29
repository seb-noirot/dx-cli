package gitlabstack

import (
	"dx-cli/config"
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
	Short: "Checkout all GitLab projects for a stack",
	Long: `This command allows you to select a GitLab context and a stack within that context.
Then, it checks out all projects associated with the selected stack into the specified directory. 
Perfect for getting your development environment set up quickly! 🚀`,
	Run: func(cmd *cobra.Command, args []string) {
		selectedGitLab, err := utils.SelectGitlabDefinition()
		if err != nil {
			utils.LogError("❌ Failed to select GitLab definition: %s", err)
			return
		}
		if selectedGitLab == nil {
			utils.LogInfo("ℹ️ No GitLab context selected.")
			return
		}

		selectedStack, err := utils.SelectGitlabStack(selectedGitLab)
		if err != nil {
			utils.LogError("❌ Failed to select stack: %s", err)
			return
		}
		if selectedStack == nil {
			utils.LogInfo("ℹ️ No stacks defined.")
			return
		}

		err = checkoutProjects(selectedStack, selectedGitLab)
		if err != nil {
			utils.LogError("❌ Failed to checkout projects: %s", err)
		}
	},
}

func checkoutProjects(stack *config.GitlabStack, gitlab *config.GitLabContext) error {
	directory := filepath.Join(os.Getenv("HOME"), stack.Path)
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.MkdirAll(directory, os.ModePerm)
		if err != nil {
			return err
		}
		utils.LogInfo(fmt.Sprintf("📁 Created directory: %s", directory))
	}

	for _, project := range stack.Projects {
		err := checkoutSingleProject(directory, project, stack, gitlab)
		if err != nil {
			utils.LogError(fmt.Sprintf("❌ Failed to clone project %s", project), err)
		}
	}

	return nil
}

func checkoutSingleProject(directory string, project string, stack *config.GitlabStack, gitlab *config.GitLabContext) error {
	projectDir := filepath.Join(directory, filepath.Base(strings.TrimSuffix(project, ".git")))
	trimmedHost := strings.TrimSuffix(gitlab.Host, "/")
	cloneURL := fmt.Sprintf("git@%s:%s", strings.Replace(trimmedHost, "https://", "", 1), project)

	if _, err := os.Stat(projectDir); os.IsNotExist(err) {
		cmd := exec.Command("git", "clone", cloneURL, projectDir)
		err := cmd.Run()
		if err != nil {
			return err
		}
		utils.LogInfo(fmt.Sprintf("Successfully cloned project %s in folder %s.", cloneURL, projectDir))
	} else {
		utils.LogInfo(fmt.Sprintf("Project %s already exists in %s, skipping.", cloneURL, projectDir))
	}

	return nil
}

func init() {
	StackCmd.AddCommand(installStackCmd)
}
