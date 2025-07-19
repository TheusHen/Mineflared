package cmd

import (
	"archive/zip"
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

func validateFileExt(file string) bool {
	ext := strings.ToLower(filepath.Ext(file))
	switch ext {
	case ".jar", ".zip", ".mcaddon", ".mcpack":
		return true
	default:
		return false
	}
}

func moveFiles(files []string, dstDir string) error {
	for _, file := range files {
		if !validateFileExt(file) {
			fmt.Printf(internal.GetTranslation("INVALID_FILE"), file)
			fmt.Println()
			continue
		}
		base := filepath.Base(file)
		dst := filepath.Join(dstDir, base)
		if err := os.Rename(file, dst); err != nil {
			return err
		}
		fmt.Printf(internal.GetTranslation("FILE_MOVED"), file, dst)
		fmt.Println()
	}
	return nil
}
func unzip(src, destDir string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()
	for _, f := range r.File {
		fpath := filepath.Join(destDir, f.Name)
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

type ServerOption struct {
	Name           string
	URL            string
	SupportsMods   bool
	SupportsPlugin bool
	FileName       string
}

var javaOptions = []ServerOption{
	{
		Name:           "PaperMC 1.21.4",
		URL:            "https://api.papermc.io/v2/projects/paper/versions/1.21.4/builds/231/downloads/paper-1.21.4-231.jar",
		SupportsMods:   false,
		SupportsPlugin: true,
		FileName:       "paper-1.21.4-231.jar",
	},
	{
		Name:           "Vanilla 1.21.1",
		URL:            "https://piston-data.mojang.com/v1/objects/6bce4ef400e4efaa63a13d5e6f6b500be969ef81/server.jar",
		SupportsMods:   false,
		SupportsPlugin: false,
		FileName:       "vanilla-1.21.1.jar",
	},
	{
		Name:           "Purpur 1.21.8",
		URL:            "https://api.purpurmc.org/v2/purpur/1.21.8/2478/download",
		SupportsMods:   false,
		SupportsPlugin: true,
		FileName:       "purpur-1.21.8-2478.jar",
	},
	// You can add more options here.
}

var bedrockOption = ServerOption{
	Name:           "Bedrock 1.21.95.1",
	URL:            "https://www.minecraft.net/content/dam/minecraft/bedrock-server/bedrock-server-1.21.95.1.zip",
	SupportsMods:   true,
	SupportsPlugin: false,
	FileName:       "bedrock-server-1.21.95.1.zip",
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: internal.GetTranslation("CREATE_SHORT_DESC"),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := internal.GetConfig()
		if cfg.Token == "" {
			fmt.Println(internal.GetTranslation("CREATE_LOGIN_REQUIRED"))
			return
		}

		var serverName string
		fmt.Print(internal.GetTranslation("CREATE_SERVER_NAME_PROMPT"))
		fmt.Scanln(&serverName)
		serverName = strings.TrimSpace(serverName)
		if serverName == "" {
			fmt.Println(internal.GetTranslation("CREATE_SERVER_NAME_EMPTY"))
			return
		}

		usr, _ := user.Current()
		serverDir := filepath.Join(usr.HomeDir, "mineflared", serverName)
		if err := os.MkdirAll(serverDir, 0755); err != nil {
			fmt.Println(internal.GetTranslation("CREATE_DIR_ERROR"), err)
			return
		}

		fmt.Println(internal.GetTranslation("CREATE_SERVER_TYPE_PROMPT"))
		fmt.Println(internal.GetTranslation("CREATE_JAVA_OPTION"))
		fmt.Println(internal.GetTranslation("CREATE_BEDROCK_OPTION"))
		var option int
		fmt.Print(internal.GetTranslation("CREATE_OPTION_PROMPT"))
		fmt.Scanln(&option)

		var serverChoice ServerOption
		var tipo string
		switch option {
		case 1:
			tipo = "java"
			fmt.Println(internal.GetTranslation("CREATE_JAVA_VERSION_PROMPT"))
			for i, opt := range javaOptions {
				fmt.Printf(internal.GetTranslation("CREATE_JAVA_VERSION_OPTION"), i+1, opt.Name, opt.SupportsMods, opt.SupportsPlugin)
				fmt.Println()
			}
			var javaIdx int
			fmt.Print(internal.GetTranslation("CREATE_OPTION_PROMPT"))
			fmt.Scanln(&javaIdx)
			if javaIdx < 1 || javaIdx > len(javaOptions) {
				fmt.Println(internal.GetTranslation("CREATE_INVALID_OPTION"))
				return
			}
			serverChoice = javaOptions[javaIdx-1]
		case 2:
			tipo = "bedrock"
			serverChoice = bedrockOption
			fmt.Println(internal.GetTranslation("CREATE_BEDROCK_WARNING"))
			fmt.Println(internal.GetTranslation("CREATE_BEDROCK_CANCEL"))
			var resp string
			fmt.Scanln(&resp)
			if strings.ToLower(resp) == "n" {
				fmt.Println(internal.GetTranslation("CREATE_ACTION_CANCELLED"))
				return
			}
		default:
			fmt.Println(internal.GetTranslation("CREATE_INVALID_OPTION"))
			return
		}

		filePath := filepath.Join(serverDir, serverChoice.FileName)
		fmt.Println(internal.GetTranslation("CREATE_DOWNLOADING"))
		if err := downloadFile(serverChoice.URL, filePath); err != nil {
			fmt.Println(internal.GetTranslation("CREATE_DOWNLOAD_ERROR"), err)
			return
		}
		fmt.Println(internal.GetTranslation("CREATE_DOWNLOAD_COMPLETE"), filePath)

		var extraFiles []string
		if tipo == "java" {
			if serverChoice.SupportsMods {
				fmt.Println(internal.GetTranslation("CREATE_MODS_SUPPORT"))
			} else if serverChoice.SupportsPlugin {
				fmt.Println(internal.GetTranslation("CREATE_PLUGINS_SUPPORT"))
			} else {
				fmt.Println(internal.GetTranslation("CREATE_NO_MODS_PLUGINS"))
			}
			if serverChoice.SupportsMods || serverChoice.SupportsPlugin {
				var links string
				scanner := bufio.NewScanner(os.Stdin)
				if scanner.Scan() {
					links = strings.TrimSpace(scanner.Text())
				}
				if links != "" {
					for _, link := range strings.Fields(links) {
						link = strings.TrimSpace(link)
						if link != "" {
							file := filepath.Join(serverDir, filepath.Base(link))
							fmt.Println(internal.GetTranslation("CREATE_DOWNLOADING_FILE"), link)
							if err := downloadFile(link, file); err != nil {
								fmt.Printf(internal.GetTranslation("CREATE_DOWNLOAD_FILE_ERROR"), link, err)
								fmt.Println()
								continue
							}
							extraFiles = append(extraFiles, file)
						}
					}
				} else {
					fmt.Println(internal.GetTranslation("CREATE_LOCAL_FILES_PROMPT"))
					fmt.Scanln()
					files, err := os.ReadDir(".")
					if err != nil {
						fmt.Println(internal.GetTranslation("CREATE_READ_DIR_ERROR"), err)
					}
					for _, f := range files {
						if !f.IsDir() && validateFileExt(f.Name()) {
							extraFiles = append(extraFiles, f.Name())
						}
					}
				}
				var dstDir string
				if serverChoice.SupportsMods {
					dstDir = filepath.Join(serverDir, "mods")
				} else if serverChoice.SupportsPlugin {
					dstDir = filepath.Join(serverDir, "plugins")
				}
				if dstDir != "" {
					os.MkdirAll(dstDir, 0755)
					for i := range extraFiles {
						extraFiles[i] = filepath.Join(".", extraFiles[i])
					}
					if err := moveFiles(extraFiles, dstDir); err != nil {
						fmt.Println(internal.GetTranslation("CREATE_MOVE_FILES_ERROR"), err)
					}
				}
			}
			fmt.Println(internal.GetTranslation("CREATE_JAVA_SERVER_STARTING"))
			runCmd := fmt.Sprintf("java -Xmx1024M -Xms1024M -jar %s nogui", serverChoice.FileName)
			fmt.Printf(internal.GetTranslation("CREATE_JAVA_SERVER_EXECUTE"), runCmd)
			fmt.Println()
		} else if tipo == "bedrock" {
			fmt.Println(internal.GetTranslation("CREATE_BEDROCK_EXTRACTING"))
			if err := unzip(filePath, serverDir); err != nil {
				fmt.Println(internal.GetTranslation("CREATE_BEDROCK_EXTRACT_ERROR"), err)
				return
			}
			fmt.Println(internal.GetTranslation("CREATE_BEDROCK_PACKS_PROMPT"))
			var links string
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				links = strings.TrimSpace(scanner.Text())
			}
			var packFiles []string
			if links != "" {
				for _, link := range strings.Fields(links) {
					link = strings.TrimSpace(link)
					if link != "" {
						file := filepath.Join(serverDir, filepath.Base(link))
						fmt.Println(internal.GetTranslation("CREATE_DOWNLOADING_FILE"), link)
						if err := downloadFile(link, file); err != nil {
							fmt.Printf(internal.GetTranslation("CREATE_DOWNLOAD_FILE_ERROR"), link, err)
							fmt.Println()
							continue
						}
						packFiles = append(packFiles, file)
					}
				}
			} else {
				fmt.Println(internal.GetTranslation("CREATE_BEDROCK_LOCAL_FILES_PROMPT"))
				fmt.Scanln()
				files, err := os.ReadDir(".")
				if err != nil {
					fmt.Println(internal.GetTranslation("CREATE_READ_DIR_ERROR"), err)
				}
				for _, f := range files {
					if !f.IsDir() && validateFileExt(f.Name()) {
						packFiles = append(packFiles, f.Name())
					}
				}
			}
			behaviorDir := filepath.Join(serverDir, "bedrock-server-1.21.95.1", "behavior_packs")
			resourceDir := filepath.Join(serverDir, "bedrock-server-1.21.95.1", "resource_packs")
			os.MkdirAll(behaviorDir, 0755)
			os.MkdirAll(resourceDir, 0755)
			for _, file := range packFiles {
				ext := strings.ToLower(filepath.Ext(file))
				dst := ""
				if ext == ".mcaddon" || ext == ".mcpack" {
					dst = resourceDir
				} else if ext == ".zip" {
					dst = behaviorDir
				}
				if dst != "" {
					if err := moveFiles([]string{file}, dst); err != nil {
						fmt.Println(internal.GetTranslation("CREATE_BEDROCK_MOVE_ERROR"), err)
					}
				}
			}
			fmt.Println(internal.GetTranslation("CREATE_BEDROCK_READY"))
		}

		body := map[string]string{
			"type":       tipo,
			"username":   cfg.Username,
			"serverUrl":  serverChoice.URL,
			"serverName": serverName,
		}
		jsonBody, _ := json.Marshal(body)
		serverCreateURL := internal.GetEnv("SERVER_CREATE_URL", "http://localhost:3000/api/server/create")
		req, _ := http.NewRequest("POST", serverCreateURL, io.NopCloser(bytes.NewReader(jsonBody)))
		req.Header.Set("Authorization", "Bearer "+cfg.Token)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil || resp.StatusCode != 200 {
			fmt.Println(internal.GetTranslation("CREATE_API_ERROR"))
			return
		}

		fmt.Println(internal.GetTranslation("CREATE_SERVER_CREATED"))
		fmt.Printf(internal.GetTranslation("CREATE_SERVER_ACCESS"), cfg.Username)
		fmt.Println()
	},
}
