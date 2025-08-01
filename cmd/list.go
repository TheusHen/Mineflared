package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"mineflared-cli/internal"
	"os"
	"path/filepath"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: internal.GetTranslation("LIST_SHORT_DESC"),
	Run: func(cmd *cobra.Command, args []string) {
		serversDir := filepath.Join(string(os.PathSeparator), "servers")
		files, err := os.ReadDir(serversDir)
		if err != nil {
			fmt.Println(internal.GetTranslation("LIST_SERVERS_DIR_ERROR"), err)
			return
		}
		cfg := internal.GetConfig()
		fmt.Println(internal.GetTranslation("LIST_HEADER"))
		for _, f := range files {
			if f.IsDir() {
				fmt.Printf("• %s  |  %s\n", f.Name(), cfg.Username)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
