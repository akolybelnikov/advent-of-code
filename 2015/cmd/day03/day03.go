package main

import (
	"fmt"
	"image"
	"os"
)

func main() {
	data, err := os.ReadFile("inputs/day03.txt")
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
	cur := image.Pt(0, 0)
	visited := map[image.Point]int{cur: 1}
	for _, c := range input {
		cur = next(cur, c)
		visited[cur]++
	}

	return len(visited)
}

// part two
func part2(input string) int {
	cur, curRobot := image.Pt(0, 0), image.Pt(0, 0)
	visited := map[image.Point]int{cur: 2}
	for i, c := range input {
		turn := i % 2
		if turn == 0 {
			cur = next(cur, c)
			visited[cur]++
		} else {
			curRobot = next(curRobot, c)
			visited[curRobot]++
		}

	}
	return len(visited)
}

func next(cur image.Point, dir rune) image.Point {
	switch dir {
	case '^':
		return image.Pt(cur.X, cur.Y-1)
	case '>':
		return image.Pt(cur.X+1, cur.Y)
	case 'v':
		return image.Pt(cur.X, cur.Y+1)
	case '<':
		return image.Pt(cur.X-1, cur.Y)
	default:
		return cur
	}
}
