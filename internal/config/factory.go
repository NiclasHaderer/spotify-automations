package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"spotify-automations/internal/models"
	"spotify-automations/internal/textarea"
)

func Get() models.Config {
	return getOrCreate()
}

func getOrCreate() models.Config {
	path := getPath()
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		Save(empty())
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
	path := getPath()
	data, _ := os.ReadFile(path)

	var model models.Config
	_ = json.Unmarshal(data, &model)

	return model
}

func Save(instance models.Config) {
	path := getPath()
	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		_ = os.MkdirAll(dir, os.ModePerm)
	}

	data, _ := json.Marshal(instance)

	_ = os.WriteFile(path, data, 0644)
}

func Print(instance models.Config) {
	data, _ := json.MarshalIndent(instance, "", "  ")
	textarea.New("Config", string(data), true)
}
