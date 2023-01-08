package day_21_test

import (
	_ "embed"
	"fmt"
	utils "github.com/akolybelnikov/advent-of-code"
	"github.com/akolybelnikov/advent-of-code/day_21"
	"github.com/go-playground/assert/v2"
	"testing"
	"time"
)

//go:embed testdata/input_test.txt
var testInput []byte

//go:embed testdata/input.txt
var input []byte

func TestPart1(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		start := time.Now()
		dataArr, _ := utils.MakeBytesArray(&testInput)
		root := day_21.MonkeyMath(dataArr)
		fmt.Println(time.Since(start))
		assert.Equal(t, root, 152)
	})

	t.Run("input", func(t *testing.T) {
		start := time.Now()
		dataArr, _ := utils.MakeBytesArray(&input)
		root := day_21.MonkeyMath(dataArr)
		fmt.Println(time.Since(start))
		assert.Equal(t, root, 232974643455000)
	})
}
func TestPart2(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		start := time.Now()
		dataArr, _ := utils.MakeBytesArray(&testInput)
		humn := day_21.MonkeyMath2(dataArr)
		fmt.Println(time.Since(start))
		assert.Equal(t, humn, 301)
	})

	t.Run("input", func(t *testing.T) {
		start := time.Now()
		dataArr, _ := utils.MakeBytesArray(&input)
		humn := day_21.MonkeyMath2(dataArr)
		fmt.Println(time.Since(start))
		assert.Equal(t, humn, 3740214169961)
	})
}
