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
	Short: "Get the current context",
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

		fmt.Println("Current Context:", configObj.CurrentContext)
	},
}

func init() {
	ContextCmd.AddCommand(getCurrentContextCmd)
}
