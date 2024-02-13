package automation

import (
	"github.com/erikgeiser/promptkit/selection"
	"spotify-automations/internal/automation/library_mirror"
	"spotify-automations/internal/models"
	"spotify-automations/internal/utils"
)

var Options = []models.AutomationOption{library_mirror.PlaylistAutomationOption}

func SelectAutomation() {
	options := utils.Map(func(option models.AutomationOption) string {
		return option.Name
	}, Options)
	options = append(options, "Back")
	sp := selection.New("Select Automation", options)
	choice, _ := sp.RunPrompt()
	retrievedOption := utils.Find(func(option models.AutomationOption) bool {
		return option.Name == choice
	}, Options)
	if retrievedOption == nil {
		return
	}

	retrievedOption.CreateOrModify()
}
