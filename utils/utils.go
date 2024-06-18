package utils

import (
	"github.com/AlecAivazis/survey/v2"
)

func RemoveDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	var result []string

	for v := range elements {
		if !encountered[elements[v]] {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}

// AskConfirmation will request confirmation from the user
func AskConfirmation(message string) bool {
	var confirmation bool
	prompt := &survey.Confirm{
		Message: message,
	}
	if err := survey.AskOne(prompt, &confirmation); err != nil {
		return false
	}
	return confirmation
}
