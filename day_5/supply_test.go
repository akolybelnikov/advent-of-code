package day_5_test

import (
	utils "github.com/akolybelnikov/advent-of-code"
	a "github.com/akolybelnikov/advent-of-code/day_5"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestFindCrates9000(t *testing.T) {
	t.Run("should find top crates in the last state", func(t *testing.T) {
		data, err := utils.ReadDataBytes("testdata/input.txt")
		if err != nil {
			t.Errorf("error reading data bytes: %v", err)
		}

		idx, lines := a.ParseInput(&data)
		assert.Equal(t, idx, 325)

		stacks := a.CreateStacks(lines)
		assert.Equal(t, len(lines[0]), len(*stacks))

		topCrates, err := a.FindTopCrates9000(&data)
		if err != nil {
			t.Errorf("error finding top crates: %v", err)
		}
		assert.Equal(t, len(topCrates), len(*stacks))
		assert.Equal(t, topCrates, "CNSZFDVLJ")
	})
}

func TestFindCrates9001(t *testing.T) {
	t.Run("should find top crates in the last state", func(t *testing.T) {
		data, err := utils.ReadDataBytes("testdata/input.txt")
		if err != nil {
			t.Errorf("error reading data bytes: %v", err)
		}

		idx, lines := a.ParseInput(&data)
		assert.Equal(t, idx, 325)

		stacks := a.CreateStacks(lines)
		assert.Equal(t, len(lines[0]), len(*stacks))

		topCrates, err := a.FindTopCrates9001(&data)
		if err != nil {
			t.Errorf("error finding top crates: %v", err)
		}
		assert.Equal(t, len(topCrates), len(*stacks))
		assert.Equal(t, topCrates, "QNDWLMGNS")
	})
}
