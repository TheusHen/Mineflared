package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"mineflared-cli/internal"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var startCmd = &cobra.Command{
	Use:   "start [server-name]",
	Short: internal.GetTranslation("START_SHORT_DESC"),
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := internal.GetConfig()
		if cfg.Token == "" {
			fmt.Println(internal.GetTranslation("START_LOGIN_REQUIRED"))
			return
		}

		serverName := args[0]
		serverDir := filepath.Join(string(os.PathSeparator), "servers", serverName)
		lockFile := filepath.Join(serverDir, ".running")
		configLock := filepath.Join(serverDir, ".configuring")

		if _, err := os.Stat(configLock); err == nil {
			fmt.Println(internal.GetTranslation("START_CONFIGURING_BLOCK"))
			return
		}
		os.WriteFile(lockFile, []byte("1"), 0644)
		defer os.Remove(lockFile)

		isJava := false
		isBedrock := false
		var jarFile string

		files, err := os.ReadDir(serverDir)
		if err != nil {
			fmt.Println(internal.GetTranslation("START_SERVER_DIR_ERROR"), err)
			return
		}
		for _, f := range files {
			if strings.HasSuffix(f.Name(), ".jar") {
				isJava = true
				jarFile = f.Name()
				break
			}
			if strings.HasPrefix(f.Name(), "bedrock-server") && f.IsDir() {
				isBedrock = true
			}
		}

		if !isJava && !isBedrock {
			fmt.Println(internal.GetTranslation("START_SERVER_TYPE_UNKNOWN"))
			return
		}

		if isJava {
			javaOK, javaVersion := checkJava21()
			if !javaOK {
				fmt.Println()
				fmt.Println("=====================================")
				fmt.Println("[!!!] JAVA 21 IS REQUIRED TO RUN THIS SERVER [!!!]")
				fmt.Println("Detected Java version:", javaVersion)
				fmt.Println()
				fmt.Println("Please install Java 21 or newer before running this server.")
				fmt.Println("Download Java 21 (Adoptium): https://adoptium.net/temurin/releases/?version=21")
				fmt.Println("After installing, close and reopen your terminal/Prompt/PowerShell.")
				fmt.Println("Then run this command again:")
				fmt.Printf("mineflared start %s\n", serverName)
				fmt.Println("=====================================")
				return
			}

			eulaPath := filepath.Join(serverDir, "eula.txt")
			needAskEULA := true
			if _, err := os.Stat(eulaPath); err == nil {
				eulaFile, err := os.ReadFile(eulaPath)
				if err == nil && strings.Contains(string(eulaFile), "eula=true") {
					needAskEULA = false
				}
			}
			if needAskEULA {
				fmt.Print(internal.GetTranslation("START_EULA_PROMPT"))
				reader := bufio.NewReader(os.Stdin)
				answer, _ := reader.ReadString('\n')
				answer = strings.TrimSpace(strings.ToLower(answer))
				if answer == "y" || answer == "yes" || answer == "s" || answer == "sim" {
					f, err := os.Create(eulaPath)
					if err != nil {
						fmt.Println(internal.GetTranslation("START_EULA_WRITE_ERROR"), err)
						return
					}
					f.WriteString("eula=true\n")
					f.Close()
					fmt.Println(internal.GetTranslation("START_EULA_ACCEPTED"))
				} else {
					fmt.Println(internal.GetTranslation("START_EULA_DECLINED"))
					return
				}
			}
			fmt.Println(internal.GetTranslation("START_JAVA_STARTING"))
			runCmd := exec.Command("java", "-Xmx1024M", "-Xms1024M", "-jar", jarFile, "nogui")
			runCmd.Dir = serverDir
			runCmd.Stdout = os.Stdout
			runCmd.Stderr = os.Stderr
			err := runCmd.Run()
			if err != nil {
				fmt.Printf(internal.GetTranslation("START_JAVA_ERROR"), err)
				fmt.Println()
			}
		} else if isBedrock {
			bedrockDir := ""
			for _, f := range files {
				if strings.HasPrefix(f.Name(), "bedrock-server") && f.IsDir() {
					bedrockDir = filepath.Join(serverDir, f.Name())
					break
				}
			}
			if bedrockDir == "" {
				fmt.Println(internal.GetTranslation("START_BEDROCK_NOT_FOUND"))
				return
			}

			bedrockBinary := filepath.Join(bedrockDir, "bedrock_server")
			if _, err := os.Stat(bedrockBinary); err != nil {
				bedrockBinary += ".exe"
			}
			fmt.Println(internal.GetTranslation("START_BEDROCK_STARTING"))
			runCmd := exec.Command(bedrockBinary)
			runCmd.Dir = bedrockDir
			runCmd.Stdout = os.Stdout
			runCmd.Stderr = os.Stderr
			err := runCmd.Run()
			if err != nil {
				fmt.Printf(internal.GetTranslation("START_BEDROCK_ERROR"), err)
				fmt.Println()
			}
		}
	},
}

func checkJava21() (bool, string) {
	cmd := exec.Command("java", "-version")
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	if err := cmd.Run(); err != nil {
		return false, "not found"
	}
	ver := errb.String()
	version := parseJavaVersion(ver)
	if version == "" {
		version = strings.TrimSpace(ver)
	}
	re := regexp.MustCompile(`\b(2[1-9]|[3-9][0-9])(\.|")`)
	if re.FindString(version) != "" {
		return true, version
	}
	return false, version
}

func parseJavaVersion(ver string) string {
	re := regexp.MustCompile(`version "(.*?)"`)
	match := re.FindStringSubmatch(ver)
	if len(match) > 1 {
		return match[1]
	}
	re2 := regexp.MustCompile(`openjdk version "(.*?)"`)
	match = re2.FindStringSubmatch(ver)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func init() {
	rootCmd.AddCommand(startCmd)
}
