package internal

import (
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		execPath, err := os.Executable()
		if err == nil {
			execDir := filepath.Dir(execPath)
			envPath := filepath.Join(execDir, ".env")
			_ = godotenv.Load(envPath)
		}
	}
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}