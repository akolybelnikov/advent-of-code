package day_3_test

import (
	a "github.com/akolybelnikov/advent-of-code"
	c "github.com/akolybelnikov/advent-of-code/day_3"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestFindTotalPriorities(t *testing.T) {
	t.Run("should handle default input", func(t *testing.T) {
		data, _ := a.ReadDataBytes("testdata/input.txt")
		total, err := c.FindTotalPriorities(data)

		if err != nil {
			t.Errorf("encountered error %v", err)
		}
		assert.Equal(t, total, 8088)
	})
}

func TestFindTotalBadges(t *testing.T) {
	t.Run("should handle the default input", func(t *testing.T) {
		data, _ := a.ReadDataBytes("testdata/input.txt")
		total, err := c.FindTotalBadges(data)

		if err != nil {
			t.Errorf("encountered error %v", err)
		}
		assert.Equal(t, total, 2522)
	})

	t.Run("should find a duplicate badge", func(t *testing.T) {
		data, _ := a.ReadDataBytes("testdata/dup_badge_input.txt")
		_, err := c.FindTotalBadges(data)

		assert.NotEqual(t, err, nil)
	})
}
