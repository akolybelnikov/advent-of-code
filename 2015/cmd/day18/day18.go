package main

import (
	"fmt"
	"github.com/akolybelnikov/aoc-2015/internal/utils"
	"image"
	"os"
)

type grid map[image.Point]int32

func (g *grid) nextState() *grid {
	next := make(grid)
	for k, v := range *g {
		on := 0
		for _, n := range [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
			p := image.Pt(k.X+n[0], k.Y+n[1])
			if c, ok := (*g)[p]; ok && c == 35 {
				on++
			}
		}
		if v == 35 {
			if on == 2 || on == 3 {
				next[k] = 35
			} else {
				next[k] = 46
			}
		} else {
			if on == 3 {
				next[k] = 35
			} else {
				next[k] = 46
			}
		}
	}
	return &next
}

func (g *grid) count() int {
	var res int
	for _, v := range *g {
		if v == 35 {
			res++
		}
	}
	return res
}

func (g *grid) bounds() (int, int, int, int) {
	// Determine grid bounds
	minX, minY, maxX, maxY := 0, 0, 0, 0
	for pt := range *g {
		if pt.X < minX {
			minX = pt.X
		}
		if pt.Y < minY {
			minY = pt.Y
		}
		if pt.X > maxX {
			maxX = pt.X
		}
		if pt.Y > maxY {
			maxY = pt.Y
		}
	}
	return minX, minY, maxX, maxY
}

func (g *grid) setCorners() {
	minX, minY, maxX, maxY := g.bounds()
	corners := [4]image.Point{
		image.Pt(minX, minY),
		image.Pt(minX, maxY),
		image.Pt(maxX, minY),
		image.Pt(maxX, maxY),
	}
	for _, c := range corners {
		(*g)[c] = 35
	}
}

func (g *grid) print() {
	minX, minY, maxX, maxY := g.bounds()
	fmt.Println()
	// Iterate through coordinates and print grid
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			p := image.Pt(x, y)
			if (*g)[p] == 35 { // ASCII for `#`
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	data, err := os.ReadFile("inputs/day18.txt")
	if err != nil {
		fmt.Println("Failed to read input file.", err)
		os.Exit(0)
	}
	input := string(data)

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input, 100))

	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input, 100))

	os.Exit(0)
}

// part one
func part1(input string, transitions int) int {
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)
	g := createGrid(lines)
	for range transitions {
		ng := g.nextState()
		g = *ng
	}
	return g.count()
}

// part two
func part2(input string, transitions int) int {
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)
	g := createGrid(lines)
	g.setCorners()
	for range transitions {
		ng := g.nextState()
		g = *ng
		g.setCorners()
	}
	return g.count()
}

func createGrid(lines []string) grid {
	g := make(grid)
	for i, line := range lines {
		for j, c := range line {
			p := image.Pt(j, i)
			g[p] = c
		}
	}
	return g
}
