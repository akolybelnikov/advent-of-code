package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assertions := assert.New(t)
	input := `
Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3
`

	t.Run("part 1", func(t *testing.T) {
		expected := 62842880
		actual := part1(input)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 57600000
		actual := part2(input)

		assertions.Equal(expected, actual)
	})
}

func TestMakeIngredient(t *testing.T) {
	// Table-driven test cases
	tests := []struct {
		name     string
		input    string
		expected *ingredient
	}{
		{
			name:  "valid input - Butterscotch",
			input: "Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8",
			expected: &ingredient{
				name:       "Butterscotch",
				capacity:   -1,
				durability: -2,
				flavor:     6,
				texture:    3,
				calories:   8,
			},
		},
		{
			name:  "valid input - Cinnamon",
			input: "Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3",
			expected: &ingredient{
				name:       "Cinnamon",
				capacity:   2,
				durability: 3,
				flavor:     -2,
				texture:    -1,
				calories:   3,
			},
		},
		{
			name:  "edge case - negative all values",
			input: "Pepper: capacity -2, durability -2, flavor -2, texture -2, calories -2",
			expected: &ingredient{
				name:       "Pepper",
				capacity:   -2,
				durability: -2,
				flavor:     -2,
				texture:    -2,
				calories:   -2,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := makeIngredient(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
