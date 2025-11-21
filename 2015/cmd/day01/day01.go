package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("inputs/day01.txt")
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
	var up, down int
	for _, c := range input {
		if c == '(' {
			up++
		} else {
			down++
		}
	}

	return up - down
}

// part two
func part2(input string) int {
	var up, down, pos int
	for i, c := range input {
		if c == '(' {
			up++
		} else {
			down++
		}
		if up-down == -1 {
			return i + 1
		}
	}

	return pos
}
