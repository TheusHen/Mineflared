#!/usr/bin/env node
const path = require('path');
const os = require('os');
const { spawn } = require('child_process');
const fs = require('fs');

const platform = os.platform();
const arch = os.arch();
let binary = '';

if (platform === 'linux') {
    if (arch === 'arm64' || arch === 'aarch64') {
        binary = 'mineflared-linux-arm64';
    } else if (arch === 'arm') {
        binary = 'mineflared-linux-arm';
    } else {
        binary = 'mineflared-linux';
    }
} else if (platform === 'darwin') {
    if (arch === 'arm64') {
        binary = 'mineflared-darwin-arm64';
    } else {
        binary = 'mineflared-darwin';
    }
} else if (platform === 'win32' || platform === 'win64') {
    binary = 'mineflared-windows.exe';
} else {
    console.error(`Unsupported platform: ${platform}`);
    process.exit(1);
}

const binPath = path.join(__dirname, 'bin', binary);

if (!fs.existsSync(binPath)) {
    console.error(`Binary not found: ${binPath}`);
    process.exit(1);
}

const args = process.argv.slice(2);
spawn(binPath, args, { stdio: 'inherit' });