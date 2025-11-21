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
		{"part 1 input 1", part1, "ugknbfddgicrmopn", 1},
		{"part 1 input 2", part1, "aaa", 1},
		{"part 1 input 3", part1, "jchzalrnumimnmhp", 0},
		{"part 1 input 4", part1, "haegwjzuvuyypxyu", 0},
		{"part 1 input 5", part1, "dvszwmarrgswjxmb", 0},

		// Test cases for part2
		{"part 2 input 4", part2, "qjhvhtzxzqqjkmpb", 1},
		{"part 2 input 2", part2, "xxyxx", 1},
		{"part 2 input 3", part2, "uurcxstgmygtbstg", 0},
		{"part 2 input 4", part2, "ieodomkazucvgmuy", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.function(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
