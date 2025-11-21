package main

import (
	"fmt"
	"github.com/akolybelnikov/aoc-2015/internal/utils"
	"os"
	"strconv"
	"strings"
)

type reindeer struct {
	name      string
	speed     int
	flyTime   int
	restTime  int
	distance  int
	points    int
	cycleTime int
}

type reindeerList []*reindeer

func main() {
	data, err := os.ReadFile("inputs/day14.txt")
	if err != nil {
		fmt.Println("Failed to read input file.", err)
		os.Exit(0)
	}
	input := string(data)

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input, 2503))

	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input, 2503))

	os.Exit(0)
}

// part one
func part1(input string, sec int) int {
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)

	reindeers := make([]*reindeer, len(lines))
	for i, line := range lines {
		reindeers[i] = newReindeer(line)
	}

	for _, r := range reindeers {
		r.calculateDistance(sec)
	}

	maxDist := 0
	for _, r := range reindeers {
		if r.distance > maxDist {
			maxDist = r.distance
		}
	}

	return maxDist
}

// part two
func part2(input string, sec int) int {
	maxPoints := 0
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)

	reindeers := make(reindeerList, len(lines))
	for i, line := range lines {
		reindeers[i] = newReindeer(line)
	}

	for i := 1; i < sec; i++ {
		reindeers.race(i)
		reindeers.reward()
	}

	for _, r := range reindeers {
		if r.points > maxPoints {
			maxPoints = r.points
		}
	}

	return maxPoints
}

// newReindeer creates and returns a pointer to a reindeer struct by parsing the provided input line.
// The input line is expected to be a string containing the reindeer's name, speed, flight time, and rest time.
// It uses helper functions to parse and handle numerical values from the input string.
func newReindeer(line string) *reindeer {
	const (
		nameIndex     = 0
		speedIndex    = 3
		flyTimeIndex  = 6
		restTimeIndex = 13
	)

	fields := strings.Fields(line)

	speed := parseAndHandle(fields[speedIndex])
	flyTime := parseAndHandle(fields[flyTimeIndex])
	restTime := parseAndHandle(fields[restTimeIndex])

	return &reindeer{
		name:      fields[nameIndex],
		speed:     speed,
		flyTime:   flyTime,
		restTime:  restTime,
		cycleTime: flyTime + restTime,
	}
}

// Helper function to parse an integer and handle errors.
func parseAndHandle(value string) int {
	parsedValue, err := strconv.Atoi(value)
	utils.HandleErr(err)
	return parsedValue
}

// calculateDistance computes the total distance a reindeer travels in the given number of seconds.
func (r *reindeer) calculateDistance(seconds int) {
	completedCycles := seconds / r.cycleTime
	remainingTime := seconds % r.cycleTime

	// Calculate distance for completed cycles
	r.distance += r.calculateBaseDistance(completedCycles)

	// Calculate distance for remaining time
	if remainingTime > r.flyTime {
		r.distance += r.speed * r.flyTime
	} else {
		r.distance += r.speed * remainingTime
	}
}

// Extracted helper function for calculating base distance during full cycles
func (r *reindeer) calculateBaseDistance(cycles int) int {
	return cycles * r.speed * r.flyTime
}

func (l *reindeerList) race(sec int) {
	for _, r := range *l {
		mtime := sec % r.cycleTime
		if mtime > 0 && mtime <= r.flyTime {
			r.distance += r.speed
		}
	}
}

func (l *reindeerList) reward() {
	lead := 0
	for _, r := range *l {
		if r.distance > lead {
			lead = r.distance
		}
	}
	for _, r := range *l {
		if r.distance == lead {
			r.points++
		}
	}
}
