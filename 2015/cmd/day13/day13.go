package main

import (
	"fmt"
	"github.com/akolybelnikov/aoc-2015/internal/utils"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.ReadFile("inputs/day13.txt")
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
	happiness := parseInput(lines)
	return maxHappiness(happiness)
}

// part two
func part2(input string) int {
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)
	happiness := parseInput(lines)
	// Second puzzle: add "Me" with happiness of 0 sitting next to everyone
	peopleMap := make(map[string]bool)
	for _, happy := range happiness {
		peopleMap[happy.name] = true
	}
	// Add "Me" to happiness relationships
	for person := range peopleMap {
		happiness = append(happiness, Happiness{"Me", 0, person}) // "Me" → other
		happiness = append(happiness, Happiness{person, 0, "Me"}) // other → "Me"
	}

	return maxHappiness(happiness)
}

// Happiness represents a structure for happiness values between two people.
type Happiness struct {
	name  string
	gain  int
	other string
}

// Parse input into a slice of Happiness structures.
func parseInput(lines []string) []Happiness {
	pattern := regexp.MustCompile(`([A-Z][a-z]+) would ([a-z]+) ([0-9]+) happiness units by sitting next to ([A-Z][a-z]+)\.`)
	var happiness []Happiness

	for _, line := range lines {
		match := pattern.FindStringSubmatch(line)
		if len(match) > 0 {
			name := match[1]
			gain, _ := strconv.Atoi(match[3])
			other := match[4]
			if match[2] == "lose" {
				gain = -gain
			}
			happiness = append(happiness, Happiness{name, gain, other})
		}
	}
	return happiness
}

// Calculate total happiness for a given seating arrangement.
func calculate(table []string, h []Happiness) int {
	first := table[0]
	last := table[len(table)-1]
	fullTable := append([]string{last}, append(table, first)...)

	totalHappiness := 0
	for i := 1; i < len(fullTable)-1; i++ {
		totalHappiness += happinessBetween(fullTable[i], fullTable[i-1], h) +
			happinessBetween(fullTable[i], fullTable[i+1], h)
	}
	return totalHappiness
}

// Get happiness value between two people.
func happinessBetween(name, other string, h []Happiness) int {
	for _, happy := range h {
		if happy.name == name && happy.other == other {
			return happy.gain
		}
	}
	return math.MinInt // Default to a very small value, though this shouldn't occur.
}

// Generate all permutations of people.
func permutations(elements []string) [][]string {
	if len(elements) <= 1 {
		return [][]string{elements}
	}

	var result [][]string
	for i, elem := range elements {
		rest := append([]string{}, elements[:i]...)
		rest = append(rest, elements[i+1:]...)
		for _, perm := range permutations(rest) {
			result = append(result, append([]string{elem}, perm...))
		}
	}
	return result
}

// Find the maximum happiness achievable.
func maxHappiness(h []Happiness) int {
	peopleMap := make(map[string]bool)
	for _, happy := range h {
		peopleMap[happy.name] = true
	}

	// Extract unique people
	var people []string
	for person := range peopleMap {
		people = append(people, person)
	}

	// Generate all permutations and calculate happiness
	allPermutations := permutations(people)
	mh := math.MinInt
	for _, perm := range allPermutations {
		happiness := calculate(perm, h)
		if happiness > mh {
			mh = happiness
		}
	}
	return mh
}
