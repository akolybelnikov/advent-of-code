package day_20_test

import (
	_ "embed"
	"fmt"
	utils "github.com/akolybelnikov/advent-of-code"
	"github.com/akolybelnikov/advent-of-code/day_20"
	"github.com/go-playground/assert/v2"
	"testing"
	"time"
)

//go:embed testdata/input_test.txt
var testInput []byte

//go:embed testdata/input.txt
var input []byte

func TestGrovePositioning(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		start := time.Now()
		dataArr, _ := utils.MakeBytesArray(&testInput)
		grove := day_20.GrovePositioning(dataArr)
		fmt.Println(time.Since(start))
		assert.Equal(t, grove, 3)
	})

	t.Run("input", func(t *testing.T) {
		start := time.Now()
		dataArr, _ := utils.MakeBytesArray(&input)
		grove := day_20.GrovePositioning(dataArr)
		fmt.Println(time.Since(start))
		assert.Equal(t, grove, 14888)
	})
}
