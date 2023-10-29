package maven

import (
	"dx-cli/config"
	"dx-cli/utils"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var verbose bool

var setupSettings = &cobra.Command{
	Use:   "settings",
	Short: "Supercharge Your Maven Settings ⚙️",
	Long: `Worried about managing your complex Maven settings? 🤔
	
The 'settings' command handles all the heavy lifting. 
Simply provide your Maven context and let the magic happen. 🌟`,
	Run: func(cmd *cobra.Command, args []string) {
		currentContext, err := utils.GetCurrentContext(config.ConfigFilePath, verbose) // true for verbose

		if err != nil {
			utils.Printf(true, "🚨 Error fetching current context: %s\n", err)
			return
		}

		if currentContext == nil {
			utils.Println(true, "🚫 No current context defined.")
			return
		}

		mvnContext := currentContext.MavenContext
		if mvnContext == nil {
			utils.Println(true, "🚫 No Maven context defined.")
			return
		}

		setupMavenSettings(mvnContext)
	},
}

func setupMavenSettings(mavenContext *config.MavenContext) error {
	settingsPath := filepath.Join(os.Getenv("HOME"), ".m2", "settings.xml")

	if _, err := os.Stat(settingsPath); os.IsNotExist(err) {
		dir := filepath.Join(os.Getenv("HOME"), ".m2")
		if err := os.MkdirAll(dir, 0755); err != nil {
			utils.Printf(true, "🚨 Failed to create directory: %s\n", err)
			return nil
		}
	}

	content := mavenContext.Settings.Content
	for _, token := range mavenContext.Settings.Tokens {
		replacement := resolveToken(token)
		content = strings.ReplaceAll(content, fmt.Sprintf("{{%s}}", token.Name), replacement)
	}

	if err := os.WriteFile(settingsPath, []byte(content), 0644); err != nil {
		utils.Printf(true, "🚨 Failed to write settings file: %s\n", err)
		return nil
	}

	utils.Println(true, "✅ Maven settings.xml has been configured.")
	return nil
}

func resolveToken(token config.Token) string {
	var value string

	// Display the token name
	utils.Printf(true, "🔑 Token: %s\n", token.Name)

	// Display the token description if available
	if token.Description != nil && *token.Description != "" {
		utils.Printf(true, "📝 Description: %s\n", *token.Description)
	}

	// Offer to open the link if available
	if token.Link != nil && *token.Link != "" {
		if utils.PromptUser("🔗 Do you want to open the link for more info?   ", []string{"Yes", "No"}) == "Yes" {
			err := exec.Command("open", *token.Link).Start()
			if err != nil {
				utils.Printf(true, "🚨 Failed to open link: %s\n", err)
			}
		}
	}

	// Get the value for the token
	valuePrompt := &survey.Input{
		Message: fmt.Sprintf("Enter the value for token %s: ", token.Name),
	}
	survey.AskOne(valuePrompt, &value)

	return value
}

func init() {
	setupSettings.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	MavenCmd.AddCommand(setupSettings)
}
