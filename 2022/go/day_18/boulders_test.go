package day_18_test

import (
	utils "github.com/akolybelnikov/advent-of-code"
	"github.com/akolybelnikov/advent-of-code/day_18"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCubes(t *testing.T) {
	t.Run("with test input", func(t *testing.T) {
		data, _ := utils.ReadDataBytes("testdata/input_test.txt")
		arr, _ := utils.MakeBytesArray(&data)
		total := day_18.FindNotConnected(arr)
		assert.Equal(t, 64, total)
	})

	t.Run("with long input", func(t *testing.T) {
		data, _ := utils.ReadDataBytes("testdata/input.txt")
		arr, _ := utils.MakeBytesArray(&data)
		total := day_18.FindNotConnected(arr)
		assert.Equal(t, 4320, total)
	})
}

func TestCubes2(t *testing.T) {
	t.Run("with test input", func(t *testing.T) {
		data, _ := utils.ReadDataBytes("testdata/input_test.txt")
		arr, _ := utils.MakeBytesArray(&data)
		total := day_18.FindNotConnected2(arr)
		assert.Equal(t, 58, total)
	})

	t.Run("with long input", func(t *testing.T) {
		data, _ := utils.ReadDataBytes("testdata/input.txt")
		arr, _ := utils.MakeBytesArray(&data)
		total := day_18.FindNotConnected2(arr)
		assert.Equal(t, 2456, total)
	})
}
