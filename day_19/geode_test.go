package day_19_test

import (
	utils "github.com/akolybelnikov/advent-of-code"
	"github.com/akolybelnikov/advent-of-code/day_19"
	"testing"
)

func TestGeodes(t *testing.T) {
	t.Run("with test input", func(t *testing.T) {
		data, _ := utils.ReadDataBytes("testdata/input_test.txt")
		arr, _ := utils.MakeBytesArray(&data)
		blueprints := day_19.MakeBlueprints(arr)
		t.Log(len(blueprints))
	})

	t.Run("with test input", func(t *testing.T) {
		data, _ := utils.ReadDataBytes("testdata/input.txt")
		arr, _ := utils.MakeBytesArray(&data)
		blueprints := day_19.MakeBlueprints(arr)
		t.Log(len(blueprints))
	})
}
