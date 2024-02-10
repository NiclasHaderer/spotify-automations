package main

import (
	"github.com/joho/godotenv"
	"log"
	"spotify-automations/internal/cli/start"
	"spotify-automations/internal/config"
	"spotify-automations/internal/spotify"
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
		case start.ShowConfig:
			config.Print(config.Get())
		case start.ShowConfigPath:
			config.PrintPath()
		case start.Exit:
			return
		}
	}
}
