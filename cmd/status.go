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
	Short: "Verifica o status do servidor",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := internal.GetConfig()
		if cfg.Username == "" {
			fmt.Println("Você precisa estar logado primeiro!")
			return
		}
		url := fmt.Sprintf("https://mineflared.theushen.me/api/server/status?username=%s", cfg.Username)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Bearer "+cfg.Token)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("Erro ao consultar status")
			return
		}
		var status map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&status)
		fmt.Println("Status do servidor:")
		fmt.Println(status)
	},
}
