package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"mineflared-cli/internal"
	"net/http"
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

		serverStatusURL := internal.GetEnv("SERVER_STATUS_URL", "http://localhost:3000/api/server/status")
		url := fmt.Sprintf("%s?username=%s", serverStatusURL, cfg.Username)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(internal.GetTranslation("STATUS_REQUEST_ERROR"), err)
			return
		}

		req.Header.Set("Authorization", "Bearer "+cfg.Token)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(internal.GetTranslation("STATUS_QUERY_ERROR"), err)
			return
		}
		defer resp.Body.Close()

		var result map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			fmt.Println(internal.GetTranslation("STATUS_DECODE_ERROR"), err)
			return
		}

		if resp.StatusCode != 200 {
			if msg, ok := result["message"]; ok {
				fmt.Printf(internal.GetTranslation("STATUS_API_ERROR"), msg)
				fmt.Println()
			} else if errMsg, ok := result["error"]; ok {
				fmt.Printf(internal.GetTranslation("STATUS_API_ERROR"), errMsg)
				fmt.Println()
			} else {
				fmt.Printf(internal.GetTranslation("STATUS_UNKNOWN_ERROR"), resp.StatusCode)
				fmt.Println()
			}
			return
		}

		fmt.Println(internal.GetTranslation("STATUS_SERVER_STATUS"))
		fmt.Printf(internal.GetTranslation("STATUS_STATUS_LINE"), result["status"])
		fmt.Println()
		if msg, ok := result["message"]; ok {
			fmt.Printf(internal.GetTranslation("STATUS_MESSAGE_LINE"), msg)
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
