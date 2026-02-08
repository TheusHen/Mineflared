package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"mineflared-cli/internal"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

var configCmd = &cobra.Command{
	Use:   "config [server-name]",
	Short: internal.GetTranslation("CONFIG_SHORT_DESC"),
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		serverName := args[0]
		serverDir := filepath.Join(string(os.PathSeparator), "servers", serverName)
		lockFile := filepath.Join(serverDir, ".running")
		configLock := filepath.Join(serverDir, ".configuring")

		if _, err := os.Stat(lockFile); err == nil {
			fmt.Println(internal.GetTranslation("CONFIG_RUNNING_BLOCK"))
			return
		}

		os.WriteFile(configLock, []byte("1"), 0644)
		defer os.Remove(configLock)

		backendPath := filepath.Join("web", "backend", "main.go")
		cmdBackend := exec.Command("go", "run", backendPath, serverName)
		cmdBackend.Stdout = os.Stdout
		cmdBackend.Stderr = os.Stderr
		cmdBackend.Dir, _ = os.Getwd()

		go func() {
			time.Sleep(2 * time.Second)
			url := "http://localhost:3000/"
			fmt.Println(internal.GetTranslation("CONFIG_OPEN_BROWSER"), url)
			openBrowser(url)
		}()

		err := cmdBackend.Run()
		if err != nil {
			fmt.Println("Backend process exited with error:", err)
		}
	},
}

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	}
	if err != nil {
		fmt.Println("Please open this URL manually:", url)
	}
}

func init() {
	rootCmd.AddCommand(configCmd)
}
