package utils

import "github.com/AlecAivazis/survey/v2"

func PromptUser(msg string, options []string) string {
	var input string
	prompt := &survey.Select{
		Message: msg,
		Options: options,
	}
	survey.AskOne(prompt, &input)
	return input
}
