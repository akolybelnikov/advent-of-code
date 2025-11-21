# Advent of Code

Solutions for the [Advent calendar of small programming puzzles](https://adventofcode.com/).

This mono-repository contains solutions for multiple years in different languages.

## Structure

- `2015/` - Go solutions (see [README](2015/README.md))
- `2022/` - Go solutions (see [README](2022/README.md))
- `2023/` - Rust solutions (see [README](2023/README.md))
- `2024/` - Go solutions (see [README](2024/README.md))
- `2025/` - Elixir solutions (see [README](2025/README.md))
- `cli/` - AOC CLI tool for automating daily challenges

## AOC CLI Tool

This repository includes a custom CLI tool written in Go that automates downloading puzzle inputs and bootstrapping daily solutions for Advent of Code.

### Installation

You can install the CLI locally to run from your terminal:

```bash
cd cli
go install
```

This will install the `aoc-cli` command to your `GOPATH/bin` directory.

### First Time Setup

**Important:** Before using the bootstrap command, authenticate with your Advent of Code session token:

```bash
aoc-cli auth
```

The CLI will prompt you for your session token and store it securely in `~/.aoc-session`.

### Supported Languages

The CLI supports bootstrapping solutions in multiple languages:
- **Go** (default) - Creates files in `cmd/dayXX/`
- **Elixir** - Creates files in `lib/` and `test/`
- **Rust** - Creates files in `src/bin/dayXX/`

### Usage Examples

Bootstrap a Go solution (default):
```bash
aoc-cli bootstrap --day 5 --path /path/to/your/project
```

Bootstrap an Elixir solution:
```bash
aoc-cli bootstrap --day 5 --path /path/to/your/project --lang elixir
```

Bootstrap a Rust solution:
```bash
aoc-cli bootstrap --day 5 --path /path/to/your/project --lang rust
```

This command will:
- Create a solution directory for the specified day
- Generate boilerplate code for the solution
- Download your puzzle input automatically

### Session Token Setup

To get your Advent of Code session token:

1. Log in to [Advent of Code](https://adventofcode.com)
2. Open Developer Tools (F12 or Ctrl+Shift+I)
3. Navigate to **Application** → **Cookies**
4. Copy the value of the `session` cookie
5. Run `aoc-cli auth` and paste the token when prompted

For more details on the CLI tool, see the [CLI README](cli/README.md).

## License

[MIT License](LICENSE)