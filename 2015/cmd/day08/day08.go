package main

import (
	"fmt"
	"github.com/akolybelnikov/aoc-2015/internal/utils"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("inputs/day08.txt")
	if err != nil {
		fmt.Println("Failed to read input file.", err)
		os.Exit(0)
	}
	input := string(data)

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input))

	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input))

	os.Exit(0)
}

// part one
func part1(input string) int {
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)

	sLit, sVal := 0, 0
	for _, line := range lines {
		sLit += len([]byte(line))
		sVal += decode(line)
	}

	return sLit - sVal
}

// part two
func part2(input string) int {
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)

	sLit, sEnc := 0, 0
	for _, line := range lines {
		sLit += len([]byte(line))
		enc := encode(line)
		sEnc += len([]byte(enc))
	}

	return sEnc - sLit
}

func decode(line string) int {
	count := 0
	cur := 1
	for cur < len(line)-1 {
		if line[cur] == '\\' && cur < len(line)-2 {
			if line[cur+1] == 'x' {
				cur += 4
			} else if line[cur+1] == '\\' || line[cur+1] == '"' {
				cur += 2
			}
		} else {
			cur++
		}
		count++
	}

	return count
}

func encode(line string) string {
	var b strings.Builder
	b.WriteByte('"')
	b.WriteByte('\\')
	b.WriteByte('"')

	cur := 1
	for cur < len(line)-1 {
		if line[cur] == '\\' && cur < len(line)-2 {
			if line[cur+1] == 'x' {
				b.WriteByte('\\')
				b.WriteByte('\\')
				b.WriteByte('x')
				b.WriteByte(line[cur+2])
				b.WriteByte(line[cur+3])
				cur += 4
			} else if line[cur+1] == '\\' || line[cur+1] == '"' {
				b.WriteByte('\\')
				b.WriteByte('\\')
				b.WriteByte('\\')
				b.WriteByte('"')
				cur += 2
			}
		} else {
			b.WriteByte(line[cur])
			cur++
		}
	}

	b.WriteByte('\\')
	b.WriteByte('"')
	b.WriteByte('"')

	return b.String()
}
