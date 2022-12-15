package advent_test

import (
	"github.com/go-playground/assert/v2"
	a "scripts/advent"
	"testing"
)

func TestFindTotalPriorities(t *testing.T) {
	t.Run("should handle default input", func(t *testing.T) {
		data, _ := a.ReadDataBytes("testdata/rucksack/input.txt")
		total, err := a.FindTotalPriorities(data)

		if err != nil {
			t.Errorf("encountered error %v", err)
		}
		assert.Equal(t, total, 8088)
	})
}

func TestFindTotalBadges(t *testing.T) {
	t.Run("should handle the default input", func(t *testing.T) {
		data, _ := a.ReadDataBytes("testdata/rucksack/input.txt")
		total, err := a.FindTotalBadges(data)

		if err != nil {
			t.Errorf("encountered error %v", err)
		}
		assert.Equal(t, total, 2522)
	})

	t.Run("should find a duplicate badge", func(t *testing.T) {
		data, _ := a.ReadDataBytes("testdata/rucksack/dup_badge_input.txt")
		_, err := a.FindTotalBadges(data)

		assert.NotEqual(t, err, nil)
	})
}
