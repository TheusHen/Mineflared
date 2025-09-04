package internal

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// MinecraftServerValidator provides security validation for Minecraft servers
type MinecraftServerValidator struct {
	allowedJavaPorts    []int
	allowedBedrockPorts []int
	knownJarSignatures  map[string]bool
}

// NewMinecraftServerValidator creates a new validator instance
func NewMinecraftServerValidator() *MinecraftServerValidator {
	return &MinecraftServerValidator{
		allowedJavaPorts:    []int{25565, 25566, 25567, 25568, 25569}, // Standard Minecraft Java ports
		allowedBedrockPorts: []int{19132, 19133, 19134, 19135, 19136}, // Standard Minecraft Bedrock ports
		knownJarSignatures:  initKnownSignatures(),
	}
}

// initKnownSignatures initializes known Minecraft server jar signatures
func initKnownSignatures() map[string]bool {
	// This would contain known hashes of legitimate Minecraft server jars
	// For now, we'll use patterns to detect Minecraft-specific content
	return map[string]bool{}
}

// ValidateServerDirectory performs comprehensive validation of a server directory
func (v *MinecraftServerValidator) ValidateServerDirectory(serverDir string) error {
	// Check if directory exists
	if _, err := os.Stat(serverDir); os.IsNotExist(err) {
		return fmt.Errorf("server directory does not exist: %s", serverDir)
	}

	// Determine server type and validate accordingly
	serverType, err := v.detectServerType(serverDir)
	if err != nil {
		return fmt.Errorf("failed to detect server type: %w", err)
	}

	switch serverType {
	case "java":
		return v.validateJavaServer(serverDir)
	case "bedrock":
		return v.validateBedrockServer(serverDir)
	default:
		return fmt.Errorf("unknown or invalid server type detected")
	}
}

// detectServerType determines if this is a Java or Bedrock server
func (v *MinecraftServerValidator) detectServerType(serverDir string) (string, error) {
	files, err := os.ReadDir(serverDir)
	if err != nil {
		return "", err
	}

	hasJarFile := false
	hasBedrockDir := false

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".jar") {
			hasJarFile = true
		}
		if strings.HasPrefix(file.Name(), "bedrock-server") && file.IsDir() {
			hasBedrockDir = true
		}
	}

	if hasJarFile && hasBedrockDir {
		return "", fmt.Errorf("directory contains both Java (.jar) and Bedrock server files - this is suspicious")
	}

	if hasJarFile {
		return "java", nil
	}
	if hasBedrockDir {
		return "bedrock", nil
	}

	return "", fmt.Errorf("no valid Minecraft server files found")
}

// validateJavaServer validates a Java Minecraft server
func (v *MinecraftServerValidator) validateJavaServer(serverDir string) error {
	// Find the main jar file
	jarFile, err := v.findMainJarFile(serverDir)
	if err != nil {
		return fmt.Errorf("failed to find main jar file: %w", err)
	}

	// Validate the jar file
	if err := v.validateJarFile(jarFile); err != nil {
		return fmt.Errorf("jar file validation failed: %w", err)
	}

	// Check for suspicious files
	if err := v.checkForSuspiciousFiles(serverDir); err != nil {
		return fmt.Errorf("suspicious files detected: %w", err)
	}

	// Validate server.properties if it exists
	serverPropsPath := filepath.Join(serverDir, "server.properties")
	if _, err := os.Stat(serverPropsPath); err == nil {
		if err := v.validateServerProperties(serverPropsPath); err != nil {
			return fmt.Errorf("server.properties validation failed: %w", err)
		}
	}

	return nil
}

// validateBedrockServer validates a Bedrock Minecraft server
func (v *MinecraftServerValidator) validateBedrockServer(serverDir string) error {
	// Find bedrock server directory
	bedrockDir, err := v.findBedrockServerDir(serverDir)
	if err != nil {
		return fmt.Errorf("failed to find bedrock server directory: %w", err)
	}

	// Validate bedrock server structure
	if err := v.validateBedrockStructure(bedrockDir); err != nil {
		return fmt.Errorf("bedrock server structure validation failed: %w", err)
	}

	// Check for suspicious files
	if err := v.checkForSuspiciousFiles(serverDir); err != nil {
		return fmt.Errorf("suspicious files detected: %w", err)
	}

	return nil
}

