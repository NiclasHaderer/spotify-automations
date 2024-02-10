package start

import (
	"github.com/erikgeiser/promptkit/selection"
	"log"
	"spotify-automations/internal/config"
)

type Command int

const (
	Start Command = iota
	Login
	Logout
	ModifyAutomations
	ShowConfig
	ShowConfigPath
	Exit
)

var options = []string{"Start", "Logout", "Modify automations", "Show config", "Show config path", "Exit"}

func fromString(s string) Command {
	switch s {
	case "Start":
		return Start
	case "Login":
		return Login
	case "Logout":
		return Logout
	case "Modify automations":
		return ModifyAutomations
	case "Show config":
		return ShowConfig
	case "Show config path":
		return ShowConfigPath
	case "Exit":
		return Exit
	}
	log.Fatalf("Invalid start command: %s", s)
	return -1
}

func NewStartCommand() Command {

	var promptOptions []string
	if config.Get().User == nil {
		return Login
	} else {
		promptOptions = options[1:]
	}

	sp := selection.New("", promptOptions)
	choice, _ := sp.RunPrompt()
	return fromString(choice)
}
