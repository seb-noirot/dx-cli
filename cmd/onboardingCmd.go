package cmd

import (
	"dx-cli/config"
	"dx-cli/utils"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"log"
)

var onboardingCmd = &cobra.Command{
	Use:   "onboarding",
	Short: "🚀 Executes onboarding commands",
	Long:  "🚀 This command performs all the steps necessary for a smooth onboarding.",
	Run: func(cmd *cobra.Command, args []string) {
		// Load the current context
		currentContext, err := utils.GetCurrentContext(config.ConfigFilePath, false)
		if err != nil {
			utils.LogError("Failed to load current context: %s", err)
			return
		}

		if currentContext.Onboarding == nil {
			utils.LogWarning("No onboarding commands defined in the current context.")
			return
		}

		totalCommands := len(currentContext.Onboarding.OnboardingCmds)
		log.Printf("🚀 Going to run %d commands...\n", totalCommands)

		// Iterate through each onboarding command
		for _, onboardingCmd := range currentContext.Onboarding.OnboardingCmds {
			// Log info
			utils.LogInfo(fmt.Sprintf("📋 Title: %s", onboardingCmd.Name))
			utils.LogInfo(fmt.Sprintf("📚 Description: %s", onboardingCmd.Description))

			// Copy to clipboard
			err := clipboard.WriteAll("go run main.go " + onboardingCmd.Execution)
			if err != nil {
				utils.LogInfo(fmt.Sprintf("📝 Command to run: go run main.go %s", onboardingCmd.Execution))
			} else {
				utils.LogInfo(fmt.Sprintf("📝 Command to run: %s (copied to your clipboard)", onboardingCmd.Execution))
			}

			installed := utils.PromptUser(fmt.Sprintf("Did you run the command %s successfully?", onboardingCmd.Execution), []string{"Yes", "Skip", "No", "Exit"})
			if installed == "No" {
				utils.LogInfo("Hope you won't regret it!")
			}
			if installed == "Exit" {
				utils.LogInfo("See you later!!!")
				return
			}
			utils.LogInfo("                     ")
		}
	},
}

func init() {
	rootCmd.AddCommand(onboardingCmd)
}