// findMainJarFile finds the main server jar file
func (v *MinecraftServerValidator) findMainJarFile(serverDir string) (string, error) {
	files, err := os.ReadDir(serverDir)
	if err != nil {
		return "", err
	}

	var jarFiles []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".jar") && !file.IsDir() {
			jarFiles = append(jarFiles, file.Name())
		}
	}

	if len(jarFiles) == 0 {
		return "", fmt.Errorf("no jar files found")
	}

	// Prioritize known Minecraft server jar patterns
	for _, jar := range jarFiles {
		if v.isKnownMinecraftJar(jar) {
			return filepath.Join(serverDir, jar), nil
		}
	}

	// If no known patterns, check if there's only one jar file
	if len(jarFiles) == 1 {
		return filepath.Join(serverDir, jarFiles[0]), nil
	}

	return "", fmt.Errorf("multiple jar files found, cannot determine main server jar")
}

// isKnownMinecraftJar checks if a jar file name matches known Minecraft server patterns
func (v *MinecraftServerValidator) isKnownMinecraftJar(filename string) bool {
	lowerName := strings.ToLower(filename)
	
	// Exact matches
	exactMatches := []string{
		"server.jar",
		"minecraft_server.jar",
	}
	
	for _, exact := range exactMatches {
		if lowerName == exact {
			return true
		}
	}
	
	// Prefix patterns (must be at start)
	prefixPatterns := []string{
		"paper-",
		"spigot-",
		"bukkit-",
		"purpur-",
		"vanilla-",
		"minecraft_server.",
		"minecraft-server-",
		"forge-",
		"fabric-server-",
	}

	for _, pattern := range prefixPatterns {
		if strings.HasPrefix(lowerName, pattern) {
			return true
		}
	}
	
	return false
}

// validateJarFile validates that a jar file is a legitimate Minecraft server
func (v *MinecraftServerValidator) validateJarFile(jarPath string) error {
	// Open the jar file as a zip
	reader, err := zip.OpenReader(jarPath)
	if err != nil {
		return fmt.Errorf("failed to open jar file: %w", err)
	}
	defer reader.Close()

	// Check for Minecraft-specific files and classes
	hasMinecraftClasses := false
	suspiciousFiles := []string{}

	for _, file := range reader.File {
		filename := file.Name

		// Check for Minecraft-specific packages
		if strings.Contains(filename, "net/minecraft/") ||
			strings.Contains(filename, "com/mojang/") ||
			strings.Contains(filename, "org/bukkit/") ||
			strings.Contains(filename, "io/papermc/") ||
			strings.Contains(filename, "org/spigotmc/") {
			hasMinecraftClasses = true
		}

		// Check for suspicious files
		if v.isSuspiciousFile(filename) {
			suspiciousFiles = append(suspiciousFiles, filename)
		}
	}

	if !hasMinecraftClasses {
		return fmt.Errorf("jar file does not contain Minecraft-specific classes")
	}

	if len(suspiciousFiles) > 0 {
		return fmt.Errorf("jar file contains suspicious files: %v", suspiciousFiles)
	}

	return nil
}

// findBedrockServerDir finds the bedrock server directory
func (v *MinecraftServerValidator) findBedrockServerDir(serverDir string) (string, error) {
	files, err := os.ReadDir(serverDir)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), "bedrock-server") && file.IsDir() {
			return filepath.Join(serverDir, file.Name()), nil
		}
	}

	return "", fmt.Errorf("bedrock server directory not found")
}

// validateBedrockStructure validates the structure of a bedrock server
func (v *MinecraftServerValidator) validateBedrockStructure(bedrockDir string) error {
	// Check for required bedrock server files
	requiredFiles := []string{
		"bedrock_server",     // Linux binary
		"bedrock_server.exe", // Windows binary
		"server.properties",
		"permissions.json",
		"allowlist.json",
	}

	hasExecutable := false
	for _, file := range requiredFiles {
		filePath := filepath.Join(bedrockDir, file)
		if _, err := os.Stat(filePath); err == nil {
			if strings.Contains(file, "bedrock_server") {
				hasExecutable = true
			}
		}
	}

	if !hasExecutable {
		return fmt.Errorf("bedrock server executable not found")
	}

	// Check for required directories
	requiredDirs := []string{
		"behavior_packs",
		"resource_packs",
		"worlds",
	}

	for _, dir := range requiredDirs {
		dirPath := filepath.Join(bedrockDir, dir)
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			return fmt.Errorf("required directory not found: %s", dir)
		}
	}

	return nil
}

