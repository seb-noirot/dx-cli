package sdkman

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var SdkmanCmd = &cobra.Command{
	Use:   "sdkman",
	Short: "Manage SDKMAN",
	Run: func(cmd *cobra.Command, args []string) {
		checkSdkman()
	},
}

func checkSdkman() {
	_, err := exec.Command("bash", "-c", "source $HOME/.sdkman/bin/sdkman-init.sh && sdk version").Output()

	if err != nil {
		fmt.Println("SDKMAN is not installed. Would you like to install it? (y/n)")

		var input string
		fmt.Scanln(&input)

		if input == "y" || input == "Y" {
			installSdkman()
		} else {
			fmt.Println("To install SDKMAN manually, visit: https://sdkman.io/install")
		}
		return
	}

	fmt.Println("SDKMAN is installed.")
}

func installSdkman() {
	// Logic to install SDKMAN
	fmt.Println("Installing SDKMAN...")

	// SDKMAN installation typically involves a curl command
	_, err := exec.Command("bash", "-c", "curl -s \"https://get.sdkman.io\" | bash").Output()
	if err != nil {
		fmt.Println("Failed to install SDKMAN: ", err)
		return
	}

	// Source bash profile to complete the installation
	_, err = exec.Command("bash", "-c", "source \"$HOME/.sdkman/bin/sdkman-init.sh\"").Output()
	if err != nil {
		fmt.Println("Failed to source bash profile: ", err)
		return
	}

	fmt.Println("SDKMAN installed successfully.")
}

func addToZshrc() {
	zshrcPath := fmt.Sprintf("%s/.zshrc", os.Getenv("HOME"))
	file, err := os.OpenFile(zshrcPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Printf("Failed to open or create .zshrc. Error: %v\n", err)
		return
	}
	defer file.Close()

	lines := `
export SDKMAN_DIR="$HOME/.sdkman"
[[ -s "$HOME/.sdkman/bin/sdkman-init.sh" ]] && source "$HOME/.sdkman/bin/sdkman-init.sh"
`

	if _, err := file.WriteString(lines); err != nil {
		fmt.Printf("Failed to write to .zshrc. Error: %v\n", err)
		return
	}

	fmt.Println("Added SDKMAN initialization to .zshrc")
}
