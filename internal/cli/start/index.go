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
	PrintConfig
	ShowConfigPath
	Exit
)

var options = []string{"Start", "Login", "Logout", "Modify automations", "Print config", "Show config path", "Exit"}

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
	case "Print config":
		return PrintConfig
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
	if config.Instance.User == nil {
		promptOptions = append(options[:2], options[3:]...)
	} else {
		promptOptions = append(options[:1], options[2:]...)
	}

	sp := selection.New("", promptOptions)
	choice, _ := sp.RunPrompt()
	return fromString(choice)
}
