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
    { goos: 'darwin', goarch: 'amd64', output: 'mineflared-darwin' },
    { goos: 'windows', goarch: 'amd64', output: 'mineflared-windows.exe' },
];

function build(goos, goarch, output) {
    const env = {
        ...process.env,
        GOOS: goos,
        GOARCH: goarch,
    };

    const result = spawnSync('go', ['build', '-o', path.join(BIN_DIR, output)], {
        env,
        stdio: 'inherit',
        shell: true
    });

    if (result.status !== 0) {
        console.error(`❌ Failed to build for ${goos}/${goarch}`);
        process.exit(1);
    }
}

for (const { goos, goarch, output } of builds) {
    build(goos, goarch, output);
}

console.log('✅ All binaries built successfully.');
