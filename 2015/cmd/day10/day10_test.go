package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		part     func(string) string
	}{
		{
			name:     "part 1 - example 1",
			input:    "1",
			expected: "11",
			part:     lookAndSay,
		},
		{
			name:     "part 1 - example 2",
			input:    "11",
			expected: "21",
			part:     lookAndSay,
		},
		{
			name:     "part 1 - example 3",
			input:    "21",
			expected: "1211",
			part:     lookAndSay,
		},
		{
			name:     "part 1 - example 4",
			input:    "1211",
			expected: "111221",
			part:     lookAndSay,
		},
		{
			name:     "part 1 - example 5",
			input:    "111221",
			expected: "312211",
			part:     lookAndSay,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.part(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
