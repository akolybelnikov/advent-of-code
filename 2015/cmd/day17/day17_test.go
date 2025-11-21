package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assertions := assert.New(t)
	input := `
20
15
10
5
5
`

	t.Run("part 1", func(t *testing.T) {
		expected := 4
		actual := part1(input, 25)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 3
		actual := part2(input, 25)

		assertions.Equal(expected, actual)
	})
}
