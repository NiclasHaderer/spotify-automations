package models

type Config struct {
	Automations []Automation `json:"automations"`
	User        *User        `json:"user"`
}
