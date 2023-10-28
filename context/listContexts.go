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
	Short: "List all contexts",
	Run: func(cmd *cobra.Command, args []string) {
		// Read YAML file
		data, err := os.ReadFile(config.ConfigFilePath)
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
