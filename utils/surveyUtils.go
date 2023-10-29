package utils

import "github.com/AlecAivazis/survey/v2"

func PromptUser(prompt string, options []string) string {
	var result string

	// If options are provided, use select prompt
	if options != nil && len(options) > 0 {
		prompt := &survey.Select{
			Message: prompt,
			Options: options,
		}
		survey.AskOne(prompt, &result, nil)
	} else {
		// Else, use input prompt
		prompt := &survey.Input{
			Message: prompt,
		}
		survey.AskOne(prompt, &result, nil)
	}

	return result
}
