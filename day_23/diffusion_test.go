package day_23_test

import (
	_ "embed"
	utils "github.com/akolybelnikov/advent-of-code"
	"github.com/akolybelnikov/advent-of-code/day_23"
	"testing"
)

//go:embed testdata/input_test.txt
var testInput []byte

//go:embed testdata/input.txt
var input []byte

func TestDiffusion(t *testing.T) {
	t.Run("Part 1 test input", func(t *testing.T) {
		arr, _ := utils.MakeBytesArray(&testInput)
		day_23.UnstableDiffusion(arr)
	})
	t.Run("Part 1 long input", func(t *testing.T) {
		arr, _ := utils.MakeBytesArray(&input)
		day_23.UnstableDiffusion(arr)
	})
}
