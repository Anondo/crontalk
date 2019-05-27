package generator

import (
	"github.com/manifoldco/promptui"
)

// GenerateCron generates a cron expression by prompting english word options
func GenerateCron() (string, error) {
	cronExpression := ""

	prompt := promptui.Select{
		Label: "Minute",
		Items: []string{"every"},
	}

	_, min, err := prompt.Run()
	if err != nil {
		return "", err
	}

	if min == "every" {
		cronExpression += "*"
	}
	return cronExpression, nil
}
