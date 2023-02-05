package day_19_test

import (
	"fmt"
	utils "github.com/akolybelnikov/advent-of-code"
	"github.com/akolybelnikov/advent-of-code/day_19"
	"github.com/go-playground/assert/v2"
	"testing"
	"time"
)

func TestGeodes(t *testing.T) {
	t.Run("with test input", func(t *testing.T) {
		data, _ := utils.ReadDataBytes("testdata/input_test.txt")
		arr, _ := utils.MakeBytesArray(&data)
		blueprints := day_19.MakeBlueprints(arr)
		start := time.Now()
		res := day_19.Run(blueprints, 24)
		fmt.Println(time.Since(start))
		assert.Equal(t, res, 33)
	})

	t.Run("with real input", func(t *testing.T) {
		data, _ := utils.ReadDataBytes("testdata/input.txt")
		arr, _ := utils.MakeBytesArray(&data)
		blueprints := day_19.MakeBlueprints(arr)
		start := time.Now()
		res := day_19.Run(blueprints, 24)
		t.Log(time.Since(start))
		assert.Equal(t, res, 600)
	})
}

func TestGeodes2(t *testing.T) {
	t.Run("with test input", func(t *testing.T) {
		data, _ := utils.ReadDataBytes("testdata/input_test.txt")
		arr, _ := utils.MakeBytesArray(&data)
		blueprints := day_19.MakeBlueprints(arr)
		start := time.Now()
		res := day_19.Run2(blueprints, 32)
		fmt.Println(time.Since(start))
		assert.Equal(t, res, 62*56)
	})

	t.Run("with real input", func(t *testing.T) {
		data, _ := utils.ReadDataBytes("testdata/input.txt")
		arr, _ := utils.MakeBytesArray(&data)
		bps := (*arr)[:3]
		blueprints := day_19.MakeBlueprints(&bps)
		start := time.Now()
		res := day_19.Run2(blueprints, 32)
		t.Log(time.Since(start))
		assert.Equal(t, res, 6000)
	})
}
