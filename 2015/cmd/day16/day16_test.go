package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		element  string
		expected bool
	}{
		{"element in slice", []string{"a", "b", "c"}, "b", true},
		{"element not in slice", []string{"a", "b", "c"}, "d", false},
		{"empty slice", []string{}, "a", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, contains(tt.slice, tt.element))
		})
	}
}

func TestParseThings(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected1 string
		expected2 int
		expectErr bool
	}{
		{"valid input", "children: 3", "children", 3, false},
		{"invalid number", "cats: invalid", "", 0, true},
		{"missing colon", "trees 3", "", 0, true},
		{"empty input", "", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual1, actual2, err := parseThings(tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected1, actual1)
				assert.Equal(t, tt.expected2, actual2)
			}
		})
	}
}

func TestMakeAunts(t *testing.T) {
	tests := []struct {
		name        string
		lines       []string
		expectedLen int
		expectErr   bool
	}{
		{
			"valid input",
			[]string{
				"Sue 1: children: 1, cats: 2, samoyeds: 1",
				"Sue 2: pomeranians: 3, akitas: 4, vizslas: 0",
			},
			2,
			false,
		},
		{
			"empty input",
			[]string{},
			0,
			false,
		},
		{
			"invalid input",
			[]string{"invalid line"},
			0,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !tt.expectErr {
					t.Errorf("TestMakeAunts panicked: %v", r)
				}
			}()
			aunts := makeAunts(tt.lines)
			assert.Equal(t, tt.expectedLen, len(aunts))
		})
	}
}
