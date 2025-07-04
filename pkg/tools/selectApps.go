package tools

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/manifoldco/promptui"
	"github.com/scriptnsam/blip-v2/internal/models"
)

func PromptSelectApps(apps []models.App) ([]models.App, error) {
	var selected []models.App

	for _, app := range apps {
		prompt := promptui.Select{
			Label: fmt.Sprintf("Add app: %s (%s)?", app.Name, app.Source),
			Items: []string{"Yes", "No"},
		}

		_, result, err := prompt.Run()
		if err != nil {
			return nil, err
		}

		if result == "Yes" {
			selected = append(selected, app)
		}
	}

	return selected, nil
}

func PromptMultiSelectApps(apps []models.App) ([]models.App, error) {
	var options []string
	appMap := make(map[string]models.App)

	for _, app := range apps {
		label := fmt.Sprintf("%s (%s)", app.Name, app.Source)
		options = append(options, label)
		appMap[label] = app
	}

	var selectedLabels []string
	prompt := &survey.MultiSelect{
		Message: "Select apps to save:",
		Options: options,
	}

	err := survey.AskOne(prompt, &selectedLabels)
	if err != nil {
		return nil, err
	}

	var selectedApps []models.App
	for _, label := range selectedLabels {
		selectedApps = append(selectedApps, appMap[label])
	}

	return selectedApps, nil
}
