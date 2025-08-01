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
	internal.LoadConfig()

	rootCmd.Short = internal.GetTranslation("CLI_SHORT_DESC")

	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(createCmd)

	rootCmd.Execute()
}
