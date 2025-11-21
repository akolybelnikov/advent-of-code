package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplace(t *testing.T) {
	tests := []struct {
		name         string
		replacements map[element][][]rune
		molecule     []rune
		expected     map[string]int
	}{
		{
			name: "single replacement",
			replacements: map[element][][]rune{
				{'H', 0}: {{'H', 'O'}, {'O', 'H'}},
			},
			molecule: []rune("HOH"),
			expected: map[string]int{
				"HOOH": 2,
				"OHOH": 1,
				"HOHO": 1,
			},
		},
		{
			name: "multiple one rune replacements",
			replacements: map[element][][]rune{
				{'H', 0}: {{'H', 'O'}},
				{'O', 0}: {{'O', 'H'}},
			},
			molecule: []rune("HOH"),
			expected: map[string]int{
				"HOOH": 1,
				"HOHO": 1,
				"HOHH": 1,
			},
		},
		{
			name: "multiple one rune replacements example",
			replacements: map[element][][]rune{
				{'H', 0}: {{'H', 'O'}, {'O', 'H'}},
				{'O', 0}: {{'H', 'H'}},
			},
			molecule: []rune("HOH"),
			expected: map[string]int{
				"HOOH": 2,
				"HOHO": 1,
				"OHOH": 1,
				"HHHH": 1,
			},
		},
		{
			name: "no replacements",
			replacements: map[element][][]rune{
				{'X', 0}: {{'Y'}},
			},
			molecule: []rune("HOH"),
			expected: map[string]int{},
		},
		{
			name:         "empty molecule",
			replacements: map[element][][]rune{{'H', 0}: {{'O'}}},
			molecule:     []rune(""),
			expected:     map[string]int{},
		},
		{
			name:         "empty replacements",
			replacements: map[element][][]rune{},
			molecule:     []rune("HOH"),
			expected:     map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &plant{
				replacements: tt.replacements,
				molecules:    make(map[string]int),
			}

			p.replace(tt.molecule)

			assert.Equal(t, tt.expected, p.molecules)
		})
	}
}

func TestSolution(t *testing.T) {
	assertions := assert.New(t)
	input := `
H => HO
H => OH
O => HH

HOH`

	input2 := `
H => HO
H => OH
O => HH

HOHOHO`

	input3 := `
e => H
e => O
H => HO
H => OH
O => HH

HOH`

	t.Run("part 1", func(t *testing.T) {
		expected := 4
		actual := part1(input)

		assertions.Equal(expected, actual)
	})

	t.Run("part 1 input 2", func(t *testing.T) {
		expected := 7
		actual := part1(input2)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 3
		actual := part2(input3)

		assertions.Equal(expected, actual)
	})
}
