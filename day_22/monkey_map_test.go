package day_22_test

import (
	_ "embed"
	utils "github.com/akolybelnikov/advent-of-code"
	"github.com/akolybelnikov/advent-of-code/day_22"
	"github.com/go-playground/assert/v2"
	"testing"
	"time"
)

//go:embed testdata/input_test.txt
var testInput []byte

//go:embed testdata/input.txt
var input []byte

func TestMonkeyMap(t *testing.T) {

	t.Run("test input", func(t *testing.T) {
		start := time.Now()
		arr, _ := utils.MakeBytesArray(&testInput)
		pswd := day_22.MonkeyMap(arr)
		t.Logf("password = %d, took %s", pswd, time.Since(start))
		assert.Equal(t, 6032, pswd)
	})

	t.Run("long input", func(t *testing.T) {
		start := time.Now()
		arr, _ := utils.MakeBytesArray(&input)
		pswd := day_22.MonkeyMap(arr)
		t.Logf("password = %d, took %s", pswd, time.Since(start))
		assert.Equal(t, 60362, pswd)
	})
}

func TestMonkeyMap2(t *testing.T) {

	t.Run("test input", func(t *testing.T) {
		start := time.Now()
		arr, _ := utils.MakeBytesArray(&testInput)
		pswd := day_22.MonkeyMap2(arr)
		t.Logf("password = %d, took %s", pswd, time.Since(start))
		assert.Equal(t, 5031, pswd)
	})

	t.Run("long input", func(t *testing.T) {
		start := time.Now()
		arr, _ := utils.MakeBytesArray(&input)
		pswd := day_22.MonkeyMap2(arr)
		t.Logf("password = %d, took %s", pswd, time.Since(start))
		assert.Equal(t, 74288, pswd)
	})
}
