package cmd

import (
	"fmt"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"mineflared-cli/internal"
	"net"
	"net/http"
	"os"
	"encoding/json"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: internal.GetTranslation("LOGIN_SHORT_DESC"),
	Run: func(cmd *cobra.Command, args []string) {
		internal.LoadConfig()
		ip := internal.GetIP()
		if ip == "" {
			fmt.Println(internal.GetTranslation("LOGIN_IP_ERROR"))
			return
		}

		backendURL := internal.GetEnv("BACKEND_URL", "http://localhost:3000")
		callbackURL := internal.GetEnv("CALLBACK_URL", "http://localhost:4500/callback")
		authURL := fmt.Sprintf("%s/auth/github/login?ip=%s&callback=%s", backendURL, ip, callbackURL)

		fmt.Println(internal.GetTranslation("LOGIN_BROWSER_OPENING"))
		browser.OpenURL(authURL)

		token := waitForTokenViaLocalhost()
		if token == "" {
			fmt.Println(internal.GetTranslation("LOGIN_TOKEN_ERROR"))
			return
		}

		cfg := internal.GetConfig()
		cfg.Token = token

		req, _ := http.NewRequest("GET", backendURL+"/api/user", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		if err == nil && resp.StatusCode == 200 {
			var data map[string]string
			json.NewDecoder(resp.Body).Decode(&data)
			cfg.Username = data["username"]
		}

		internal.SaveConfig()
		fmt.Println(internal.GetTranslation("LOGIN_SUCCESS"))
	},
}

func waitForTokenViaLocalhost() string {
	listener, err := net.Listen("tcp", "localhost:4500")
	if err != nil {
		fmt.Println(internal.GetTranslation("LOGIN_PORT_ERROR"), err)
		return ""
	}
	defer listener.Close()

	tokenChan := make(chan string)

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		if token == "" {
			http.Error(w, internal.GetTranslation("LOGIN_TOKEN_NOT_RECEIVED"), http.StatusBadRequest)
			tokenChan <- ""
			return
		}
		fmt.Fprintf(w, internal.GetTranslation("LOGIN_BROWSER_SUCCESS"))
		tokenChan <- token
	})

	go http.Serve(listener, nil)
	fmt.Println(internal.GetTranslation("LOGIN_WAITING"))

	token := <-tokenChan
	return token
}
