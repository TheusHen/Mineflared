package main

import (
	"fmt"
	"mineflared-cli/cmd"
	"mineflared-cli/internal"
	"os"
	"path/filepath"
)

func main() {
	internal.SelfInstall()

	configDir, _ := os.UserConfigDir()
	mineDir := filepath.Join(configDir, "minecli")
	configPath := filepath.Join(mineDir, "config.json")

	firstRun := false
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		firstRun = true
	}

	internal.LoadConfig()

	if firstRun {
		selectLanguageOnFirstRun()
	}

	cmd.UpdateDNS()
	cmd.Execute()
}

func selectLanguageOnFirstRun() {
	fmt.Println(internal.GetTranslation("FIRST_RUN_LANGUAGE"))
	fmt.Println(internal.GetTranslation("FIRST_RUN_ENGLISH"))
	fmt.Println(internal.GetTranslation("FIRST_RUN_PORTUGUESE"))

	var option int
	fmt.Print(internal.GetTranslation("LANGUAGE_OPTION_PROMPT"))
	fmt.Scanln(&option)

	switch option {
	case 1:
		internal.SetLanguage("en")
	case 2:
		internal.SetLanguage("pt")
	default:
		internal.SetLanguage("pt")
	}
}
