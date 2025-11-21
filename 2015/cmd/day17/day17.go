package main

import (
	"fmt"
	"github.com/akolybelnikov/aoc-2015/internal/utils"
	"math"
	"os"
	"strconv"
)

func main() {
	data, err := os.ReadFile("inputs/day17.txt")
	if err != nil {
		fmt.Println("Failed to read input file.", err)
		os.Exit(0)
	}
	input := string(data)

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input, 150))

	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input, 150))

	os.Exit(0)
}

// part one
func part1(input string, vol int) int {
	ints, err := makeInts(input)
	utils.HandleErr(err)
	var res [][]int
	findValidCombinations(ints, []int{}, vol, 0, &res)
	return len(res)
}

// part two
func part2(input string, vol int) int {
	ints, err := makeInts(input)
	utils.HandleErr(err)
	var res [][]int
	findValidCombinations(ints, []int{}, vol, 0, &res)
	minCombinations := findMinCombinations(&res)
	return len(minCombinations)
}

func makeInts(input string) ([]int, error) {
	var res []int
	lines, err := utils.ParseLines(input)
	if err != nil {
		return nil, err
	}
	for _, line := range lines {
		i, iErr := strconv.Atoi(line)
		if iErr != nil {
			return nil, iErr
		}
		res = append(res, i)
	}
	return res, nil
}

func findValidCombinations(nums, cur []int, target, index int, res *[][]int) {
	if target == 0 {
		combination := make([]int, len(cur))
		copy(combination, cur)
		*res = append(*res, combination)
		return
	}

	if target < 0 || index >= len(nums) {
		return
	}

	findValidCombinations(nums, append(cur, nums[index]), target-nums[index], index+1, res)
	findValidCombinations(nums, cur, target, index+1, res)
}

func findMinCombinations(combinations *[][]int) []int {
	var minCombinations []int
	minContainers := math.MaxInt
	for i, combination := range *combinations {
		if len(combination) < minContainers {
			minContainers = len(combination)
			minCombinations = []int{i}
		} else if len(combination) == minContainers {
			minCombinations = append(minCombinations, i)
		}
	}
	return minCombinations
}
