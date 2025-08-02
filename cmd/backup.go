package cmd

import (
	"archive/zip"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"mineflared-cli/internal"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func zipDir(source, target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name, _ = filepath.Rel(filepath.Dir(source), path)
		if info.IsDir() {
			header.Name += string(os.PathSeparator)
		} else {
			header.Method = zip.Deflate
		}
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return nil
}

var backupCmd = &cobra.Command{
	Use:   "backup [server-name]",
	Short: internal.GetTranslation("BACKUP_SHORT_DESC"),
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		serverName := args[0]
		serverDir := filepath.Join(string(os.PathSeparator), "servers", serverName)
		if _, err := os.Stat(serverDir); os.IsNotExist(err) {
			fmt.Println(internal.GetTranslation("BACKUP_SERVER_NOT_FOUND"))
			return
		}

		fmt.Print(internal.GetTranslation("BACKUP_FORMAT_PROMPT"))
		var format string
		fmt.Scanln(&format)
		format = strings.ToLower(strings.TrimSpace(format))

		if format != "zip" && format != "rar" {
			fmt.Println(internal.GetTranslation("BACKUP_INVALID_FORMAT"))
			return
		}

		backupFileName := fmt.Sprintf("%s_backup.%s", serverName, format)
		fmt.Println(internal.GetTranslation("BACKUP_CHOOSE_PATH"))
		cwd, _ := os.Getwd()
		dirs, _ := os.ReadDir(cwd)
		fmt.Println(internal.GetTranslation("BACKUP_CURRENT_DIR"), cwd)
		for _, d := range dirs {
			if d.IsDir() {
				fmt.Println("- " + d.Name())
			}
		}
		fmt.Print(internal.GetTranslation("BACKUP_PATH_PROMPT"))
		var savePath string
		fmt.Scanln(&savePath)
		savePath = strings.TrimSpace(savePath)
		if savePath == "" {
			savePath = cwd
		}

		backupFullPath := filepath.Join(savePath, backupFileName)

		if format == "zip" {
			fmt.Println(internal.GetTranslation("BACKUP_CREATING_ZIP"))
			if err := zipDir(serverDir, backupFullPath); err != nil {
				fmt.Printf(internal.GetTranslation("BACKUP_ZIP_ERROR"), err)
				return
			}
			fmt.Printf(internal.GetTranslation("BACKUP_DONE"), backupFullPath)
		} else if format == "rar" {
			fmt.Println(internal.GetTranslation("BACKUP_CREATING_RAR"))
			cmd := exec.Command("rar", "a", "-r", backupFullPath, serverDir)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Printf(internal.GetTranslation("BACKUP_RAR_ERROR"), err)
				return
			}
			fmt.Printf(internal.GetTranslation("BACKUP_DONE"), backupFullPath)
		}
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
}
