package cmd

import (
	"archive/zip"
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"mineflared-cli/internal"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func unzipFile(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()
	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, f.Mode())
			continue
		}
		if err := os.MkdirAll(filepath.Dir(fpath), 0755); err != nil {
			return err
		}
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return err
		}
		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: internal.GetTranslation("RESTORE_SHORT_DESC"),
	Run: func(cmd *cobra.Command, args []string) {
		// Navegação simples
		fmt.Println(internal.GetTranslation("RESTORE_CHOOSE_PATH"))
		cwd, _ := os.Getwd()
		files, _ := os.ReadDir(cwd)
		fmt.Println(internal.GetTranslation("RESTORE_CURRENT_DIR"), cwd)
		for _, f := range files {
			if strings.HasSuffix(f.Name(), ".zip") || strings.HasSuffix(f.Name(), ".rar") {
				fmt.Println("- " + f.Name())
			}
		}
		fmt.Print(internal.GetTranslation("RESTORE_FILE_PROMPT"))
		reader := bufio.NewReader(os.Stdin)
		backupFile, _ := reader.ReadString('\n')
		backupFile = strings.TrimSpace(backupFile)
		if backupFile == "" {
			fmt.Println(internal.GetTranslation("RESTORE_FILE_EMPTY"))
			return
		}
		if _, err := os.Stat(backupFile); err != nil {
			fmt.Println(internal.GetTranslation("RESTORE_FILE_NOT_FOUND"))
			return
		}

		// Nome do servidor será o nome do arquivo sem _backup e extensão
		base := filepath.Base(backupFile)
		serverName := base
		serverName = strings.ReplaceAll(serverName, "_backup.zip", "")
		serverName = strings.ReplaceAll(serverName, "_backup.rar", "")
		serverName = strings.TrimSuffix(serverName, ".zip")
		serverName = strings.TrimSuffix(serverName, ".rar")
		serverDir := filepath.Join(string(os.PathSeparator), "servers", serverName)

		if _, err := os.Stat(serverDir); err == nil {
			fmt.Println(internal.GetTranslation("RESTORE_SERVER_EXISTS"))
			return
		}
		os.MkdirAll(serverDir, 0755)

		if strings.HasSuffix(backupFile, ".zip") {
			fmt.Println(internal.GetTranslation("RESTORE_UNZIPPING"))
			if err := unzipFile(backupFile, serverDir); err != nil {
				fmt.Printf(internal.GetTranslation("RESTORE_UNZIP_ERROR"), err)
				return
			}
			fmt.Printf(internal.GetTranslation("RESTORE_DONE"), serverDir)
			fmt.Println()
		} else if strings.HasSuffix(backupFile, ".rar") {
			fmt.Println(internal.GetTranslation("RESTORE_UNRAR"))
			cmd := exec.Command("unrar", "x", "-o+", backupFile, serverDir)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Printf(internal.GetTranslation("RESTORE_UNRAR_ERROR"), err)
				return
			}
			fmt.Printf(internal.GetTranslation("RESTORE_DONE"), serverDir)
			fmt.Println()
		} else {
			fmt.Println(internal.GetTranslation("RESTORE_INVALID_FORMAT"))
		}
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)
}
