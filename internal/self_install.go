package internal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func SelfInstall() {
	exePath, err := os.Executable()
	if err != nil {
		return
	}

	if strings.Contains(exePath, os.TempDir()) {
		return
	}

	switch runtime.GOOS {
	case "windows":
		programFiles := os.Getenv("ProgramFiles")
		if programFiles == "" {
			return
		}
		targetDir := filepath.Join(programFiles, "mineflared-cli")
		targetExe := filepath.Join(targetDir, "mineflared-cli.exe")

		if isSameFile(exePath, targetExe) {
			return
		}

		if err := os.MkdirAll(targetDir, 0755); err != nil {
			fmt.Println("❌ Permission denied to create install directory.")
			fmt.Println("Please run the program as administrator and try again.")
			os.Exit(1)
		}

		err = copyFileWithErr(exePath, targetExe)
		if err != nil {
			fmt.Println("❌ Permission denied to copy executable.")
			fmt.Println("Please run the program as administrator and try again.")
			os.Exit(1)
		}

		err = addToPathWindows(targetDir)
		if err != nil {
			fmt.Println("⚠️ Warning: Could not add directory to system PATH.")
			fmt.Println("You will need to add it manually:", targetDir)
		}

		fmt.Println("✅ mineflared-cli installed in", targetDir)
		fmt.Println("Please open a new terminal and run: mineflared-cli")

		cmd := exec.Command(targetExe, os.Args[1:]...)
		cmd.Start()
		os.Exit(0)

	case "linux":
		targetDir := "/usr/local/bin"
		targetExe := filepath.Join(targetDir, "mineflared-cli")

		if isSameFile(exePath, targetExe) {
			return
		}

		if err := os.MkdirAll(targetDir, 0755); err != nil {
			fmt.Println("❌ Permission denied to create install directory.")
			fmt.Println("Please run the program with sudo and try again.")
			os.Exit(1)
		}

		err = copyFileWithErr(exePath, targetExe)
		if err != nil {
			fmt.Println("❌ Permission denied to copy executable.")
			fmt.Println("Please run the program with sudo and try again.")
			os.Exit(1)
		}

		os.Chmod(targetExe, 0755)

		err = addToPathLinux(targetDir)
		if err != nil {
			fmt.Println("⚠️ Warning: Could not add directory to PATH.")
			fmt.Println("You will need to add it manually:", targetDir)
		}

		fmt.Println("✅ mineflared-cli installed in", targetDir)
		fmt.Println("Please open a new terminal and run: mineflared-cli")

		cmd := exec.Command(targetExe, os.Args[1:]...)
		cmd.Start()
		os.Exit(0)
	}
}

func isSameFile(a, b string) bool {
	absA, _ := filepath.Abs(a)
	absB, _ := filepath.Abs(b)
	return absA == absB
}

func copyFileWithErr(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.ReadFrom(in)
	if err != nil {
		return err
	}

	return out.Sync()
}

func addToPathWindows(dir string) error {
	psCmd := fmt.Sprintf(`$p=[Environment]::GetEnvironmentVariable("Path", "Machine"); if ($p -notlike "*%s*") {[Environment]::SetEnvironmentVariable("Path", "$p;%s", "Machine")}`, dir, dir)
	cmd := exec.Command("powershell", "-Command", psCmd)
	return cmd.Run()
}

func addToPathLinux(dir string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	profile := filepath.Join(home, ".profile")
	data, err := os.ReadFile(profile)
	if err != nil {
		return err
	}
	if !strings.Contains(string(data), dir) {
		f, err := os.OpenFile(profile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.WriteString("\nexport PATH=\"$PATH:" + dir + "\"\n")
		if err != nil {
			return err
		}
	}
	return nil
}
