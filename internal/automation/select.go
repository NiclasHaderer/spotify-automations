package automation

import (
	"github.com/erikgeiser/promptkit/selection"
	"spotify-automations/internal/models"
	"spotify-automations/internal/spotify_wrapper"
	"spotify-automations/internal/utils"
)

func SelectAutomation() {
	options := utils.Map(Options, func(option models.AutomationOption) string {
		return option.Name
	})
	options = append(options, "Back")
	sp := selection.New("Select Automation", options)
	choice, _ := sp.RunPrompt()
	retrievedOption := utils.Find(func(option models.AutomationOption) bool {
		return option.Name == choice
	}, Options)
	if retrievedOption == nil {
		return
	}

	client := spotify_wrapper.NewClient()
	retrievedOption.CreateOrModify(client)
}