// checkForSuspiciousFiles checks for files that shouldn't be in a Minecraft server
func (v *MinecraftServerValidator) checkForSuspiciousFiles(serverDir string) error {
	suspiciousPatterns := []string{
		"index.html",
		"index.php",
		"index.jsp",
		".htaccess",
		"web.config",
		"package.json",
		"Dockerfile",
		"docker-compose",
		".git",
		"node_modules",
		"vendor",
		"www",
		"public_html",
		"webroot",
	}

	return filepath.Walk(serverDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath := strings.TrimPrefix(path, serverDir)
		for _, pattern := range suspiciousPatterns {
			if strings.Contains(strings.ToLower(relPath), pattern) {
				return fmt.Errorf("suspicious file/directory found: %s", relPath)
			}
		}

		return nil
	})
}

// isSuspiciousFile checks if a file inside a jar is suspicious
func (v *MinecraftServerValidator) isSuspiciousFile(filename string) bool {
	lowerName := strings.ToLower(filename)
	
	// Check for web-related file extensions
	webExtensions := []string{
		".jsp",
		".php", 
		".html",
		".htm",
		".css",
	}
	
	for _, ext := range webExtensions {
		if strings.HasSuffix(lowerName, ext) {
			return true
		}
	}
	
	// Check for JavaScript files (but not JSON)
	if strings.HasSuffix(lowerName, ".js") && !strings.HasSuffix(lowerName, ".json") {
		return true
	}
	
	// Check for web-related directories and files
	webPatterns := []string{
		"web.xml",
		"servlet",
		"/webapp/",
		"/web/",
		"/www/",
	}
	
	for _, pattern := range webPatterns {
		if strings.Contains(lowerName, pattern) {
			return true
		}
	}
	
	return false
}

// validateServerProperties validates server.properties file
func (v *MinecraftServerValidator) validateServerProperties(propsPath string) error {
	content, err := os.ReadFile(propsPath)
	if err != nil {
		return err
	}

	props := string(content)

	// Extract server port
	portRegex := regexp.MustCompile(`server-port=(\d+)`)
	matches := portRegex.FindStringSubmatch(props)
	if len(matches) > 1 {
		// Note: We would parse the port and validate it's in allowed range
		// For now, just check it's present
	}

	// Check for suspicious configurations
	if strings.Contains(props, "enable-rcon=true") {
		// RCON can be legitimate but should be noted
	}

	// Check for web-related configurations (suspicious)
	suspiciousConfigs := []string{
		"web-server",
		"http-port",
		"https-port",
		"web-root",
		"document-root",
	}

	for _, config := range suspiciousConfigs {
		if strings.Contains(strings.ToLower(props), config) {
			return fmt.Errorf("suspicious configuration found: %s", config)
		}
	}

	return nil
}

// ValidateServerURL validates that a server URL points to legitimate Minecraft server software
func (v *MinecraftServerValidator) ValidateServerURL(url string) error {
	// List of known legitimate Minecraft server download sources
	legitimateSources := []string{
		"api.papermc.io",
		"piston-data.mojang.com",
		"api.purpurmc.org",
		"ci.md-5.net", // Spigot
		"maven.minecraftforge.net",
		"maven.fabricmc.net",
		"minecraft.net",
		"mojang.com",
	}

	lowerURL := strings.ToLower(url)
	
	for _, source := range legitimateSources {
		if strings.Contains(lowerURL, source) {
			return nil
		}
	}

	return fmt.Errorf("server URL is not from a known legitimate source: %s", url)
}

// CalculateFileHash calculates SHA256 hash of a file
func (v *MinecraftServerValidator) CalculateFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}