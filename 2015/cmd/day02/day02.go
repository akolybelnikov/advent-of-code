package main

import (
	"fmt"
	"github.com/akolybelnikov/aoc-2015/internal/utils"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("inputs/day02.txt")
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
	res := 0
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)

	for _, line := range lines {
		b := newBox(line)
		res += b.totalSqFeet()
	}

	return res
}

// part two
func part2(input string) int {
	res := 0
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)

	for _, line := range lines {
		b := newBox(line)
		res += b.totalRibbon()
	}

	return res
}

type box struct {
	l, w, h, s int
}

func (b *box) totalSqFeet() int {
	return 2*b.l*b.w + 2*b.w*b.h + 2*b.h*b.l + b.s
}

func (b *box) totalRibbon() int {
	return b.l*b.w*b.h + b.smallestPerimeter()*2
}

func (b *box) smallestPerimeter() int {
	s := b.l + b.w
	if b.w+b.h < s {
		s = b.w + b.h
	}
	if b.h+b.l < s {
		s = b.h + b.l
	}

	return s
}

func newBox(line string) *box {
	split := strings.Split(line, "x")
	l, _ := strconv.Atoi(split[0])
	w, _ := strconv.Atoi(split[1])
	h, _ := strconv.Atoi(split[2])

	return &box{l, w, h, smallestArea(l*w, w*h, h*l)}
}

func smallestArea(l, w, h int) int {
	m := l
	if w < m {
		m = w
	}
	if h < m {
		m = h
	}
	return m
}
