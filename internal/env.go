package internal

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		execPath, execErr := os.Executable()
		if execErr == nil {
			execDir := filepath.Dir(execPath)
			envPath := filepath.Join(execDir, ".env")
			err = godotenv.Load(envPath)
		}

		if err != nil {
			fmt.Println("Warning: Could not load .env file")
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

// IsTermux detects if the program is running in Termux environment
func IsTermux() bool {
	prefix := os.Getenv("PREFIX")
	if prefix == "" {
		return false
	}
	
	// Termux's PREFIX typically looks like /data/data/com.termux/files/usr
	// Check if the path contains "com.termux"
	absPrefix, err := filepath.Abs(prefix)
	if err != nil {
		return false
	}
	
	return filepath.Base(absPrefix) == "usr" && 
		(filepath.Dir(absPrefix) == "/data/data/com.termux/files" ||
		 filepath.Base(filepath.Dir(absPrefix)) == "com.termux")
}
