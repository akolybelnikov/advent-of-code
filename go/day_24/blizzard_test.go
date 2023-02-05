package day_24_test

import (
	_ "embed"
	utils "github.com/akolybelnikov/advent-of-code"
	"github.com/akolybelnikov/advent-of-code/day_24"
	"github.com/go-playground/assert/v2"
	"testing"
)

//go:embed testdata/input_test.txt
var inputTest []byte

//go:embed testdata/input.txt
var input []byte

func TestMakeBlizzards(t *testing.T) {
	arr, _ := utils.MakeBytesArray(&inputTest)
	blizzards := day_24.MakeBlizzards(arr)
	assert.Equal(t, 19, len(*blizzards))
}

func TestFindPath(t *testing.T) {
	t.Run("test input", func(t *testing.T) {
		arr, _ := utils.MakeBytesArray(&inputTest)
		minutes := day_24.FindPath(arr)
		t.Logf("minutes: %d", minutes)
	})

	t.Run("input", func(t *testing.T) {
		arr, _ := utils.MakeBytesArray(&input)
		minutes := day_24.FindPath(arr)
		t.Logf("minutes: %d", minutes)
	})
}

func TestFindPath2(t *testing.T) {
	t.Run("test input", func(t *testing.T) {
		arr, _ := utils.MakeBytesArray(&inputTest)
		minutes := day_24.FindPath2(arr)
		t.Logf("minutes: %d", minutes)
	})

	t.Run("input", func(t *testing.T) {
		arr, _ := utils.MakeBytesArray(&input)
		minutes := day_24.FindPath2(arr)
		t.Logf("minutes: %d", minutes)
	})
}
