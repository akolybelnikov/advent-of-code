package main

import (
	"fmt"
	"github.com/akolybelnikov/aoc-2015/internal/utils"
	"os"
	"strconv"
	"strings"
)

const message = `
children: 3
cats: 7
samoyeds: 2
pomeranians: 3
akitas: 0
vizslas: 0
goldfish: 5
trees: 3
cars: 2
perfumes: 1
`

type aunt struct {
	num       int
	things    []string
	thingsMap map[string]int
}

func main() {
	data, err := os.ReadFile("inputs/day16.txt")
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
	aunts := makeAunts(lines)
	messageLines, err := utils.ParseLines(message)
	utils.HandleErr(err)
	for _, a := range aunts {
		if contains(messageLines, a.things[0]) &&
			contains(messageLines, a.things[1]) && contains(messageLines, a.things[2]) {
			return a.num
		}
	}

	return 0
}

// part two
func part2(input string) int {
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)
	aunts := makeAunts(lines)
	mm := makeMessageMap()
	for _, a := range aunts {
		oks := 3
		for k, v := range a.thingsMap {
			if k == "cats" || k == "trees" {
				if v <= mm[k] {
					oks--
				}
			} else if k == "pomeranians" || k == "goldfish" {
				if v >= mm[k] {
					oks--
				}
			} else {
				if v != mm[k] {
					oks--
				}
			}
		}
		if oks == 3 {
			return a.num
		}
	}

	return 0
}

func makeAunts(lines []string) []*aunt {
	aunts := make([]*aunt, len(lines))
	for i, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		numStr := strings.Split(parts[0], " ")[1]
		num, err := strconv.Atoi(numStr)
		utils.HandleErr(err)
		things := strings.Split(parts[1], ", ")
		a := aunt{
			num:       num,
			things:    []string{},
			thingsMap: map[string]int{},
		}
		for _, thing := range things {
			a.things = append(a.things, strings.TrimSpace(thing))
			t, n, err := parseThings(thing)
			utils.HandleErr(err)
			a.thingsMap[t] = n
		}
		aunts[i] = &a
	}

	return aunts
}

func parseThings(line string) (string, int, error) {
	parts := strings.SplitN(line, ":", 2)
	if len(parts) != 2 {
		return "", 0, fmt.Errorf("invalid line: %s", line)
	}
	num, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	return strings.TrimSpace(parts[0]), num, err
}

func makeMessageMap() map[string]int {
	return map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
