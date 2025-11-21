package main

import (
	"github.com/akolybelnikov/aoc-2015/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assertions := assert.New(t)
	input := `
123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`

	t.Run("run circuit", func(t *testing.T) {
		expected := map[string]uint16{
			"d": 72,
			"e": 507,
			"f": 492,
			"g": 114,
			"h": 65412,
			"i": 65079,
			"x": 123,
			"y": 456,
		}
		lines, _ := utils.ParseLines(input)
		ws := runCircuit(lines, 0)

		assertions.Equal(ws.circuit, expected)
	})

	t.Run("part 1", func(t *testing.T) {
		expected := uint16(0)
		actual := part1(input)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := uint16(0)
		actual := part2(input)

		assertions.Equal(expected, actual)
	})
}
