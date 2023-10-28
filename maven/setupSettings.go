package maven

import (
	"dx-cli/config"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var setupSettings = &cobra.Command{
	Use:   "settings",
	Short: "Setup settings",
	Run: func(cmd *cobra.Command, args []string) {
		currentContext, err := config.GetCurrentContext()
		if err != nil {
			fmt.Printf("Error fetching current context: %s\n", err)
			return
		}

		// Handle null or empty cases
		if currentContext == nil {
			fmt.Println("No current context defined.")
			return
		}

		mvnContext := currentContext.MavenContext
		if mvnContext == nil {
			fmt.Println("No Maven context defined.")
			return
		}

		setupMavenSettings(mvnContext)

	},
}

func setupMavenSettings(mavenContext *config.MavenContext) error {
	settingsPath := filepath.Join(os.Getenv("HOME"), ".m2", "settings.xml")

	// Create the directory if it doesn't exist
	if _, err := os.Stat(settingsPath); os.IsNotExist(err) {
		dir := filepath.Join(os.Getenv("HOME"), ".m2")
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %s", err)
		}
	}

	content := mavenContext.Settings.Content

	// Replace tokens in the content
	for _, token := range mavenContext.Settings.Tokens {
		// Assuming you have some way to resolve these tokens to actual values
		replacement := resolveToken(token)
		content = strings.ReplaceAll(content, fmt.Sprintf("{{%s}}", token.Name), replacement)
	}

	// Write the settings to the file
	if err := os.WriteFile(settingsPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write settings file: %s", err)
	}

	fmt.Println("Maven settings.xml has been configured.")
	return nil
}

func resolveToken(token config.Token) string {
	fmt.Printf("Token: %s\n", token.Name)
	if token.Description != nil {
		fmt.Printf("Description: %s\n", *token.Description)
	}
	if token.Link != nil {
		fmt.Println("Do you want to open the link for more info? [y/n]")
		var openLink string
		fmt.Scanln(&openLink)
		if openLink == "y" || openLink == "Y" {
			err := exec.Command("open", *token.Link).Start()
			if err != nil {
				fmt.Printf("Failed to open link: %s\n", err)
			}
		}
	}
	fmt.Printf("Enter the value for token %s: ", token.Name)
	var value string
	fmt.Scanln(&value)
	return value
}

func init() {
	MavenCmd.AddCommand(setupSettings)
}
