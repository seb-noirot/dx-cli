package utils

import (
	"dx-cli/config"
	"github.com/AlecAivazis/survey/v2"
)

func SelectGitlabStack(selectedGitLab *config.GitLabContext) (*config.GitlabStack, error) {
	if len(selectedGitLab.GitlabStacks) == 0 {
		LogInfo("No stacks defined.")
		return nil, nil
	}

	names := make([]string, len(selectedGitLab.GitlabStacks))
	for i, stack := range selectedGitLab.GitlabStacks {
		names[i] = stack.Name
	}

	selected := ""
	prompt := &survey.Select{
		Message: "Select a gitlabstack to install:",
		Options: names,
	}
	err := survey.AskOne(prompt, &selected)
	if err != nil {
		return nil, err
	}

	var selectedStack *config.GitlabStack
	for i, stack := range selectedGitLab.GitlabStacks {
		if stack.Name == selected {
			selectedStack = &selectedGitLab.GitlabStacks[i] // Directly reference from the slice
			break
		}
	}

	return selectedStack, nil
}
