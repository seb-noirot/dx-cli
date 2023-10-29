package maven

import (
	// ... your other imports
	"dx-cli/utils"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"os/exec"
	"runtime"
)

var cmdInstallMaven = &cobra.Command{
	Use:   "install",
	Short: "Get Maven Up and Running 🛠️",
	Long: `No more manual downloads or PATH tweaks! 🙅‍♂️

This command makes Maven installation as easy as pie. 🥧
Choose your operating system, and let the magic happen.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := exec.Command("mvn", "--version").Output()
		if err != nil {
			installMaven()
		}
		versionOutputInstalled, err := exec.Command("mvn", "--version").Output()
		utils.Printf(true, "🛠️ Maven is installed: %s\n", string(versionOutputInstalled))
	},
}

func installMaven() {
	var proceed string
	prompt := &survey.Select{
		Message: "Install Maven?",
		Options: []string{"Yes", "No"},
	}
	survey.AskOne(prompt, &proceed)

	if proceed == "No" {
		utils.Println(true, "👋 Okay, maybe next time!")
		return
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		// macOS
		err := exec.Command("brew", "install", "maven").Run()
		if err != nil {
			utils.Printf(true, "🚨 Failed to install Maven: %s\n", err)
		} else {
			utils.Println(true, "🎉 Maven installed successfully.")
		}
	case "linux":
		// Linux
		// Updating package lists
		err := exec.Command("sudo", "apt-get", "update").Run()
		if err != nil {
			utils.Printf(true, "🚨 Failed to update package lists: %s\n", err)
			return
		}

		// Installing Maven
		err = exec.Command("sudo", "apt-get", "install", "-y", "maven").Run()
		if err != nil {
			utils.Printf(true, "🚨 Failed to install Maven: %s\n", err)
		} else {
			utils.Println(true, "🎉 Maven installed successfully.")
		}
	default:
		utils.Println(true, "🤷 Your OS is not supported. Please install Maven manually.")
	}
}

func init() {
	MavenCmd.AddCommand(cmdInstallMaven)
}
