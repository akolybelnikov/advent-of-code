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
		{"part 1 input 1", part1, "turn on 0,0 through 999,999", 1000000},
		{"part 1 input 2", part1, "toggle 0,0 through 999,0", 1000},
		{"part 1 input 3", part1, "turn off 499,499 through 500,500", 0},

		// Test cases for part2
		{"part 2 input 4", part2, "turn on 0,0 through 0,0", 1},
		{"part 2 input 2", part2, "toggle 0,0 through 999,999", 2000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.function(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
