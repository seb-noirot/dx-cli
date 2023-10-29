package utils

import (
	"dx-cli/config"
	"fmt"
)

func SelectGitlabDefinition() (*config.GitLabContext, error) {
	// Fetch the current context
	currentContext, err := GetCurrentContext(config.ConfigFilePath, false)
	if err != nil {
		fmt.Printf("Error fetching current context: %s\n", err)
		return nil, nil
	}

	if currentContext == nil {
		fmt.Println("No current context defined.")
		return nil, nil
	}

	if len(currentContext.GitLabContexts) == 0 {
		fmt.Println("No GitLab definitions available.")
		return nil, nil
	}

	fmt.Println("Select a GitLab definition:")
	for i, glContext := range currentContext.GitLabContexts {
		fmt.Printf("[%d] %s (%s)\n", i+1, glContext.Name, glContext.Host)
	}

	var choice int
	fmt.Scanln(&choice)

	if choice < 1 || choice > len(currentContext.GitLabContexts) {
		fmt.Println("Invalid choice.")
		return nil, nil
	}
	selectedGitLab := currentContext.GitLabContexts[choice-1]
	return &selectedGitLab, nil
}
