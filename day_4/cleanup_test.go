package day_4_test

import (
	a "github.com/akolybelnikov/advent-of-code"
	c "github.com/akolybelnikov/advent-of-code/day_4"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestFindContainedPairs(t *testing.T) {
	t.Run("should handle the default input", func(t *testing.T) {
		data, _ := a.ReadDataBytes("testdata/input.txt")
		total, err := c.FindContainedPairs(data)

		if err != nil {
			t.Log(err)
		}

		assert.Equal(t, err, nil)
		assert.Equal(t, total, 584)
	})
}

func TestFindOverlappingPairs(t *testing.T) {
	t.Run("should handle the default input", func(t *testing.T) {
		data, _ := a.ReadDataBytes("testdata/input.txt")
		total, err := c.FindOverlappingPairs(data)

		if err != nil {
			t.Log(err)
		}

		assert.Equal(t, err, nil)
		assert.Equal(t, total, 933)
	})
}
