package advent_test

import (
	a "github.com/akolybelnikov/advent-of-code"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestFindMyStrategyTotal(t *testing.T) {
	t.Run("should handle the default input", func(t *testing.T) {
		bs, _ := a.ReadDataBytes("testdata/strategy/input.txt")
		total, err := a.FindMyStrategyTotal(bs)
		assert.Equal(t, err, nil)
		assert.Equal(t, total, 12794)
	})

	t.Run("should handle an empty input", func(t *testing.T) {
		bs, _ := a.ReadDataBytes("testdata/calories/empty_input.txt")
		total, err := a.FindMyStrategyTotal(bs)
		assert.Equal(t, err, nil)
		assert.Equal(t, total, 0)
	})

	t.Run("should handle a malformed input", func(t *testing.T) {
		bs, _ := a.ReadDataBytes("testdata/strategy/malformed.txt")
		total, err := a.FindMyStrategyTotal(bs)
		assert.Equal(t, err, nil)
		assert.Equal(t, total, 45)
	})
}

func TestFindElfStrategyTotal(t *testing.T) {
	t.Run("should handle the default input", func(t *testing.T) {
		bs, _ := a.ReadDataBytes("testdata/strategy/input.txt")
		total, err := a.FindElfStrategyTotal(bs)
		assert.Equal(t, err, nil)
		assert.Equal(t, total, 14979)
	})

	t.Run("should handle an empty input", func(t *testing.T) {
		bs, _ := a.ReadDataBytes("testdata/calories/empty_input.txt")
		total, err := a.FindElfStrategyTotal(bs)
		assert.Equal(t, err, nil)
		assert.Equal(t, total, 0)
	})

	t.Run("should handle a malformed input", func(t *testing.T) {
		bs, _ := a.ReadDataBytes("testdata/strategy/malformed.txt")
		total, err := a.FindElfStrategyTotal(bs)
		assert.Equal(t, err, nil)
		assert.Equal(t, total, 51)
	})
}
