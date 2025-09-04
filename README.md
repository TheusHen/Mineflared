# 🟩 Mineflared – Minecraft P2P Hosting CLI

[![License: MIT](https://img.shields.io/github/license/TheusHen/Mineflared?style=flat-square)](LICENSE)
[![npm version](https://img.shields.io/npm/v/mineflared?style=flat-square)](https://www.npmjs.com/package/mineflared)
[![Build](https://github.com/TheusHen/Mineflared/actions/workflows/publish.yml/badge.svg)](https://github.com/TheusHen/Mineflared/actions)

> [!CAUTION]
> Version V3 has been permanently discontinued and is awaiting V4, which will fix **critical bugs**.

Understand these bugs:
- ✅ **FIXED**: Security for creating Minecraft-only servers and prevention of scam sites implemented.

- ✅ **FIXED**: The DNS configuration for Cloudflare is incorrect and needs to be fixed.

Notes:
Proxy protection for Bedrock servers will not be possible due to UDP/IP, and Java needs to be updated to support connecting to the default Minecraft port.

> [!IMPORTANT]  
> API temporarily disabled, will be enabled again in V4. **DO NOT attempt to use Mineflared, errors are expected without the API**
---

**Host Minecraft servers with a few commands. No VPS, no static IP – Cloudflare protected.**

---

## 🚀 Features

- ⚡ Instant setup in seconds
- 🛡️ DDoS protection via Cloudflare  
- 🔒 **Security validation** to prevent non-Minecraft content
- 🌐 Dynamic DNS with custom subdomain  
- 📈 Real-time monitoring  
- 🔄 Auto-restart on crash  

---

## 📦 Installation

```bash
npm install -g mineflared
````

Or via PowerShell:

```powershell
iex (iwr -Uri "https://mineflared.theushen.me/install.ps1").Content
```

---

## 🔐 Authentication

Log in with GitHub:

```bash
mineflared login
```

---

## 🛠️ Usage

```bash
mineflared create           # Create new Java/Bedrock server
mineflared start <name>     # Start an existing server
mineflared status           # Show current server status
mineflared list             # List all your servers
mineflared config <name>    # Open server.properties editor in browser
mineflared backup <name>    # Create .zip or .rar backup
mineflared restore          # Restore from backup
mineflared language         # Change CLI language
```

---

## 🔒 Security

Mineflared includes comprehensive security validation to ensure only legitimate Minecraft servers are hosted:

### 🛡️ **Server Validation**
- **URL Verification**: Only allows downloads from official Minecraft sources (PaperMC, Mojang, Purpur, etc.)
- **Content Analysis**: Validates JAR files contain Minecraft-specific classes and packages
- **File Scanning**: Detects and blocks web content (HTML, PHP, JSP, CSS, JS files)
- **Directory Structure**: Ensures proper Minecraft server structure for both Java and Bedrock

### 🚫 **Blocked Content**
- Web applications and websites
- Non-Minecraft server software  
- Mixed Java/Bedrock configurations
- Suspicious file uploads

### ✅ **Legitimate Sources**
- `api.papermc.io` (PaperMC)
- `piston-data.mojang.com` (Vanilla Minecraft)
- `api.purpurmc.org` (Purpur)
- `maven.minecraftforge.net` (Forge)
- `maven.fabricmc.net` (Fabric)
- `minecraft.net` (Bedrock)

---

## 🔒 Privacy & Data

**Local (CLI)**:
Stored at `$HOME/.config/minecli/config.json`:

* GitHub username
* Auth token (JWT)
* Public IP
* Language

**Backend (API)**:
Stored in MongoDB:

* GitHub username & ID
* JWT token
* Last known IP
* GitHub access token

**Cloudflare**:
Used to update subdomain `<username>.mineflared.theushen.me` with IP.

✅ No passwords, emails, server files or telemetry are collected.
❌ No third-party data sharing (only GitHub & Cloudflare integrations).

---

## 🗑️ Data Deletion

Run `mineflared delete` to:

* Remove your data from the database
* Delete your Cloudflare subdomain

This action is irreversible.

---

## 📄 License

[MIT License](LICENSE)

---

## 📚 Links

* 🌍 [Website](https://mineflared.theushen.me)
* 📦 [NPM](https://www.npmjs.com/package/mineflared)
* 💻 [CLI GitHub](https://github.com/TheusHen/Mineflared)
* 🧩 [Mineserver GitHub](https://github.com/TheusHen/mineserver)
* 🖥️ [Mineflared Web GitHub](https://github.com/TheusHen/mineflared-web)
