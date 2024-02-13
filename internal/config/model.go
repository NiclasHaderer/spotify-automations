package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"spotify-automations/internal/models"
	"spotify-automations/internal/textarea"
)

type Config struct {
	Automations map[string]models.Automation `json:"automations"`
	User        *models.User                 `json:"user"`
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

func GetAutomationConfig[T any](name string) (v T, err error) {
	c := Get()

	if val, ok := c.Automations[name]; ok {
		return v, json.Unmarshal([]byte(val.Config), &v)
	}
	return v, errors.New("no configuration found")
}
