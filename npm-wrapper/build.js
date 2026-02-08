#!/usr/bin/env node

const { spawnSync } = require('child_process');
const fs = require('fs');
const path = require('path');
const os = require('os');

const BIN_DIR = path.join(__dirname, 'bin');
fs.mkdirSync(BIN_DIR, { recursive: true });

console.log('Building Go binaries...');

const builds = [
    { goos: 'linux', goarch: 'amd64', output: 'mineflared-linux' },
    { goos: 'linux', goarch: 'arm64', output: 'mineflared-linux-arm64' },
    { goos: 'linux', goarch: 'arm', output: 'mineflared-linux-arm' },
    { goos: 'android', goarch: 'arm64', output: 'mineflared-android-arm64', cgo: '0' },
    { goos: 'darwin', goarch: 'amd64', output: 'mineflared-darwin' },
    { goos: 'darwin', goarch: 'arm64', output: 'mineflared-darwin-arm64' },
    { goos: 'windows', goarch: 'amd64', output: 'mineflared-windows.exe' },
];

function build(goos, goarch, output, cgo) {
    const env = {
        ...process.env,
        GOOS: goos,
        GOARCH: goarch,
    };

    // Set CGO_ENABLED if specified
    if (cgo !== undefined) {
        env.CGO_ENABLED = cgo;
    }

    // Build from the parent directory (where main.go is), using explicit path
    const rootDir = path.join(__dirname, '..');
    const result = spawnSync('go', ['build', '-o', path.join(BIN_DIR, output), rootDir], {
        env,
        stdio: 'inherit'
    });

    if (result.status !== 0) {
        console.error(`❌ Failed to build for ${goos}/${goarch}`);
        process.exit(1);
    }
}

for (const { goos, goarch, output, cgo } of builds) {
    build(goos, goarch, output, cgo);
}

console.log('✅ All binaries built successfully.');
