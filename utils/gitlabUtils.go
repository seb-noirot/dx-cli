package utils

import (
	"dx-cli/config"
	"fmt"
)

func SelectGitlabDefinition() (*config.GitLabContext, error) {
	// Fetch the current context
	path, err := config.GetConfigFilePath()
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	currentContext, err := GetCurrentContext(path, false)
	if err != nil {
		LogError("Error fetching current context: %s", err)
		return nil, err
	}

	if currentContext == nil {
		LogWarning("No current context defined.")
		return nil, fmt.Errorf("No current context defined")
	}

	if len(currentContext.GitLabContexts) == 0 {
		LogWarning("No GitLab definitions available.")
		return nil, fmt.Errorf("No GitLab definitions available")
	}

	// Prepare names for survey
	names := make([]string, len(currentContext.GitLabContexts))
	for i, glContext := range currentContext.GitLabContexts {
		names[i] = fmt.Sprintf("%s (%s)", glContext.Name, glContext.Host)
	}

	// Prompt user for selection
	selected := PromptUser("Select a GitLab definition:", names)
	if selected == "" {
		LogWarning("Invalid choice.")
		return nil, fmt.Errorf("Invalid choice")
	}

	// Find the selected context
	var selectedGitLab *config.GitLabContext
	LogInfo("Selected name " + selected)
	for _, gitLabContext := range currentContext.GitLabContexts {
		if fmt.Sprintf("%s (%s)", gitLabContext.Name, gitLabContext.Host) == selected {
			selectedGitLab = &gitLabContext
			break
		}
	}
	return selectedGitLab, nil
}
