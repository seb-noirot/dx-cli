package gitlab

import (
	"dx-cli/config"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// addGitLabCmd represents the add command
var addGitLabCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new GitLab definition",
	Run: func(cmd *cobra.Command, args []string) {
		// Fetch the current context
		currentContext, err := config.GetCurrentContext()
		if err != nil {
			fmt.Printf("Error fetching current context: %s\n", err)
			return
		}

		if currentContext == nil {
			fmt.Println("No current context defined.")
			return
		}

		// Prompt for name and host
		var name, host string
		fmt.Print("Enter the name of the GitLab definition: ")
		fmt.Scanln(&name)
		name = strings.TrimSpace(name)
		if name == "" {
			fmt.Println("Name cannot be empty.")
			return
		}

		fmt.Print("Enter the host of the GitLab definition: ")
		fmt.Scanln(&host)
		host = strings.TrimSpace(host)
		if host == "" {
			fmt.Println("Host cannot be empty.")
			return
		}

		// Create new GitLabContext and append it
		newGitLabContext := config.GitLabContext{Name: name, Host: host}
		currentContext.GitLabContexts = append(currentContext.GitLabContexts, newGitLabContext)

		// Update the current context
		err = config.UpdateCurrentContext(currentContext)
		if err != nil {
			fmt.Printf("Error updating current context: %s\n", err)
			return
		}

		fmt.Println("New GitLab definition added.")
	},
}

func init() {
	GitlabCmd.AddCommand(addGitLabCmd)
}
