package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"mineflared-cli/internal"
	"net/http"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: internal.GetTranslation("DELETE_SHORT_DESC"),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := internal.GetConfig()
		if cfg.Token == "" {
			fmt.Println(internal.GetTranslation("DELETE_LOGIN_REQUIRED"))
			return
		}

		backendURL := internal.GetEnv("BACKEND_URL", "http://localhost:3000")
		deleteURL := backendURL + "/delete"

		fmt.Print(internal.GetTranslation("DELETE_CONFIRM_PROMPT"))
		var confirm string
		fmt.Scanln(&confirm)
		if confirm != "y" && confirm != "Y" {
			fmt.Println(internal.GetTranslation("DELETE_ABORTED"))
			return
		}

		req, err := http.NewRequest("DELETE", deleteURL, nil)
		if err != nil {
			fmt.Println(internal.GetTranslation("DELETE_REQUEST_FAILED"), err)
			return
		}
		req.Header.Set("Authorization", "Bearer "+cfg.Token)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(internal.GetTranslation("DELETE_REQUEST_FAILED"), err)
			return
		}
		defer resp.Body.Close()

		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)

		if resp.StatusCode == 200 && result["success"] == true {
			fmt.Println(internal.GetTranslation("DELETE_SUCCESS"))
			_ = internal.DeleteConfigFile()
		} else {
			errorMsg := result["error"]
			if errorMsg == nil {
				errorMsg = "Unknown error"
			}
			fmt.Printf("%s %v\n", internal.GetTranslation("DELETE_ERROR"), errorMsg)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
