#!/usr/bin/env node
const path = require('path');
const os = require('os');
const { spawn } = require('child_process');
const fs = require('fs');

const platform = os.platform();
const arch = os.arch();
let binary = '';

// Detect if running on Android/Termux
function isAndroid() {
    // Check for Android-specific environment variables
    if (process.env.ANDROID_ROOT || process.env.ANDROID_DATA) {
        return true;
    }
    
    // Check for Termux-specific environment
    if (process.env.TERMUX_VERSION || (process.env.PREFIX && process.env.PREFIX.includes('/com.termux/'))) {
        return true;
    }
    
    // Check if /system/bin/sh exists (Android-specific path)
    try {
        if (fs.existsSync('/system/bin/sh') && !fs.existsSync('/usr/bin/sh')) {
            return true;
        }
    } catch (e) {
        // Ignore errors
    }
    
    return false;
}

if (platform === 'linux') {
    // Check if running on Android/Termux
    if (isAndroid()) {
        if (arch === 'arm64' || arch === 'aarch64') {
            // Try Android-specific binary first
            const androidBinary = 'mineflared-android-arm64';
            const androidPath = path.join(__dirname, 'bin', androidBinary);
            if (fs.existsSync(androidPath)) {
                binary = androidBinary;
            } else {
                // Fall back to Linux ARM64 binary (Termux provides Linux compatibility)
                binary = 'mineflared-linux-arm64';
            }
        } else if (arch === 'arm') {
            binary = 'mineflared-linux-arm';
        } else {
            // For x64/amd64 on Android, use Linux binary (Termux compatibility)
            binary = 'mineflared-linux';
        }
    } else {
        // Regular Linux
        if (arch === 'arm64') {
            binary = 'mineflared-linux-arm64';
        } else if (arch === 'arm') {
            binary = 'mineflared-linux-arm';
        } else {
            binary = 'mineflared-linux';
        }
    }
} else if (platform === 'darwin') {
    if (arch === 'arm64') {
        binary = 'mineflared-darwin-arm64';
    } else {
        binary = 'mineflared-darwin';
    }
} else if (platform === 'win32' || platform === 'win64') {
    binary = 'mineflared-windows.exe';
} else if (platform === 'android') {
    // Handle when os.platform() directly reports 'android'
    if (arch === 'arm64' || arch === 'aarch64') {
        // Try Android-specific binary first
        const androidBinary = 'mineflared-android-arm64';
        const androidPath = path.join(__dirname, 'bin', androidBinary);
        if (fs.existsSync(androidPath)) {
            binary = androidBinary;
        } else {
            // Fall back to Linux ARM64 binary (Termux provides Linux compatibility)
            binary = 'mineflared-linux-arm64';
        }
    } else if (arch === 'arm') {
        binary = 'mineflared-linux-arm';
    } else {
        // For x64/amd64 on Android, use Linux binary (Termux compatibility)
        binary = 'mineflared-linux';
    }
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