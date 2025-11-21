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
		{"part 1 input 1", part1, "abcdef", 609043},
		{"part 1 input 2", part1, "pqrstuv", 1048970},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.function(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
