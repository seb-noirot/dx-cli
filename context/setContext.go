package context

import (
	"dx-cli/config"
	"dx-cli/utils"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"os"
	// ...
)

var verbose bool

// Define a command to set the current context
var setContextCmd = &cobra.Command{
	Use:   "set-current",
	Short: "Crown a King 👑",
	Long: `Feel like switching gears? 🔄

Use 'set-current' to specify which context you want to rule over. Just like placing a crown on your chosen king, this command sets the 'current' context to your specified name. 🌐

Be the master of your realm! 💪`,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := config.GetConfigFilePath()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		conf, err := utils.LoadConfig(path, verbose)

		if err != nil {
			return
		}

		utils.Println(verbose, "📋 Parsing contexts...")

		// Create a list of context names
		var contextNames []string
		for _, ctx := range conf.Contexts {
			contextNames = append(contextNames, ctx.Name)
		}
		utils.Println(verbose, "✅ Contexts parsed.")

		// Use survey to get user's choice for context
		var selectedContextName string
		prompt := &survey.Select{
			Message: "🌍 Choose a context:",
			Options: contextNames,
		}
		utils.Println(true, "🤖 Awaiting your selection...")
		err = survey.AskOne(prompt, &selectedContextName)
		if err != nil {
			utils.Println(true, "🚨 Oops! Something went wrong.")
			return
		}
		utils.Printf(true, "🎉 You've chosen: %s\n", selectedContextName)

		// Set current context
		utils.Println(true, "🔄 Setting current context...")
		conf.CurrentContext = selectedContextName

		// Write updated config back to file
		utils.Println(true, "📝 Updating config file...")
		newData, err := yaml.Marshal(&conf)
		if err != nil {
			utils.Printf(true, "🚨 Oops! Failed to marshal updated config: %s\n", err)
			return
		}
		err = os.WriteFile(path, newData, 0644)
		if err != nil {
			utils.Printf(true, "🚨 Oops! Failed to update config file: %s\n", err)
			return
		}
		utils.Println(true, "✅ Successfully set the current context to "+selectedContextName)
	},
}

func init() {
	setContextCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	ContextCmd.AddCommand(setContextCmd)
}
