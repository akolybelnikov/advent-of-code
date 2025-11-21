# Advent of Code

Solutions for the [Advent calendar of small programming puzzles](https://adventofcode.com/).

This mono-repository contains solutions for multiple years in different languages.

## Structure

- `2015/` - Go solutions (see [README](2015/README.md))
- `2022/` - Go solutions (see [README](2022/README.md))
- `2023/` - Rust solutions (see [README](2023/README.md))
- `2024/` - Go solutions (see [README](2024/README.md))
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

### Usage

The CLI tool helps you quickly set up solutions for each day and automatically downloads your puzzle inputs.

#### Bootstrap a day's solution

```bash
aoc-cli bootstrap --day 5 --path /path/to/your/project
```

This command will:
- Create a solution directory for the specified day
- Generate boilerplate code for the solution
- Download your puzzle input (if session token is configured)

#### Session Token Setup

To download puzzle inputs, you need to provide your Advent of Code session token:

1. Log in to [Advent of Code](https://adventofcode.com)
2. Open Developer Tools (F12 or Ctrl+Shift+I)
3. Navigate to **Application** → **Cookies**
4. Copy the value of the `session` cookie
5. Paste it into the CLI when prompted (stored securely in `~/.aoc-session`)

For more details on the CLI tool, see the [CLI README](cli/README.md).

## License

[MIT License](LICENSE)