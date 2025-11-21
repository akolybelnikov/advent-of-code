package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		function func(string) int
		input    string
		expected int
	}{
		// Test cases for part1
		{"part 1 input 1", part1, `""`, 2},
		{"part 1 input 2", part1, `"abc"`, 2},
		{"part 1 input 3", part1, `"aaa\"aaa"`, 3},
		{"part 1 input 4", part1, `"\x27"`, 5},
		{"part 1 input 5", part1, `"xs"`, 2},

		// Test cases for part2
		{"part 2 input 1", part2, `""`, 4},
		{"part 2 input 2", part2, `"abc"`, 4},
		{"part 2 input 3", part2, `"aaa\"aaa"`, 6},
		{"part 2 input 2", part2, `"\x27"`, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.function(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
