package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assertions := assert.New(t)
	input := `
London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141`

	t.Run("part 1", func(t *testing.T) {
		expected := 605
		actual := part1(input)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 982
		actual := part2(input)

		assertions.Equal(expected, actual)
	})
}
