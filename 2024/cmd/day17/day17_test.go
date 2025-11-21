package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay17(t *testing.T) {
	assertions := assert.New(t)
	input := `
Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

	t.Run("part 1", func(t *testing.T) {
		expected := "4,6,3,5,6,3,5,2,1,0"
		actual := part1(input)

		assertions.Equal(expected, actual)
	})

	input2 := `
Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

	t.Run("part 2", func(t *testing.T) {
		expected := 117440
		actual := part2(input2)

		assertions.Equal(expected, actual)
	})
}
