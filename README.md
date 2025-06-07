# Discord Bot in Go 🧠🤖

A modular discord bot built in Go with Docker support. This bot supports both traditional command prefixes (like `!ping`). It also integrates with an external API to deliver jokes, and includes a secure message deletion function for server owners.

---

## ✨ Features

- 🔁 Message-based commands (e.g. `!hello`, `!bye`, etc.)
- ⚡ Slash command support (`/ping`, `/joke`, `/delete`)
- 📡 External API integration (`https://official-joke-api.appspot.com`)
- 🔒 Owner-only bulk message deletion
- 🐳 Dockerized for local and cloud deployment
- ⚙️ Environment-based configuration (`.env` file for secrets)

---

## 💻 Commands

| Command     | Type         | Description                          |
|-------------|--------------|--------------------------------------|
| `!ping`| Text | Replies with "Pong!"                   |
| `!hello` | Text | Greets the user                        |
| `!bye` | Text | Says goodbye                           |
| `!help` | Text | Lists available commands               |
| `!whoami` | Text | Identifies the bot                     |
| `!google` | Text | Sends a link to Google                 |
| `!joke`  | Text | Fetches and tells a random joke        |
| `!delete` | Text | Deletes recent messages (owner only)   |

---

## 📦 Tech Stack

- [Go](https://golang.org/)
- [discordgo](https://github.com/bwmarrin/discordgo)
- [Docker](https://www.docker.com/)
- [Official Joke API](https://official-joke-api.appspot.com)

---

## 🚀 Getting Started

### Prerequisites
- Docker installed
- Discord bot token
- A `.env` file containing:
  ```env
  TOKEN=your_discord_bot_token_here
