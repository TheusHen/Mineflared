
# 🤝 Contributing to Mineflared

Thank you for your interest in contributing to the **Mineflared Project**! Whether you're here to fix a bug, add a feature, improve documentation, or give feedback — you're more than welcome.

This guide applies to all repositories under the [`Mineflared`](https://github.com/TheusHen) organization, including:

- 🧠 [Mineflared](https://github.com/TheusHen/Mineflared) (Core CLI)
- 🌐 [Mineflared-Web](https://github.com/TheusHen/Mineflared-Web) (Frontend)
- 🛰️ [Mineserver](https://github.com/TheusHen/Mineserver) (Server/Cloudflare integration)
- 📦 Any other related services/modules

---

## 📌 Before You Start

- **Check existing issues**: Avoid duplicates.
- **Use clear and respectful language**.
- Make sure your changes **follow the purpose of the project**.
- For major features, open an **issue first** to discuss ideas before implementing.

---

## 🛠️ Local Setup

Each repository has its own setup instructions in the `README.md`. Make sure you:

- Clone the repo
- Install dependencies
- Run locally and test your changes

Example (for Node projects):

```bash
git clone https://github.com/TheusHen/Mineflared-Web.git
cd Mineflared-Website
npm install
npm run dev
````

---

## 🌿 Branch Strategy

* Base branch: `main`
* Feature branches: `feature/<your-feature>`
* Fix branches: `fix/<bug-name>`
* Don't push directly to `main`.

---

## ✅ Pull Request Checklist

Before submitting a PR, make sure:

* [ ] Your code works locally and builds without errors
* [ ] You ran `lint` and/or `format` if available
* [ ] You described your changes clearly in the PR
* [ ] You linked related issue(s) if applicable
* [ ] You only touched related files (no unrelated changes)

> **Note:** Small PRs are easier to review and merge.

---

## 📂 Folder Structure Conventions

If you're contributing to CLI or backend projects:

* Follow clean code principles
* Avoid global scope pollution
* Modularize by responsibility (e.g., `/lib`, `/core`, `/routes`)

For frontend contributions:

* Use meaningful components
* Keep UI and logic separated
* Prefer composition over duplication

---

## 🧪 Testing

If your changes include logic (not just UI), please:

* Add or update tests when possible
* Describe manual testing steps in the PR

---

## 📣 Feature Requests & Ideas

Open an [Issue](https://github.com/TheusHen/Mineflared/issues) with:

* What problem you’re solving
* How your feature addresses it
* Optional: Sketches, mockups, or code examples

---

## 💬 Communication

Join the discussions on the project's GitHub Discussions tab (if enabled), or open a general issue with the `question` label.

---

## 🧾 License

By contributing, you agree that your code will be licensed under the same license as the repository.

---

Made with ☕ by TheusHen.