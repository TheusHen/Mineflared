package cmd

import (
	"github.com/spf13/cobra"
	"mineflared-cli/internal"
)

var rootCmd = &cobra.Command{
	Use:   "mineflared-cli",
	Short: "CLI para criar e gerenciar servidores de Minecraft",
}

func Execute() {
	internal.InitConfig()
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.Execute()
}
