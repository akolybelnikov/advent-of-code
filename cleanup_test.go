package advent_test

import (
	"github.com/go-playground/assert/v2"
	a "scripts/advent"
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
