package main

import (
	"fmt"
	"github.com/akolybelnikov/aoc-2015/internal/utils"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("inputs/day06.txt")
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
	grid := makeGrid(1000)
	for _, line := range lines {
		instr := parseInstruction(line)
		for i := instr.rangeX[0]; i <= instr.rangeX[1]; i++ {
			for j := instr.rangeY[0]; j <= instr.rangeY[1]; j++ {
				switch instr.op {
				case "on":
					grid[i][j] = 1
				case "off":
					grid[i][j] = 0
				default:
					if grid[i][j] == 0 {
						grid[i][j] = 1
					} else {
						grid[i][j] = 0
					}
				}
			}
		}
	}
	for _, row := range grid {
		for _, cell := range row {
			if cell == 1 {
				count++
			}
		}
	}

	return count
}

// part two
func part2(input string) int {
	count := 0
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)
	grid := makeGrid(1000)
	for _, line := range lines {
		instr := parseInstruction(line)
		for i := instr.rangeX[0]; i <= instr.rangeX[1]; i++ {
			for j := instr.rangeY[0]; j <= instr.rangeY[1]; j++ {
				switch instr.op {
				case "on":
					grid[i][j]++
				case "off":
					if grid[i][j] > 0 {
						grid[i][j]--
					}
				default:
					grid[i][j] += 2
				}
			}
		}
	}
	for _, row := range grid {
		for _, cell := range row {
			count += cell
		}
	}

	return count
}

func makeGrid(size int) [][]int {
	res := make([][]int, size)
	for i := 0; i < size; i++ {
		res[i] = make([]int, size)
		for j := 0; j < size; j++ {
			res[i][j] = 0
		}
	}
	return res
}

type instruction struct {
	op     string
	rangeX [2]int
	rangeY [2]int
}

func parseInstruction(input string) instruction {
	instr := instruction{}
	var i, j int
	fields := strings.Fields(input)
	if len(fields) != 5 {
		instr.op = fields[0]
		i, j = 1, 3
	} else {
		instr.op = fields[1]
		i, j = 2, 4
	}
	start := parseRange(fields[i])
	end := parseRange(fields[j])
	instr.rangeX = [2]int{start[0], end[0]}
	instr.rangeY = [2]int{start[1], end[1]}

	return instr
}

func parseRange(input string) [2]int {
	r := [2]int{}
	fields := strings.Split(input, ",")
	r[0], _ = strconv.Atoi(fields[0])
	r[1], _ = strconv.Atoi(fields[1])
	return r
}
