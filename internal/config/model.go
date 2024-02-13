package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"spotify-automations/internal/models"
	"spotify-automations/internal/textarea"
)

type Config struct {
	Automations []models.Automation[any] `json:"automations"`
	User        *models.User             `json:"user"`
}

func Get() *Config {
	return getOrCreate()
}

func (c *Config) Save() {
	path := getPath()
	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		_ = os.MkdirAll(dir, os.ModePerm)
	}

	data, _ := json.Marshal(c)
	_ = os.WriteFile(path, data, 0644)
}

func (c *Config) Print() {
	data, _ := json.MarshalIndent(c, "", "  ")
	textarea.New("Config", string(data), true)
}
