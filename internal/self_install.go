package internal

import (
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
		if !isSameFile(exePath, targetExe) {
			os.MkdirAll(targetDir, 0755)
			copyFile(exePath, targetExe)
			addToPathWindows(targetDir)
			exec.Command(targetExe, os.Args[1:]...).Start()
			os.Exit(0)
		}
	case "linux":
		targetDir := "/usr/local/bin"
		targetExe := filepath.Join(targetDir, "mineflared-cli")
		if !isSameFile(exePath, targetExe) {
			copyFile(exePath, targetExe)
			os.Chmod(targetExe, 0755)
			addToPathLinux(targetDir)
			exec.Command(targetExe, os.Args[1:]...).Start()
			os.Exit(0)
		}
	}
}

func isSameFile(a, b string) bool {
	absA, _ := filepath.Abs(a)
	absB, _ := filepath.Abs(b)
	return absA == absB
}

func copyFile(src, dst string) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer out.Close()
	_, _ = out.ReadFrom(in)
	out.Sync()
}

func addToPathWindows(dir string) {
	exec.Command("powershell", "-Command",
		`$p=[Environment]::GetEnvironmentVariable("Path", "Machine"); if ($p -notlike "*`+dir+`*") {[Environment]::SetEnvironmentVariable("Path", "$p;`+dir+`", "Machine")}`).Run()
}

func addToPathLinux(dir string) {
	home, _ := os.UserHomeDir()
	profile := filepath.Join(home, ".profile")
	data, _ := os.ReadFile(profile)
	if !strings.Contains(string(data), dir) {
		f, _ := os.OpenFile(profile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		defer f.Close()
		f.WriteString("\nexport PATH=\"$PATH:" + dir + "\"\n")
	}
}
