package models

import "github.com/zmb3/spotify/v2"

type Automation struct {
	Config string `json:"config"`
}

type AutomationOption struct {
	Name           string
	CreateOrModify func(client *spotify.Client)
	Run            func(client *spotify.Client)
}
