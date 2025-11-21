package main

import (
	"fmt"
	"github.com/akolybelnikov/aoc-2015/internal/utils"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("inputs/day05.txt")
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
	count := 0
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)
	for _, line := range lines {
		if isInvalid(line) {
			continue
		}
		vowels := "aeiou"
		vCount := 0
		hasDouble := false
		split := strings.Split(line, "")
		for i := 0; i < len(split); i++ {
			if strings.Contains(vowels, split[i]) {
				vCount++
			}
			if i < len(split)-1 && split[i] == split[i+1] {
				hasDouble = true
			}
		}
		if vCount >= 3 && hasDouble {
			count++
		}
	}

	return count
}

func isInvalid(input string) bool {
	forbidden := []string{"ab", "cd", "pq", "xy"}
	for _, f := range forbidden {
		if strings.Contains(input, f) {
			return true
		}
	}
	return false
}

// part two
func part2(input string) int {
	count := 0
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)
	for _, line := range lines {
		mapped := mapIndices(line)
		repeated := false
		for _, indices := range mapped {
			if isRepeated(indices) {
				repeated = true
				break
			}
		}
		repeatedCombo := false
		mappedCombos := mapTwoLetters(line)
		for _, occ := range mappedCombos {
			if len(occ) >= 2 {
				repeatedCombo = true
			}
		}
		if repeated && repeatedCombo {
			count++
		}
	}

	return count
}

func mapIndices(input string) map[string][]int {
	res := make(map[string][]int)
	split := strings.Split(input, "")
	for i, s := range split {
		indices, ok := res[s]
		if !ok {
			indices = make([]int, 0)
		}
		indices = append(indices, i)
		res[s] = indices
	}

	return res
}

func isRepeated(indices []int) bool {
	if len(indices) < 2 {
		return false
	}
	for i := 0; i < len(indices)-1; i++ {
		for j := i + 1; j < len(indices); j++ {
			if indices[j]-indices[i] == 2 {
				return true
			}
		}
	}
	return false
}

func mapTwoLetters(input string) map[string][]string {
	res := make(map[string][]string)
	split := strings.Split(input, "")
	for i := 0; i < len(split)-1; i++ {
		combo := fmt.Sprintf("%s%s", split[i], split[i+1])
		valid := true
		if indices, ok := res[combo]; ok {
			for _, s := range indices {
				if strings.Contains(s, fmt.Sprintf("%d", i)) {
					valid = false
				}
			}
		} else {
			res[combo] = make([]string, 0)
		}
		if valid {
			res[combo] = append(res[combo], fmt.Sprintf("%d%d", i, i+1))
		}
	}
	return res
}
