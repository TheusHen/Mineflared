package internal

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMinecraftServerValidator_ValidateServerURL(t *testing.T) {
	validator := NewMinecraftServerValidator()

	// Test legitimate URLs
	legitimateURLs := []string{
		"https://api.papermc.io/v2/projects/paper/versions/1.21.4/builds/231/downloads/paper-1.21.4-231.jar",
		"https://piston-data.mojang.com/v1/objects/6bce4ef400e4efaa63a13d5e6f6b500be969ef81/server.jar",
		"https://api.purpurmc.org/v2/purpur/1.21.8/2478/download",
		"https://minecraft.net/content/dam/minecraft/bedrock-server/bedrock-server-1.21.95.1.zip",
	}

	for _, url := range legitimateURLs {
		err := validator.ValidateServerURL(url)
		if err != nil {
			t.Errorf("Expected legitimate URL to pass validation: %s, got error: %v", url, err)
		}
	}

	// Test suspicious URLs
	suspiciousURLs := []string{
		"https://example.com/malicious-server.jar",
		"https://sketchy-site.net/fake-minecraft.jar",
		"http://unknown-source.org/server.jar",
		"https://github.com/malicious-user/fake-server/releases/download/v1.0/server.jar",
	}

	for _, url := range suspiciousURLs {
		err := validator.ValidateServerURL(url)
		if err == nil {
			t.Errorf("Expected suspicious URL to fail validation: %s", url)
		}
	}
}

func TestMinecraftServerValidator_IsKnownMinecraftJar(t *testing.T) {
	validator := NewMinecraftServerValidator()

	// Test known Minecraft jar patterns
	knownJars := []string{
		"paper-1.21.4-231.jar",
		"vanilla-1.21.1.jar",
		"purpur-1.21.8-2478.jar",
		"server.jar",
		"minecraft_server.1.21.4.jar",
		"spigot-1.21.4.jar",
		"bukkit-1.21.4.jar",
		"forge-1.21.4-52.0.17.jar",
		"fabric-server-mc.1.21.4-loader.0.16.9-launcher.1.0.1.jar",
	}

	for _, jar := range knownJars {
		if !validator.isKnownMinecraftJar(jar) {
			t.Errorf("Expected known Minecraft jar to be recognized: %s", jar)
		}
	}

	// Test unknown/suspicious jar names
	unknownJars := []string{
		"malicious.jar",
		"webapp.jar",
		"exploit.jar",
		"backdoor.jar",
		"web-server.jar",
	}

	for _, jar := range unknownJars {
		if validator.isKnownMinecraftJar(jar) {
			t.Errorf("Expected unknown jar to not be recognized as Minecraft: %s", jar)
		}
	}
}

func TestMinecraftServerValidator_DetectServerType(t *testing.T) {
	validator := NewMinecraftServerValidator()

	// Create temporary directory for testing
	tempDir := t.TempDir()

	// Test Java server detection
	javaDir := filepath.Join(tempDir, "java-server")
	os.MkdirAll(javaDir, 0755)
	os.WriteFile(filepath.Join(javaDir, "paper-1.21.4.jar"), []byte("fake jar content"), 0644)

	serverType, err := validator.detectServerType(javaDir)
	if err != nil {
		t.Fatalf("Unexpected error detecting Java server: %v", err)
	}
	if serverType != "java" {
		t.Errorf("Expected 'java', got '%s'", serverType)
	}

	// Test Bedrock server detection
	bedrockDir := filepath.Join(tempDir, "bedrock-server")
	bedrockServerDir := filepath.Join(bedrockDir, "bedrock-server-1.21.95.1")
	os.MkdirAll(bedrockServerDir, 0755)

	serverType, err = validator.detectServerType(bedrockDir)
	if err != nil {
		t.Fatalf("Unexpected error detecting Bedrock server: %v", err)
	}
	if serverType != "bedrock" {
		t.Errorf("Expected 'bedrock', got '%s'", serverType)
	}

	// Test mixed content (should fail)
	mixedDir := filepath.Join(tempDir, "mixed-server")
	os.MkdirAll(mixedDir, 0755)
	os.WriteFile(filepath.Join(mixedDir, "server.jar"), []byte("fake jar"), 0644)
	os.MkdirAll(filepath.Join(mixedDir, "bedrock-server-1.21"), 0755)

	_, err = validator.detectServerType(mixedDir)
	if err == nil {
		t.Error("Expected error for mixed Java/Bedrock content")
	}

	// Test empty directory (should fail)
	emptyDir := filepath.Join(tempDir, "empty-server")
	os.MkdirAll(emptyDir, 0755)

	_, err = validator.detectServerType(emptyDir)
	if err == nil {
		t.Error("Expected error for empty directory")
	}
}

