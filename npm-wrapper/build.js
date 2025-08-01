// Do not run this file in Windows because it uses `execSync` which is not compatible with Windows for this purpose.
// If you want to run in Windows implement `cross-env`.

const { execSync } = require("child_process");
const fs = require("fs");
const path = require("path");

const BIN_DIR = path.join(__dirname, "bin");
fs.mkdirSync(BIN_DIR, { recursive: true });

console.log("Building Go binaries...");

try {
    execSync("GOOS=linux GOARCH=amd64 go build -o bin/mineflared-linux", { stdio: "inherit" });
    execSync("GOOS=darwin GOARCH=amd64 go build -o bin/mineflared-darwin", { stdio: "inherit" });
    execSync("GOOS=windows GOARCH=amd64 go build -o bin/mineflared-windows.exe", { stdio: "inherit" });
    console.log("✅ All binaries built successfully.");
} catch (e) {
    console.error("❌ Failed to build binaries:", e.message);
    process.exit(1);
}
