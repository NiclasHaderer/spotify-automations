package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"spotify-automations/internal/models"
)

var Path = os.Getenv("HOME") + "/.spotify-automations/config.json"
var Instance = getOrCreate()

func Save() {
	dir := filepath.Dir(Path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		_ = os.MkdirAll(dir, os.ModePerm)
	}

	data, _ := json.Marshal(Instance)

	_ = os.WriteFile(Path, data, 0644)
}

func getOrCreate() models.Config {
	if _, err := os.Stat(Path); errors.Is(err, os.ErrNotExist) {
		Instance = empty()
		Save()
	}

	return load()
}

func empty() models.Config {
	return models.Config{
		Automations: []models.Automation{},
		User:        nil,
	}
}

func load() models.Config {
	data, _ := os.ReadFile(Path)

	var model models.Config
	_ = json.Unmarshal(data, &model)

	return model
}