func TestMinecraftServerValidator_CheckForSuspiciousFiles(t *testing.T) {
	validator := NewMinecraftServerValidator()

	// Create temporary directory for testing
	tempDir := t.TempDir()

	// Test clean server directory
	cleanDir := filepath.Join(tempDir, "clean-server")
	os.MkdirAll(cleanDir, 0755)
	os.WriteFile(filepath.Join(cleanDir, "server.jar"), []byte("minecraft server"), 0644)
	os.WriteFile(filepath.Join(cleanDir, "server.properties"), []byte("server-port=25565"), 0644)

	err := validator.checkForSuspiciousFiles(cleanDir)
	if err != nil {
		t.Errorf("Clean server directory should pass: %v", err)
	}

	// Test directory with suspicious files
	suspiciousDir := filepath.Join(tempDir, "suspicious-server")
	os.MkdirAll(suspiciousDir, 0755)
	os.WriteFile(filepath.Join(suspiciousDir, "server.jar"), []byte("minecraft server"), 0644)
	os.WriteFile(filepath.Join(suspiciousDir, "index.html"), []byte("<html>malicious</html>"), 0644)

	err = validator.checkForSuspiciousFiles(suspiciousDir)
	if err == nil {
		t.Error("Directory with suspicious files should fail validation")
	}

	// Test directory with web-related files
	webDir := filepath.Join(tempDir, "web-server")
	os.MkdirAll(webDir, 0755)
	webSubDir := filepath.Join(webDir, "www")
	os.MkdirAll(webSubDir, 0755)
	os.WriteFile(filepath.Join(webSubDir, "index.php"), []byte("<?php malicious code ?>"), 0644)

	err = validator.checkForSuspiciousFiles(webDir)
	if err == nil {
		t.Error("Directory with web files should fail validation")
	}
}

func TestMinecraftServerValidator_IsSuspiciousFile(t *testing.T) {
	validator := NewMinecraftServerValidator()

	// Test legitimate files
	legitimateFiles := []string{
		"net/minecraft/server/Main.class",
		"com/mojang/authlib/Agent.class",
		"org/bukkit/Bukkit.class",
		"META-INF/MANIFEST.MF",
		"log4j2.xml",
		"version.json",
	}

	for _, file := range legitimateFiles {
		if validator.isSuspiciousFile(file) {
			t.Errorf("Legitimate file should not be flagged as suspicious: %s", file)
		}
	}

	// Test suspicious files
	suspiciousFiles := []string{
		"web/index.html",
		"webapp/login.jsp",
		"malicious.php",
		"exploit.js",
		"style.css",
		"WEB-INF/web.xml",
		"javax/servlet/Servlet.class",
	}

	for _, file := range suspiciousFiles {
		if !validator.isSuspiciousFile(file) {
			t.Errorf("Suspicious file should be flagged: %s", file)
		}
	}
}

func TestMinecraftServerValidator_FindMainJarFile(t *testing.T) {
	validator := NewMinecraftServerValidator()

	// Create temporary directory for testing
	tempDir := t.TempDir()

	// Test directory with single known Minecraft jar
	singleJarDir := filepath.Join(tempDir, "single-jar")
	os.MkdirAll(singleJarDir, 0755)
	os.WriteFile(filepath.Join(singleJarDir, "paper-1.21.4.jar"), []byte("minecraft server"), 0644)

	jarPath, err := validator.findMainJarFile(singleJarDir)
	if err != nil {
		t.Fatalf("Unexpected error finding jar file: %v", err)
	}
	expectedPath := filepath.Join(singleJarDir, "paper-1.21.4.jar")
	if jarPath != expectedPath {
		t.Errorf("Expected %s, got %s", expectedPath, jarPath)
	}

	// Test directory with multiple jars but one is known Minecraft jar
	multiJarDir := filepath.Join(tempDir, "multi-jar")
	os.MkdirAll(multiJarDir, 0755)
	os.WriteFile(filepath.Join(multiJarDir, "random.jar"), []byte("random jar"), 0644)
	os.WriteFile(filepath.Join(multiJarDir, "server.jar"), []byte("minecraft server"), 0644)
	os.WriteFile(filepath.Join(multiJarDir, "library.jar"), []byte("library"), 0644)

	jarPath, err = validator.findMainJarFile(multiJarDir)
	if err != nil {
		t.Fatalf("Unexpected error finding jar file: %v", err)
	}
	expectedPath = filepath.Join(multiJarDir, "server.jar")
	if jarPath != expectedPath {
		t.Errorf("Expected %s, got %s", expectedPath, jarPath)
	}

	// Test directory with no jar files
	noJarDir := filepath.Join(tempDir, "no-jar")
	os.MkdirAll(noJarDir, 0755)
	os.WriteFile(filepath.Join(noJarDir, "readme.txt"), []byte("no jars here"), 0644)

	_, err = validator.findMainJarFile(noJarDir)
	if err == nil {
		t.Error("Expected error for directory with no jar files")
	}

	// Test directory with multiple unknown jars (should fail)
	unknownJarsDir := filepath.Join(tempDir, "unknown-jars")
	os.MkdirAll(unknownJarsDir, 0755)
	os.WriteFile(filepath.Join(unknownJarsDir, "malicious1.jar"), []byte("bad jar"), 0644)
	os.WriteFile(filepath.Join(unknownJarsDir, "malicious2.jar"), []byte("bad jar"), 0644)

	_, err = validator.findMainJarFile(unknownJarsDir)
	if err == nil {
		t.Error("Expected error for directory with multiple unknown jars")
	}
}