package ide

import (
	"github.com/spf13/cobra"
)

var IdeCmd = &cobra.Command{
	Use:   "ide",
	Short: "IDE Management 🛠️",
	Long: `Manage your Integrated Development Environment settings and configurations with ease! 🌈

This command provides utilities to interact with your IDE, making it simple to set up, modify, or debug your development workspace. 💻

Use it to maximize productivity and make your coding life a breeze! 🚀`,
}
