package config

import (
	"encoding/json"
	"errors"
	"os"
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
		Automations: map[string]any{},
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
