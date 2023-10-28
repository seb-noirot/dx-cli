package context

import (
	"dx-cli/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	// ...
)

var deleteContextCmd = &cobra.Command{
	Use:   "delete [name]",
	Short: "Delete a context",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile(config.ConfigFilePath)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		var configObj config.Config
		err = yaml.Unmarshal(data, &configObj)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		for i, context := range configObj.Contexts {
			if context.Name == args[0] {
				configObj.Contexts = append(configObj.Contexts[:i], configObj.Contexts[i+1:]...)
				break
			}
		}

		newData, err := yaml.Marshal(&configObj)
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
	ContextCmd.AddCommand(deleteContextCmd)
}
