package config

import (
	"os"
	"path/filepath"
	"spotify-automations/internal/textarea"
)

func getPath() string {
	// Check if the SPOTIFY_AUTOMATIONS_CONFIG_FILE_LOCATION is set
	var p string
	if os.Getenv("SPOTIFY_AUTOMATIONS_CONFIG_FILE_LOCATION") != "" {
		p = os.Getenv("SPOTIFY_AUTOMATIONS_CONFIG_FILE_LOCATION")
	} else {
		p = os.Getenv("HOME") + "/.config/spotify-automations/config.json"
	}

	p, _ = filepath.Abs(p)
	return p
}

func PrintPath() {
	textarea.New("Config Path", getPath(), false)
}
