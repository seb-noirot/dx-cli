package context

import (
	"dx-cli/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	// ...
)

// Define a command to set the current context
var setContextCmd = &cobra.Command{
	Use:   "set-current",
	Short: "Set the current context",
	Run: func(cmd *cobra.Command, args []string) {
		// Read existing config
		data, err := os.ReadFile(config.ConfigFilePath)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		var conf config.Config
		err = yaml.Unmarshal(data, &conf)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		// Set current context (you might want to add validation here)
		conf.CurrentContext = args[0]

		// Write updated config back to file
		newData, err := yaml.Marshal(&conf)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		err = os.WriteFile(config.ConfigFilePath, newData, 0644)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	},
}

func init() {
	ContextCmd.AddCommand(setContextCmd)
}
