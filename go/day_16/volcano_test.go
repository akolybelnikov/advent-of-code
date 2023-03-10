package day_16_test

import (
	"fmt"
	utils "github.com/akolybelnikov/advent-of-code"
	"github.com/akolybelnikov/advent-of-code/day_16"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestValves(t *testing.T) {
	t.Run("with test input", func(t *testing.T) {
		start := time.Now()
		data, _ := utils.ReadDataBytes("testdata/input_test.txt")
		arr, _ := utils.MakeBytesArray(&data)
		vs := day_16.ParseLines(arr)
		assert.Equal(t, 10, len(*vs))
		flow := day_16.FindMaxFlow(vs)
		assert.Equal(t, 1651, flow)
		fmt.Println(time.Since(start))
	})
	t.Run("with real input", func(t *testing.T) {
		start := time.Now()
		data, _ := utils.ReadDataBytes("testdata/input.txt")
		arr, _ := utils.MakeBytesArray(&data)
		vs := day_16.ParseLines(arr)
		assert.Equal(t, 59, len(*vs))
		flow := day_16.FindMaxFlow(vs)
		assert.Equal(t, 1595, flow)
		fmt.Println(time.Since(start))
	})
}

func TestValves2(t *testing.T) {
	t.Run("with test input", func(t *testing.T) {
		start := time.Now()
		data, _ := utils.ReadDataBytes("testdata/input_test.txt")
		arr, _ := utils.MakeBytesArray(&data)
		vs := day_16.ParseLines(arr)
		assert.Equal(t, 10, len(*vs))
		flow := day_16.FindMaxFlow2(vs)
		assert.Equal(t, 1707, flow)
		fmt.Println(time.Since(start))
	})

	t.Run("with real input", func(t *testing.T) {
		start := time.Now()
		data, _ := utils.ReadDataBytes("testdata/input.txt")
		arr, _ := utils.MakeBytesArray(&data)
		vs := day_16.ParseLines(arr)
		assert.Equal(t, 59, len(*vs))
		flow := day_16.FindMaxFlow2(vs)
		assert.Equal(t, 2186, flow)
		fmt.Println(time.Since(start))
	})
}
