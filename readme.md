# Taskr

**Taskr** is a lightweight command-line tool written in Go for quickly storing, listing, and removing timestamped notes. It uses a simple `.csv` file as the backend and a TOML-based config file to control where notes are stored.

---

## 🧠 Features

- Add timestamped notes with a single command
- List all existing notes with index
- Remove notes by index
- Count the number of notes
- Configurable file location via TOML
- Human-readable CSV backend
- Clean shell integration with dynamic prompt

---

## 🛠 Installation

### 1. Build from source

```bash
git clone https://github.com/patpragman/taskr.git
cd taskr
go build -o taskr
```

### 2. Set up your config file

Create a `config.toml` file at:

```bash
~/.local/etc/taskr/config.toml
```

Example:

```toml
Version = "1.0.0"
LocalStorage = true
StorageAddress = "/home/yourusername/.local/share/taskr/notes.csv"
```

Create the storage directory if it doesn’t exist:

```bash
mkdir -p ~/.local/share/taskr
touch ~/.local/share/taskr/notes.csv
```

⚠️ **Note:** Taskr requires a valid config file at `~/.local/etc/taskr/config.toml`. It will not function without it.

---

## 🚀 Usage

```bash
taskr <command> [arguments...]
```

### Commands

| Command           | Description                                  |
|-------------------|----------------------------------------------|
| `about`           | Display version and usage instructions       |
| `add <text>`      | Add a new note                               |
| `list`            | List all notes with index and timestamp      |
| `remove <index>`  | Remove note at index                         |
| `rm <index>`      | Alias for `remove`                           |
| `n`               | Show total number of notes                   |

### Examples

```bash
taskr add Buy chicken feed
taskr list
taskr remove 3
taskr n
```

---

## 🖥️ Shell Prompt Integration

Add the following to your `.bashrc` to display the current number of notes directly in your terminal prompt:

```bash
function taskr_prompt() {
    local count
    count=$(taskr n 2>/dev/null)

    # Bright ANSI colors with bold
    RED="\[\e[1;91m\]"
    GREEN="\[\e[1;92m\]"
    BLUE="\[\e[1;34m\]"
    RESET="\[\e[0m\]"

    PS1="${RED}(${count})${RESET} ${GREEN}\u@\h${RESET}:${BLUE}\w${RESET}\$ "
}

PROMPT_COMMAND=taskr_prompt
```

---

## 📁 File Format

Your notes are stored in a CSV with the following headers:

```csv
Date,Note
2025-05-26T10:43:21Z,"Call mom"
```

---

## 📦 Future Goals

- Store notes in a private `github` repository for sync and backup
- Develop a simple mobile app to view and add notes
- Create additional tools for automation and filtering

---

## 👤 Author

Written by **Pat Pragman**  
[www.pragman.io](https://www.pragman.io)  
[github.com/patpragman](https://github.com/patpragman)

---

## 📝 License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.


---

## ✨ Contributions

Pull requests and issues are welcome. Please keep it lean and hackable.
