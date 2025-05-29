package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"mineflared-cli/internal"
	"net/http"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Faz login via GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Abrindo navegador para login...")
		browser.OpenURL("https://mineflared.theushen.me/auth/github")
		fmt.Println("Após login, cole aqui o token retornado na URL:")
		var token string
		fmt.Print("Token: ")
		fmt.Scanln(&token)
		cfg := internal.GetConfig()
		cfg.Token = token
		// (Opcional) Requisição para obter username
		req, _ := http.NewRequest("GET", "https://mineflared.theushen.me/api/user", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		if err == nil && resp.StatusCode == 200 {
			var data map[string]string
			json.NewDecoder(resp.Body).Decode(&data)
			cfg.Username = data["username"]
		}
		internal.SaveConfig()
		fmt.Println("Login realizado com sucesso!")
	},
}
