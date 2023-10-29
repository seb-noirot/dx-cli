package sdkman

import (
	"dx-cli/utils"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var SdkmanCmd = &cobra.Command{
	Use:   "sdkman",
	Short: "Your SDK Guardian Angel 👼",
	Long: `Ever feel like you're drowning in a sea of SDK versions? 🌊

No worries, the 'sdkman' command is here to keep you afloat! 🚢
From installing to updating, and even switching between SDK versions, sdkman makes it a breeze.

Think of it as your personal SDK concierge, always at your service. 🛎️`,
	Run: func(cmd *cobra.Command, args []string) {
		checkSdkman()
	},
}

func checkSdkman() {
	utils.Println(true, "🔍 Checking for SDKMAN installation...")
	_, err := exec.Command("bash", "-c", "source $HOME/.sdkman/bin/sdkman-init.sh && sdk version").Output()

	if err != nil {
		var install bool
		prompt := &survey.Confirm{
			Message: "❗ SDKMAN is not installed. Would you like to install it?",
		}
		survey.AskOne(prompt, &install)

		if install {
			installSdkman()
		} else {
			utils.Println(true, "🔗 To install SDKMAN manually, visit: https://sdkman.io/install")
		}
		return
	}

	utils.Println(true, "✅ SDKMAN is already installed.")
}

func installSdkman() {
	utils.Println(true, "📦 Installing SDKMAN...")

	_, err := exec.Command("bash", "-c", "curl -s \"https://get.sdkman.io\" | bash").Output()
	if err != nil {
		utils.Printf(true, "🚨 Failed to install SDKMAN: %v\n", err)
		return
	}

	_, err = exec.Command("bash", "-c", "source \"$HOME/.sdkman/bin/sdkman-init.sh\"").Output()
	if err != nil {
		utils.Printf(true, "🚨 Failed to source bash profile: %v\n", err)
		return
	}

	utils.Println(true, "✅ SDKMAN installed successfully.")
	addToZshrc()
}

func addToZshrc() {
	utils.Println(true, "📝 Adding SDKMAN initialization to .zshrc...")

	zshrcPath := fmt.Sprintf("%s/.zshrc", os.Getenv("HOME"))
	file, err := os.OpenFile(zshrcPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Printf("🚨 Failed to open or create .zshrc. Error: %v\n", err)
		return
	}
	defer file.Close()

	lines := `
export SDKMAN_DIR="$HOME/.sdkman"
[[ -s "$HOME/.sdkman/bin/sdkman-init.sh" ]] && source "$HOME/.sdkman/bin/sdkman-init.sh"
`

	if _, err := file.WriteString(lines); err != nil {
		fmt.Printf("🚨 Failed to write to .zshrc. Error: %v\n", err)
		return
	}

	utils.Println(true, "✅ Added SDKMAN initialization to .zshrc")
}
