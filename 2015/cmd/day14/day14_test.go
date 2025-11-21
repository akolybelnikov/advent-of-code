package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assertions := assert.New(t)
	input := `
Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`

	t.Run("part 1", func(t *testing.T) {
		expected := 1120
		actual := part1(input, 1000)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 689
		actual := part2(input, 1000)

		assertions.Equal(expected, actual)
	})
}
