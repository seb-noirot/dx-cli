package git

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var checkGitCmd = &cobra.Command{
	Use:   "check",
	Short: "Check for Git installation",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := exec.LookPath("git")
		if err != nil {
			fmt.Println("Git is not installed. Would you like to install it? (y/n)")
			var response string
			fmt.Scanln(&response)
			if strings.ToLower(response) == "y" {
				return installGit()
			}
			return fmt.Errorf("Git is required for this program to function")
		}

		versionOutput, err := exec.Command("git", "--version").Output()
		if err != nil {
			return err
		}

		fmt.Printf("Git is installed: %s\n", string(versionOutput))
		return nil
	},
}

func installGit() error {
	switch runtime.GOOS {
	case "darwin":
		fmt.Println("Installing Git via Homebrew...")
		err := exec.Command("brew", "install", "git").Run()
		return err
	case "linux":
		fmt.Println("Installing Git via apt-get...")
		err := exec.Command("sudo", "apt-get", "install", "-y", "git").Run()
		return err
	default:
		return fmt.Errorf("Unsupported OS. Please install Git manually.")
	}
}

func init() {
	GitCmd.AddCommand(checkGitCmd)
}
