package context

import (
	"dx-cli/config"
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

// Define a command to list contexts
var listContextsCmd = &cobra.Command{
	Use:   "list",
	Short: "Roll Call! 📜",
	Long: `Wondering about your available choices? 🤔

The 'list' command showcases all the contexts you've set up. It's your command-line directory! 📚

Scroll through your options and take your pick! 🎯`,
	Run: func(cmd *cobra.Command, args []string) {
		// Read YAML file
		path, err := config.GetConfigFilePath()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		data, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		// Parse into Config struct
		var config config.Config
		err = yaml.Unmarshal(data, &config)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		// List contexts
		for _, context := range config.Contexts {
			if context.Name == config.CurrentContext {
				fmt.Println("-> ", context.Name) // Arrow indicates the current context
			} else {
				fmt.Println("   ", context.Name)
			}
		}

	},
}

func init() {
	ContextCmd.AddCommand(listContextsCmd)
}
