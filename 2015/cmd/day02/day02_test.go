package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assertions := assert.New(t)
	input := `
2x3x4
1x1x10`

	t.Run("part 1", func(t *testing.T) {
		expected := 101
		actual := part1(input)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 48
		actual := part2(input)

		assertions.Equal(expected, actual)
	})
}

func TestBoxTotalSqFeet(t *testing.T) {
	tests := []struct {
		name     string
		box      box
		expected int
	}{
		{
			name:     "all positive dimensions",
			box:      box{l: 2, w: 3, h: 4, s: 6},
			expected: 58, // 2*6 + 2*12 + 2*8 + 6
		},
		{
			name:     "zero dimensions",
			box:      box{l: 0, w: 0, h: 0, s: 0},
			expected: 0,
		},
		{
			name:     "one side zero",
			box:      box{l: 0, w: 2, h: 3, s: 0},
			expected: 12, // 2*0 + 2*6 + 2*0 + 0
		},
		{
			name:     "large dimensions",
			box:      box{l: 10, w: 20, h: 30, s: 200},
			expected: 2400, // 2*200 + 2*600 + 2*300 + 200
		},
		{
			name:     "small dimensions",
			box:      box{l: 1, w: 1, h: 10, s: 1},
			expected: 43, // 2*1 + 2*10 + 2*10 + 1
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.box.totalSqFeet()
			assert.Equal(t, tt.expected, actual)
		})
	}
}
