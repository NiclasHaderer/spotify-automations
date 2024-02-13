package main

import (
	"github.com/joho/godotenv"
	"log"
	"spotify-automations/internal/automation"
	"spotify-automations/internal/config"
	"spotify-automations/internal/spotify_wrapper"
	"spotify-automations/internal/start"
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
			spotify_wrapper.Login()
		case start.Logout:
			spotify_wrapper.Logout()
		case start.ModifyAutomations:
			automation.SelectAutomation()
		case start.ShowConfig:
			config.Get().Print()
		case start.ShowConfigPath:
			config.PrintPath()
		case start.Exit:
			return
		}
	}
}
