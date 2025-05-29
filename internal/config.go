package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}

var config Config
var configPath string

func InitConfig() {
	home, _ := os.UserHomeDir()
	configDir := filepath.Join(home, ".mineflared-cli")
	configPath = filepath.Join(configDir, "config.json")
	os.MkdirAll(configDir, os.ModePerm)

	file, err := os.ReadFile(configPath)
	if err == nil {
		json.Unmarshal(file, &config)
	}
}

func SaveConfig() {
	file, _ := json.MarshalIndent(config, "", "  ")
	os.WriteFile(configPath, file, 0644)
}

func GetConfig() *Config {
	return &config
}
