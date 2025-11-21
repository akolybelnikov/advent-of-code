package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assertions := assert.New(t)
	input := `
.#.#.#
...##.
#....#
..#...
#.#..#
####..
`

	t.Run("part 1", func(t *testing.T) {
		expected := 4
		actual := part1(input, 4)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 17
		actual := part2(input, 5)

		assertions.Equal(expected, actual)
	})
}
