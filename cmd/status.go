package cmd

import (
	"fmt"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"mineflared-cli/internal"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: internal.GetTranslation("STATUS_SHORT_DESC"),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := internal.GetConfig()
		if cfg.Username == "" {
			fmt.Println(internal.GetTranslation("STATUS_LOGIN_REQUIRED"))
			return
		}

		url := fmt.Sprintf("https://mineflared.theushen.me/status?username=%s", cfg.Username)
		fmt.Println(internal.GetTranslation("STATUS_OPENING_BROWSER"))
		err := browser.OpenURL(url)
		if err != nil {
			fmt.Println(internal.GetTranslation("STATUS_BROWSER_ERROR"), err)
			fmt.Printf("URL: %s\n", url)
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
