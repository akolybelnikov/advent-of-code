package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("inputs/day11.txt")
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
func part1(input string) string {
	input = strings.TrimSpace(input)
	s := increment(input)
	for !isValid(s) {
		s = increment(s)
	}

	return s
}

// part two
func part2(input string) string {
	input = part1(input)
	s := increment(input)
	for !isValid(s) {
		s = increment(s)
	}

	return s
}

func increment(s string) string {
	bs := []byte(s)
	idx := len(bs) - 1
	carry := true
	for carry {
		bs[idx]++
		if bs[idx] > 'z' {
			bs[idx] = 'a'
			idx--
			if idx < 0 {
				break
			}
		} else {
			carry = false
		}
	}

	return string(bs)
}

func isValid(s string) bool {
	if !strings.ContainsAny(s, "iol") && hasTwoDistinctPairs(s) {
		for i := 0; i <= len(s)-3; i++ {
			substring := s[i : i+3] // Extract 3 consecutive letters
			if isSequential(substring) {
				return true
			}
		}
	}

	return false
}

func isSequential(s string) bool {
	return len(s) == 3 &&
		s[1] == s[0]+1 &&
		s[2] == s[1]+1
}

func hasTwoDistinctPairs(s string) bool {
	pairs := make(map[string]bool) // To store distinct pairs

	// Loop through the string to find pairs
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] { // Check if two consecutive characters are the same
			pair := string(s[i]) + string(s[i+1])
			pairs[pair] = true
			i++ // Skip the next character to ensure pairs are non-overlapping
		}
	}

	// Check if there are at least two distinct pairs
	return len(pairs) >= 2
}
