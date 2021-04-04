package generator

import "github.com/manifoldco/promptui"

// runPrompt executes a prompt to chose an item from the cli
func runPrompt(label string, items []string) (string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	_, res, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return res, err
}
