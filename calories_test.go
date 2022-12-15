package advent_test

import (
	a "github.com/akolybelnikov/advent-of-code"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestElvesStrings(t *testing.T) {
	t.Run("opTest input", func(t *testing.T) {
		total := a.FindCaloriesStrings("testdata/calories/input.txt")
		assert.Equal(t, total, 212117)
	})

	t.Run("opTest empty file", func(t *testing.T) {
		total := a.FindCaloriesStrings("testdata/calories/empty_input.txt")
		assert.Equal(t, total, 0)
	})

	t.Run("opTest malformed input", func(t *testing.T) {
		total := a.FindCaloriesStrings("testdata/calories/malformed_input.txt")
		assert.Equal(t, total, 17)
	})
}

func TestElvesBytes(t *testing.T) {
	t.Run("opTest input bytes", func(t *testing.T) {
		bs, _ := a.ReadDataBytes("testdata/calories/input.txt")
		cal, err := a.FindCaloriesBytes(bs)
		assert.Equal(t, err, nil)
		assert.Equal(t, cal, 212117)
	})

	t.Run("opTest empty bytes", func(t *testing.T) {
		bs, _ := a.ReadDataBytes("testdata/calories/empty_input.txt")
		cal, err := a.FindCaloriesBytes(bs)
		assert.Equal(t, err, nil)
		assert.Equal(t, cal, 0)
	})

	t.Run("opTest malformed bytes", func(t *testing.T) {
		bs, _ := a.ReadDataBytes("testdata/calories/malformed_input.txt")
		cal, err := a.FindCaloriesBytes(bs)
		assert.Equal(t, err, nil)
		assert.Equal(t, cal, 17)
	})
}

func BenchmarkFindCaloriesBytes(b *testing.B) {
	var result int
	bs, _ := a.ReadDataBytes("testdata/calories/input.txt")
	var y int
	for x := 0; x < b.N; x++ {
		y, _ = a.FindCaloriesBytes(bs)
	}
	result = y
	b.Log(result)
}

func BenchmarkFindCaloriesStrings(b *testing.B) {
	var result int
	var y int
	for i := 0; i < b.N; i++ {
		y = a.FindCaloriesStrings("testdata/calories/input.txt")
	}
	result = y
	b.Log(result)
}
