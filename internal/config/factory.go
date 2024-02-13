package config

import (
	"encoding/json"
	"errors"
	"os"
	"spotify-automations/internal/models"
)

func getOrCreate() *Config {
	path := getPath()
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		empty().Save()
	}

	return load()
}

func empty() *Config {
	return &Config{
		Automations: map[string]models.Automation{},
		User:        nil,
	}
}

func load() *Config {
	path := getPath()
	data, _ := os.ReadFile(path)

	var model *Config
	_ = json.Unmarshal(data, &model)

	return model
}
