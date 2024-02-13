package main

import (
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"spotify-automations/internal/automation"
	"spotify-automations/internal/config"
	"spotify-automations/internal/spotify_wrapper"
	"spotify-automations/internal/start"
)

func main() {
	loadEnv()
	app := &cli.App{
		Name:  "spotify-automations",
		Usage: "A simple CLI program for running different spotify automations",
		Action: func(c *cli.Context) error {
			runConfiguration()
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "start",
				Usage:   "Start the automation process",
				Aliases: []string{"s"},
				Action: func(c *cli.Context) error {
					automation.Run()
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func runConfiguration() {
	for {
		option := start.NewStartCommand()
		switch option {
		case start.Start:
			automation.Run()
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
