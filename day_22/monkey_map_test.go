package day_22_test

import (
	_ "embed"
	utils "github.com/akolybelnikov/advent-of-code"
	"github.com/akolybelnikov/advent-of-code/day_22"
	"testing"
)

//go:embed testdata/input_test.txt
var testInput []byte

//go:embed testdata/input.txt
var input []byte

func TestMonkeyMap(t *testing.T) {

	t.Run("test input", func(t *testing.T) {
		arr, _ := utils.MakeBytesArray(&testInput)
		pswd := day_22.MonkeyMap(arr)
		t.Logf("password is: %d", pswd)
	})

	t.Run("long input", func(t *testing.T) {
		arr, _ := utils.MakeBytesArray(&input)
		day_22.MonkeyMap(arr)
		pswd := day_22.MonkeyMap(arr)
		t.Logf("password is: %d", pswd)
	})
}
