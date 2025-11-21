package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assertions := assert.New(t)
	input := "()())"

	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "test 1",
			input:    "(())",
			expected: 0,
		},
		{
			name:     "test 2",
			input:    "()()",
			expected: 0,
		},
		{
			name:     "test 3",
			input:    "(((",
			expected: 3,
		},
		{
			name:     "test 4",
			input:    "(()(()(",
			expected: 3,
		},
		{
			name:     "test 5",
			input:    "))(((((",
			expected: 3,
		},
		{
			name:     "test 6",
			input:    "())",
			expected: -1,
		},
		{
			name:     "test 7",
			input:    "))(",
			expected: -1,
		},
		{
			name:     "test 8",
			input:    ")))",
			expected: -3,
		},
		{
			name:     "test 9",
			input:    ")())())",
			expected: -3,
		},
	}

	for _, tt := range tests {
		t.Run("part 1", func(t *testing.T) {
			t.Logf(
				"Running: %s",
				tt.name)
			expected := tt.expected
			actual := part1(tt.input)

			assertions.Equal(expected, actual)
		})
	}

	t.Run("part 2", func(t *testing.T) {
		expected := 5
		actual := part2(input)

		assertions.Equal(expected, actual)
	})
}
