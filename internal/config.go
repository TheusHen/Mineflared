package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	IP       string `json:"ip"`
	Language string `json:"language"`
}

var config Config

func GetConfig() *Config {
	return &config
}

func LoadConfig() {
	configDir, _ := os.UserConfigDir()
	mineDir := filepath.Join(configDir, "minecli")
	tokenPath := filepath.Join(mineDir, "config.json")

	if data, err := os.ReadFile(tokenPath); err == nil {
		json.Unmarshal(data, &config)
	}
}

func SaveConfig() {
	configDir, _ := os.UserConfigDir()
	mineDir := filepath.Join(configDir, "minecli")
	os.MkdirAll(mineDir, 0700)
	tokenPath := filepath.Join(mineDir, "config.json")
	data, _ := json.Marshal(config)
	os.WriteFile(tokenPath, data, 0600)
}
