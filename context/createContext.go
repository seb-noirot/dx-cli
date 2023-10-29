package context

import (
	"dx-cli/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	// ...
)

var createContextCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Spawn a New Context in a Flash ⚡",
	Long: `Say hello to your new playground! 🎉

Use 'create [name]' to set up a new context that houses all your project-specific configurations.
Just give it a name, and you're ready to configure away! 🛠

Start building your perfect environment now! 🌈`,
	Args: cobra.ExactArgs(1),
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

		newContext := config.Context{Name: args[0]}
		configObj.Contexts = append(configObj.Contexts, newContext)

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
	ContextCmd.AddCommand(createContextCmd)
}
