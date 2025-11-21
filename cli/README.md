# Advent of Code CLI

A CLI tool to automate downloading inputs for Advent of Code, written in Go.

## Installation

You can install the CLI directly using Go without cloning the repository.

### Requirements
- Go 1.19+ installed

### Install via Go
```bash
go install github.com/akolybelnikov/aoc-cli@latest
```

## Bootstrap a day's solution and download your puzzle inputs

The `bootstrap` command helps you quickly set up your solution for a specific day and download the corresponding puzzle 
inputs. It creates a new solution directory for the given day and populates it with the necessary boilerplate code.

### Supported Languages

The CLI supports multiple languages for bootstrapping solutions:
- **Go** (default) - Creates files in `cmd/dayXX/`
- **Elixir** - Creates files in `lib/` and `test/`

### Examples

Bootstrap a Go solution (default):
```shell
aoc-cli bootstrap --day 5 --path /projects/my-project
```

Bootstrap an Elixir solution:
```shell
aoc-cli bootstrap --day 5 --path /projects/my-project --lang elixir
```

### What Gets Created

**For Go:**
- `cmd/day05/day05.go` - Solution file with `part1()` and `part2()` functions
- `cmd/day05/day05_test.go` - Test file
- `inputs/day05.txt` - Input file (if session is configured)

**For Elixir:**
- `lib/day05.ex` - Solution module with `part1()`, `part2()`, and `solve()` functions
- `test/day05_test.exs` - ExUnit test file
- `inputs/day05.txt` - Input file (if session is configured)

If valid session credentials are already configured, the input file will be downloaded automatically.

## Auth - Session Management

### Why Session Management?
Advent of Code (AoC) puzzles and inputs require authentication. AoC uses a session cookie (`session` token) to verify users. 
This CLI stores the session locally, allowing automated puzzle downloads without repeated logins.

### How It Works:
1. **First Run (No Session Found):**
    - CLI checks for `.aoc-session` in the user's home directory.
    - If not found, the CLI prompts for the session token.
    - The token is saved securely in `.aoc-session` (with `0600` permissions).

2. **Subsequent Runs (Session Exists):**
    - CLI reads the token from `.aoc-session`.
    - The token is used to authenticate and fetch puzzles/inputs.

3. **Security:**
    - Token stored with restricted permissions (user-only read/write).

---

### How to Get Your Session Token:
1. Log in to [Advent of Code](https://adventofcode.com).
2. Open Developer Tools (`F12` or `Ctrl+Shift+I`).
3. Navigate to **Application** -> **Cookies**.
4. Copy the value of the `session` cookie.
5. Paste it into the CLI when prompted.

---

### Notes:
- **Session Expiry:** If the session expires, the CLI will prompt for a new token.
- **No Auto-Login:** The CLI cannot extract the session directly from your browser for security reasons.
- **No OAuth:** AoC does not offer an API, so session cookies are the only way to authenticate.

---

This method keeps authentication simple and mirrors the manual download process while enabling automation.

