package maven

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"runtime"
)

var cmdInstallMaven = &cobra.Command{
	Use:   "install",
	Short: "Installs Maven",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := exec.Command("git", "--version").Output()
		if err != nil {
			installMaven()
		}
		versionOutputInstalled, err := exec.Command("git", "--version").Output()
		fmt.Printf("Git is installed: %s\n", string(versionOutputInstalled))
		installMaven()
	},
}

func installMaven() {
	switch os := runtime.GOOS; os {
	case "darwin":
		// macOS
		err := exec.Command("brew", "install", "maven").Run()
		if err != nil {
			fmt.Printf("Failed to install Maven: %s\n", err)
		} else {
			fmt.Println("Maven installed successfully.")
		}
	case "linux":
		// Linux
		// Updating package lists
		err := exec.Command("sudo", "apt-get", "update").Run()
		if err != nil {
			fmt.Printf("Failed to update package lists: %s\n", err)
			return
		}

		// Installing Maven
		err = exec.Command("sudo", "apt-get", "install", "-y", "maven").Run()
		if err != nil {
			fmt.Printf("Failed to install Maven: %s\n", err)
		} else {
			fmt.Println("Maven installed successfully.")
		}
	default:
		fmt.Println("Your OS is not supported. Please install Maven manually.")
	}
}

func init() {
	MavenCmd.AddCommand(cmdInstallMaven)
}
