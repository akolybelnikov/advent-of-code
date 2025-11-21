package day_2_test

import (
	u "github.com/akolybelnikov/advent-of-code"
	a "github.com/akolybelnikov/advent-of-code/day_2"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestFindMyStrategyTotal(t *testing.T) {
	t.Run("should handle the default input", func(t *testing.T) {
		bs, _ := u.ReadDataBytes("testdata/input.txt")
		total, err := a.FindMyStrategyTotal(bs)
		assert.Equal(t, err, nil)
		assert.Equal(t, total, 12794)
	})

	t.Run("should handle a malformed input", func(t *testing.T) {
		bs, _ := u.ReadDataBytes("testdata/malformed.txt")
		total, err := a.FindMyStrategyTotal(bs)
		assert.Equal(t, err, nil)
		assert.Equal(t, total, 45)
	})
}

func TestFindElfStrategyTotal(t *testing.T) {
	t.Run("should handle the default input", func(t *testing.T) {
		bs, _ := u.ReadDataBytes("testdata/input.txt")
		total, err := a.FindElfStrategyTotal(bs)
		assert.Equal(t, err, nil)
		assert.Equal(t, total, 14979)
	})

	t.Run("should handle a malformed input", func(t *testing.T) {
		bs, _ := u.ReadDataBytes("testdata/malformed.txt")
		total, err := a.FindElfStrategyTotal(bs)
		assert.Equal(t, err, nil)
		assert.Equal(t, total, 51)
	})
}
