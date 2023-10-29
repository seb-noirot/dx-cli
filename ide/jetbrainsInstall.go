package ide

import (
	"dx-cli/utils"
	"github.com/spf13/cobra"
	"os/exec"
)

var installJetbrains = &cobra.Command{
	Use:   "jetbrains",
	Short: "Installs JetBrains Toolbox 🛠️",
	Long:  "Installs or updates the JetBrains Toolbox, the gateway to all JetBrains IDEs 🌈",
	Run: func(cmd *cobra.Command, args []string) {
		err := exec.Command("open", "-a", "JetBrains Toolbox").Run()
		if err != nil {
			installJetBrainsToolbox()
		} else {
			utils.LogInfo("JetBrains Toolbox is already installed. 👍")
		}

	},
}

// InstallJetBrainsToolbox installs JetBrains Toolbox if not already installed
func installJetBrainsToolbox() {
	if utils.PromptUser("Would you like to proceed with the installation of JetBrains Toolbox? 🤔", []string{"Yes", "No"}) == "Yes" {
		// Install JetBrains Toolbox using Homebrew Cask
		err := exec.Command("brew", "install", "--cask", "jetbrains-toolbox").Run()

		if err != nil {
			utils.LogError("Failed to install JetBrains Toolbox: 😢", err)
		} else {
			utils.LogInfo("JetBrains Toolbox installed successfully. 🎉")
		}
	} else {
		utils.LogInfo("Installation aborted by the user. 🚫")
	}
}

func init() {
	IdeCmd.AddCommand(installJetbrains)
}
