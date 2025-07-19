package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"mineflared-cli/internal"
)

var languageCmd = &cobra.Command{
	Use:   "language",
	Short: "Altera o idioma da interface",
	Run: func(cmd *cobra.Command, args []string) {
		
		fmt.Printf(internal.GetTranslation("LANGUAGE_CURRENT"), internal.GetLanguageName())
		fmt.Println()
		
		fmt.Println(internal.GetTranslation("LANGUAGE_PROMPT"))
		var option int
		fmt.Print(internal.GetTranslation("LANGUAGE_OPTION_PROMPT"))
		fmt.Scanln(&option)
		
		switch option {
		case 1:
			internal.SetLanguage("en")
			fmt.Printf(internal.GetTranslation("LANGUAGE_CHANGED"), "English")
			fmt.Println()
		case 2:
			internal.SetLanguage("pt")
			fmt.Printf(internal.GetTranslation("LANGUAGE_CHANGED"), "Português")
			fmt.Println()
		default:
			fmt.Println(internal.GetTranslation("LANGUAGE_INVALID_OPTION"))
		}
	},
}

func init() {
	rootCmd.AddCommand(languageCmd)
}