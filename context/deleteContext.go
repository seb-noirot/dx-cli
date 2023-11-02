package context

import (
	"dx-cli/config"
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	// ...
)

var deleteContextCmd = &cobra.Command{
	Use:   "delete [name]",
	Short: "Banish from the Kingdom 🚫🏰",
	Long: `Oops, made a wrong choice? No worries! 🙌

Use 'delete' to remove a context you no longer need. Think of it as banishing it from your kingdom, making way for new possibilities and cleaner config. 🗑️

Gone but not forgotten! 😢`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path, err := config.GetConfigFilePath()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		data, err := os.ReadFile(path)
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

		err = os.WriteFile(path, newData, 0644)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	},
}

func init() {
	ContextCmd.AddCommand(deleteContextCmd)
}
