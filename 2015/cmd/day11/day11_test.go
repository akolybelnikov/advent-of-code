package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{name: "valid string with sequential letters and 2 pairs", input: "aabbcde", expected: true},
		{name: "invalid string with disallowed characters", input: "abcdeiol", expected: false},
		{name: "valid string with sequential letters and no 'iol'", input: "abcdffaa", expected: true},
		{name: "invalid string missing sequential letters", input: "aabbccdd", expected: false},
		{name: "valid string at the boundary", input: "xxxyzz", expected: true},
		{name: "invalid empty string", input: "", expected: false},
		{name: "invalid string with a single pair", input: "aabcde", expected: false},
		{name: "valid string with sequential letters and exactly 2 pairs", input: "abcddeff", expected: true},
		{name: "invalid string with i and l included", input: "hijklmmn", expected: false},
		{name: "invalid string without a straight", input: "abbceffg", expected: false},
		{name: "invalid string with only one double", input: "abbcegjk", expected: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := isValid(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestSolution(t *testing.T) {
	assertions := assert.New(t)
	input := "abcdefgh"
	input2 := "ghijklmn"

	t.Run("part 1", func(t *testing.T) {
		expected := "abcdffaa"
		actual := part1(input)

		assertions.Equal(expected, actual)
	})

	t.Run("part 1 input 2", func(t *testing.T) {
		expected := "ghjaabcc"
		actual := part1(input2)

		assertions.Equal(expected, actual)
	})
}

func TestIncrement(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "simple increment", input: "abc", expected: "abd"},
		{name: "single carry", input: "azz", expected: "baa"},
		{name: "multiple carry", input: "zz", expected: "aa"},
		{name: "long carry chain", input: "aazz", expected: "abaa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := increment(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
