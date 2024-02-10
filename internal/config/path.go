package config

import (
	"fmt"
	"os"
)

func getPath() string {
	// Check if the SPOTIFY_AUTOMATIONS_CONFIG_FILE_LOCATION is set
	if os.Getenv("SPOTIFY_AUTOMATIONS_CONFIG_FILE_LOCATION") != "" {
		return os.Getenv("SPOTIFY_AUTOMATIONS_CONFIG_FILE_LOCATION")
	}
	return os.Getenv("HOME") + "/.config/spotify-automations/config.json"
}

func PrintPath() {
	fmt.Printf("Config path: %s", getPath())
}
