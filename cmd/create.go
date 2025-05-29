package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"mineflared-cli/internal"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func downloadFile(url, dest string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Cria um novo servidor Minecraft",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := internal.GetConfig()
		if cfg.Token == "" {
			fmt.Println("Você precisa estar logado. Use 'mineflared-cli login'")
			return
		}

		// 1. Perguntar nome do servidor
		var serverName string
		fmt.Print("Digite o nome do servidor: ")
		fmt.Scanln(&serverName)
		serverName = strings.TrimSpace(serverName)
		if serverName == "" {
			fmt.Println("Nome do servidor não pode ser vazio!")
			return
		}

		// 2. Montar caminho da pasta do servidor
		usr, _ := user.Current()
		serverDir := filepath.Join(usr.HomeDir, "mineflared", serverName)
		if err := os.MkdirAll(serverDir, 0755); err != nil {
			fmt.Println("Erro ao criar diretório do servidor:", err)
			return
		}

		// 3. Escolher tipo do servidor
		fmt.Println("Escolha o tipo de servidor:")
		fmt.Println("[1] Java (Paper)")
		fmt.Println("[2] Bedrock")
		var option int
		fmt.Print("Opção: ")
		fmt.Scanln(&option)

		var tipo, serverURL, fileName string
		switch option {
		case 1:
			tipo = "java"
			serverURL = "https://api.papermc.io/v2/projects/paper/versions/1.21.4/builds/231/downloads/paper-1.21.4-231.jar"
			fileName = "paper-1.21.4-231.jar"
		case 2:
			tipo = "bedrock"
			serverURL = "https://www.minecraft.net/bedrockdedicatedserver/bin-win/bedrock-server-1.21.83.1.zip"
			fileName = "bedrock-server-1.21.83.1.zip"
		default:
			fmt.Println("Opção inválida.")
			return
		}

		// 4. Baixar arquivo do servidor para a pasta
		filePath := filepath.Join(serverDir, fileName)
		fmt.Println("Baixando arquivo do servidor...")
		if err := downloadFile(serverURL, filePath); err != nil {
			fmt.Println("Erro ao baixar o arquivo:", err)
			return
		}
		fmt.Println("Arquivo baixado em:", filePath)

		// 5. Perguntar links mods/plugins (opcional)
		fmt.Println("Insira links para mods/plugins (separados por espaço, ou pressione Enter para nenhum):")
		var links string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			links = strings.TrimSpace(scanner.Text())
		}

		// 6. Enviar dados para API (como antes)
		body := map[string]string{
			"type":       tipo,
			"mods":       links,
			"username":   cfg.Username,
			"serverUrl":  serverURL,
			"serverName": serverName,
		}
		jsonBody, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", "https://mineflared.theushen.me/api/server/create", io.NopCloser(bytes.NewReader(jsonBody)))
		req.Header.Set("Authorization", "Bearer "+cfg.Token)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil || resp.StatusCode != 200 {
			fmt.Println("Erro ao criar servidor")
			return
		}

		fmt.Println("Servidor criado! Aguarde a propagação do DNS...")
		fmt.Printf("Acesse: https://%s.mineserver.theushen.me\n", cfg.Username)
	},
}
