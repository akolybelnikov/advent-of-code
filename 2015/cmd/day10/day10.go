package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("inputs/day10.txt")
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
	input = strings.TrimSpace(input)
	for i := 0; i < 40; i++ {
		input = lookAndSay(input)
	}

	return len(input)
}

// part two
func part2(input string) int {
	input = strings.TrimSpace(input)
	for i := 0; i < 50; i++ {
		input = lookAndSay(input)
	}

	return len(input)
}

func lookAndSay(input string) string {
	b := strings.Builder{}
	bs := []byte(input)
	idx := 0
	for idx < len(bs) {
		cur := bs[idx]
		count := 0
		for idx < len(bs) && bs[idx] == cur {
			idx++
			count++
		}
		b.WriteString(fmt.Sprintf("%d%c", count, cur))
	}

	return b.String()
}
