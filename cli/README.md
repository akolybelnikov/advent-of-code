# Advent of Code CLI

A small command-line tool to bootstrap Advent of Code solutions and download puzzle inputs.

This README describes how to build, install, and use the CLI (including new flags to control template source and skipping downloads), and how to troubleshoot the most common issues (templates not found on Windows, authentication/400 errors).

## Requirements

- Go 1.19+ to build from source

## Install

Install the released binary (if available) or build from source.

Build from source (recommended if you want the latest fixes):

```bash
cd /c/Users/akoly/GolandProjects/advent-of-code/cli
go build -o ../bin/aoc-cli.exe main.go
```

After building, the executable will be at `../bin/aoc-cli.exe` relative to the `cli` folder.

You can also install via `go install`:

```bash
go install github.com/akolybelnikov/aoc-cli@latest
```

## Quick usage

Generate a day's solution scaffold and download the puzzle input (default behavior):

```bash
aoc-cli bootstrap --day 1 --path /path/to/project --lang elixir
```

If you only want to create the template files and skip downloading the input (handy when offline or when you don't have a valid session):

```bash
aoc-cli bootstrap --day 1 --path /path/to/project --lang elixir --no-download
```

You can override the template source and point the CLI at a local `templates` folder (useful when debugging or developing templates):

```bash
aoc-cli bootstrap --day 1 --path /path/to/project --lang elixir --templates-path /path/to/cli/internal/templates --no-download
```

Supported languages:
- `go` (default) — creates `cmd/dayXX/`
- `elixir` — creates `lib/` and `test/`
- `rust` — creates `src/bin/dayXX/`

## Flags summary

- `--day, -d` — Day number (1–25)
- `--path, -p` — Project root where files will be created
- `--lang, -l` — Language (`go`, `elixir`, `rust`)
- `--templates-path` — Optional override to load templates from a local folder instead of embedded assets
- `--no-download` — Create templates but skip downloading the puzzle input
- `--year, -y` — Advent of Code year (defaults to the current year)

## Auth / Session management

The CLI requires your Advent of Code session cookie to download inputs. The session token is stored in a file named `.aoc-session` in your home directory (for example, `C:\Users\<you>\.aoc-session` on Windows).

How to obtain the session token:
1. Log into https://adventofcode.com
2. Open Developer Tools -> Application -> Cookies
3. Copy the value of the `session` cookie
4. Run `aoc-cli auth` (if implemented) or manually create the file `~/.aoc-session` and paste the token on a single line

Notes:
- The CLI trims whitespace when reading the session token. If you paste with a newline, that will be handled.
- If session validation fails, the CLI will print a helpful message and stop; re-run `aoc-cli auth` or update `~/.aoc-session` with a fresh token.

## Troubleshooting

Templates not found (Windows / embed-related issue)
- Symptom: `Failed to copy template: open templates\elixir: file does not exist` or similar.
- Cause: Go's `embed.FS` uses forward slashes (`/`) for paths even on Windows. If code constructs the template path using `filepath.Join`, Windows backslashes may be passed to the embedded FS, which prevents files from being found.
- Fixes:
  1. Use the built-in embedded templates (the CLI does this by default). If you see the error, rebuild the CLI from the repository's `cli` directory (so embedded assets are included) and try again.
  2. Use `--templates-path` to point to a local `templates` directory (for example, `cli/internal/templates`).
  3. As a quick workaround, run `aoc-cli` from the repository root where a `templates/` folder exists or copy the `templates` folder next to the binary.

Download returns 400 (Bad Request) or authentication errors
- Symptom: `Failed to download https://adventofcode.com/2025/day/1/input: 400 Bad Request` or `401 Unauthorized`.
- Causes & fixes:
  - No session token present or the token is empty. Ensure `~/.aoc-session` exists and contains the `session` token on a single line.
  - Session token expired or invalid. Re-obtain the `session` cookie and re-run `aoc-cli auth` (or update the file) to refresh it.
  - AoC sometimes rejects requests without a standard `User-Agent`. The CLI now sets a User-Agent header but if you still see errors, try the `--no-download` flag to just scaffold templates and debug auth separately.

To debug authentication:
- Check the `.aoc-session` file exists and is non-empty.

PowerShell example (check file size):

```powershell
# Check session file
Get-Item $env:USERPROFILE\.aoc-session | Select-Object Name, Length
```

If the file is missing or empty, create it and paste your session value.

## Development notes / Contributing

- Templates are embedded into the CLI with Go `//go:embed` and are located under `cli/internal/templates` in the repository.
- When updating templates, rebuild the CLI so embedded assets are included.
- Tests and PRs are welcome.

## Example: Full flow (build, bootstrap, download)

```bash
cd /c/Users/akoly/GolandProjects/advent-of-code/cli
go build -o ../bin/aoc-cli.exe main.go
../bin/aoc-cli.exe bootstrap --day 1 --path /c/Users/akoly/GolandProjects/advent-of-code/2025/elixir --lang elixir
```

If you want to avoid network calls while iterating on templates:

```bash
../bin/aoc-cli.exe bootstrap --day 1 --path /c/Users/akoly/GolandProjects/advent-of-code/2025/elixir --lang elixir --no-download
```

## License

This project follows the repository license (see top-level `LICENSE`).

---

If you want, I can also:
- Add a short example showing the contents of generated Elixir/Go/Rust templates, or
- Add a one-line `--verbose` log that prints which template source the CLI chose at runtime.

Pick one and I'll add it to the README and (optionally) the CLI.
