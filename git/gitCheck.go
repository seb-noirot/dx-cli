package git

import (
	"dx-cli/utils"
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"runtime"
)

var checkGitCmd = &cobra.Command{
	Use:   "check",
	Short: "Check for Git installation",
	Long:  "This command checks for the presence of Git on your system and offers to install it if it's not found.",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := exec.LookPath("git")
		if err != nil {
			if utils.PromptUser("Git is not installed. Would you like to install it?? 🎵", []string{"Yes", "No"}) == "Yes" {
				return installGit()
			}
			return fmt.Errorf("Git is required for this program to function.")
		}

		versionOutput, err := exec.Command("git", "--version").Output()
		if err != nil {
			return err
		}

		utils.LogInfo(fmt.Sprintf("Git is installed: %s", string(versionOutput)))
		return nil
	},
}

func installGit() error {
	var err error
	switch runtime.GOOS {
	case "darwin":
		utils.LogInfo("Installing Git via Homebrew...")
		err = exec.Command("brew", "install", "git").Run()
	case "linux":
		utils.LogInfo("Installing Git via apt-get...")
		err = exec.Command("sudo", "apt-get", "install", "-y", "git").Run()
	default:
		err = fmt.Errorf("Unsupported OS. Please install Git manually.")
	}
	return err
}

func init() {
	GitCmd.AddCommand(checkGitCmd)
}
