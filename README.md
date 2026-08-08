# 🛠️ Bertu — AI-Powered Developer CLI Tools

Bertu is a CLI toolbox built in Go that uses Google's Gemini AI to automate common development tasks. It compiles to a **single binary** you can drop into any machine and use immediately.

## ✨ Features

| Command | Description |
|---------|-------------|
| `bertu git` | Reads your staged `git diff` and generates a conventional commit message |
| `bertu explain` | Pipe any failing command's output and get an AI-powered explanation + fix |
| `bertu chat` | Ask natural language questions about your codebase (RAG with local vector DB) |

---

## 📦 Installation

### Option 1: Install with `go install` (Recommended)

If you have Go installed (1.21+), run:

```bash
go install github.com/bigisigi50/bertu@latest
```

This downloads, compiles, and places the `bertu` binary in your `$GOPATH/bin` (usually `~/go/bin`). Make sure that directory is in your `PATH`:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Add that line to your `~/.bashrc` or `~/.zshrc` to make it permanent.

### Option 2: Build from Source

```bash
git clone https://github.com/bigisigi50/bertu.git
cd bertu
go build -o bertu .
```

Then move the binary somewhere in your PATH:

```bash
sudo mv bertu /usr/local/bin/
```

---

## ⚙️ Setup

Bertu requires a **Google Gemini API Key** to work. Get one for free at [Google AI Studio](https://aistudio.google.com/apikey).

Set it as an environment variable:

```bash
export GEMINI_API_KEY="your_api_key_here"
```

To make it permanent, add it to your shell config:

```bash
echo 'export GEMINI_API_KEY="your_api_key_here"' >> ~/.bashrc
source ~/.bashrc
```

---

## 🚀 Usage

### Smart Git Assistant

Automatically generate commit messages from your staged changes:

```bash
git add .
bertu git
```

Output:
```
Generating commit message...

Generated commit message:

feat: add user authentication with JWT tokens
```

### Terminal Error Explainer

Pipe any failing command into `bertu explain`:

```bash
npm run build 2>&1 | bertu explain
go build ./... 2>&1 | bertu explain
cargo build 2>&1 | bertu explain
```

Output:
```
Analyzing error...

Explanation & Fix:

The error indicates a missing dependency...
Run 'npm install' to fix this.
```

### Codebase Chatter

First, index your project (run this once or after major code changes):

```bash
bertu chat --init
```

Then ask questions about your code:

```bash
bertu chat "How does the authentication middleware work?"
bertu chat "Where is the database connection configured?"
bertu chat "What does the handleUpload function do?"
```

---

## 🐚 Shell Autocompletion

Enable tab-completion for bertu commands:

```bash
# Bash
echo 'source <(bertu completion bash)' >> ~/.bashrc

# Zsh
echo 'source <(bertu completion zsh)' >> ~/.zshrc

# Fish
bertu completion fish | source
```

---

## 🏗️ Tech Stack

- **Language:** Go
- **CLI Framework:** [Cobra](https://github.com/spf13/cobra)
- **AI Model:** Google Gemini (gemini-2.5-flash)
- **Embeddings:** Gemini Embedding (gemini-embedding-001)
- **Vector Store:** [chromem-go](https://github.com/philippgille/chromem-go) (pure Go, no external DB needed)

---

## 📄 License

MIT
