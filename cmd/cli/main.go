package main

import (
	"github.com/joho/godotenv"
	"log"
	"spotify-automations/internal/config"
	"spotify-automations/internal/spotify"
	"spotify-automations/internal/start"
	"spotify-automations/internal/textarea"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadEnv()
	for {
		option := start.NewStartCommand()
		switch option {
		case start.Start:
		case start.Login:
			spotify.Login()
		case start.Logout:
			spotify.Logout()
		case start.ModifyAutomations:
			textarea.New("Modify automations", "Not implemented yet", false)
		case start.ShowConfig:
			config.Get().Print()
		case start.ShowConfigPath:
			config.PrintPath()
		case start.Exit:
			return
		}
	}
}
