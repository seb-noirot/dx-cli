package ide

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

var installJetbrains = &cobra.Command{
	Use:   "jetbrains",
	Short: "Installs jetbrains",
	Run: func(cmd *cobra.Command, args []string) {
		installJetBrainsToolbox()
	},
}

// InstallJetBrainsToolbox installs JetBrains Toolbox if not already installed
func installJetBrainsToolbox() {
	// Check if JetBrains Toolbox is already installed
	err := exec.Command("open", "-a", "JetBrains Toolbox").Run()

	if err != nil {
		fmt.Println("JetBrains Toolbox is not installed. Installing...")

		// Install JetBrains Toolbox using Homebrew Cask
		err = exec.Command("brew", "install", "--cask", "jetbrains-toolbox").Run()

		if err != nil {
			fmt.Printf("Failed to install JetBrains Toolbox: %s\n", err)
		} else {
			fmt.Println("JetBrains Toolbox installed successfully.")
		}

	} else {
		fmt.Println("JetBrains Toolbox is already installed.")
	}
}

func init() {
	IdeCmd.AddCommand(installJetbrains)
}
