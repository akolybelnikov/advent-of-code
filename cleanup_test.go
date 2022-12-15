package advent_test

import (
	a "github.com/akolybelnikov/advent-of-code"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestFindContainedPairs(t *testing.T) {
	t.Run("should handle the default input", func(t *testing.T) {
		data, _ := a.ReadDataBytes("testdata/cleanup/input.txt")
		total, err := a.FindContainedPairs(data)

		if err != nil {
			t.Log(err)
		}

		assert.Equal(t, err, nil)
		assert.Equal(t, total, 584)
	})
}

func TestFindOverlappingPairs(t *testing.T) {
	t.Run("should handle the default input", func(t *testing.T) {
		data, _ := a.ReadDataBytes("testdata/cleanup/input.txt")
		total, err := a.FindOverlappingPairs(data)

		if err != nil {
			t.Log(err)
		}

		assert.Equal(t, err, nil)
		assert.Equal(t, total, 933)
	})
}
