package internal

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

func LoadEnv() {
	// Tenta carregar o .env no diretório atual
	err := godotenv.Load()
	if err != nil {
		// Se falhar, tenta carregar no diretório do executável
		execPath, execErr := os.Executable()
		if execErr == nil {
			execDir := filepath.Dir(execPath)
			envPath := filepath.Join(execDir, ".env")
			err = godotenv.Load(envPath)
		}

		// Se ainda falhar, exibe uma mensagem de erro
		if err != nil {
			fmt.Println("Aviso: Não foi possível carregar o arquivo .env")
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
