package day_17_test

import (
	"fmt"
	utils "github.com/akolybelnikov/advent-of-code"
	"github.com/akolybelnikov/advent-of-code/day_17"
	"testing"
	"time"
)

func TestFlow(t *testing.T) {
	t.Run("test input", func(t *testing.T) {
		start := time.Now()
		data, _ := utils.ReadDataBytes("testdata/input_test.txt")
		pattern := data[:len(data)-1]
		t.Log(len(pattern))
		top := day_17.Run(&pattern, 2022)
		t.Log(top)
		fmt.Println(time.Since(start))
	})

	t.Run("test input part 2", func(t *testing.T) {
		start := time.Now()
		data, _ := utils.ReadDataBytes("testdata/input_test.txt")
		pattern := data[:len(data)-1]
		t.Log(len(pattern))
		top := day_17.Run2(&pattern, 1000000000000)
		t.Log(top)
		fmt.Println(time.Since(start))
	})

	t.Run("long input", func(t *testing.T) {
		start := time.Now()
		data, _ := utils.ReadDataBytes("testdata/input.txt")
		pattern := data[:len(data)-1]
		t.Log(len(pattern))
		top := day_17.Run(&pattern, 2022)
		t.Log(top)
		fmt.Println(time.Since(start))
	})

	t.Run("long input part 2", func(t *testing.T) {
		start := time.Now()
		data, _ := utils.ReadDataBytes("testdata/input.txt")
		pattern := data[:len(data)-1]
		t.Log(len(pattern))
		top := day_17.Run2(&pattern, 1000000000000)
		t.Log(top)
		fmt.Println(time.Since(start))
	})
}
