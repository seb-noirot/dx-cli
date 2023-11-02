package context

import (
	"dx-cli/config"
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var getCurrentContextCmd = &cobra.Command{
	Use:   "current",
	Short: "Show Your Active Playground 🌟",
	Long: `Curious where you are? 🤔

The 'current' command reveals the context you're currently working in. Think of it as your command-line GPS! 🗺

Stay oriented and keep sailing smoothly! 🚀`,
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

		fmt.Println("Current Context:", configObj.CurrentContext)
	},
}

func init() {
	ContextCmd.AddCommand(getCurrentContextCmd)
}
