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
		{"part 1 input 1", part1, ">", 2},
		{"part 1 input 2", part1, "^>v<", 4},
		{"part 1 input 3", part1, "^v^v^v^v^v", 2},

		// Test cases for part2
		{"part 2 input 4", part2, "^v", 3},
		{"part 2 input 2", part2, "^>v<", 3},
		{"part 2 input 3", part2, "^v^v^v^v^v", 11},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.function(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
