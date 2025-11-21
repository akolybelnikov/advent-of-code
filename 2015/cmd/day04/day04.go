package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("inputs/day04.txt")
	if err != nil {
		fmt.Println("Failed to read input file.", err)
		os.Exit(0)
	}
	input := string(data)
	input = strings.TrimSpace(input)

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input))

	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input))

	os.Exit(0)
}

// part one
func part1(input string) int {
	return findHash(input, "00000")
}

// part two
func part2(input string) int {
	return findHash(input, "000000")
}

func makeHash(input string) [16]byte {
	return md5.Sum([]byte(input))
}

func findHash(input string, prefix string) int {
	i := 0
	for {
		hashInput := fmt.Sprintf("%s%d", input, i)
		hash := makeHash(hashInput)
		strHash := fmt.Sprintf("%x", hash)
		if strings.HasPrefix(strHash, prefix) {
			return i
		}
		i++
	}
}
