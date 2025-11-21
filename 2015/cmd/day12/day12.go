package main

import (
	"encoding/json"
	"fmt"
	"github.com/akolybelnikov/aoc-2015/internal/utils"
	"os"
)

func main() {
	data, err := os.ReadFile("inputs/day12.txt")
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
	var data interface{}
	err := json.Unmarshal([]byte(input), &data)
	utils.HandleErr(err)
	res := process(data, false)
	return res
}

// part two
func part2(input string) int {
	var data interface{}
	err := json.Unmarshal([]byte(input), &data)
	utils.HandleErr(err)
	res := process(data, true)
	return res
}

// process prints JSON values recursively
func process(data interface{}, part2 bool) int {
	var sum int
	switch v := data.(type) {
	case map[string]interface{}: // JSON Object
		for _, value := range v {
			if part2 && value == "red" {
				return 0
			}
			sum += process(value, part2)
		}
	case []interface{}: // JSON Array
		for _, value := range v {
			sum += process(value, part2)
		}
	case float64: // JSON Number
		return int(v)
	default:
	}

	return sum
}
