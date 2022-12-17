package day_14_test

import (
	"fmt"
	a "github.com/akolybelnikov/advent-of-code"
	c "github.com/akolybelnikov/advent-of-code/day_14"
	"github.com/go-playground/assert/v2"
	"testing"
	"time"
)

func TestMakeGridTestInput(t *testing.T) {
	data, _ := a.ReadDataBytes("testdata/test_input.txt")
	bytesArray, _ := a.MakeBytesArray(&data)
	grid, leftEdge, rightEdge := c.MakeGrid(bytesArray)
	count := grid.DropSand(leftEdge, rightEdge)
	grid.Render()
	assert.Equal(t, 24, count)
}

func TestMakeGrid(t *testing.T) {
	start := time.Now()
	data, _ := a.ReadDataBytes("testdata/input.txt")
	bytesArray, _ := a.MakeBytesArray(&data)
	grid, leftEdge, rightEdge := c.MakeGrid(bytesArray)
	count := grid.DropSand(leftEdge, rightEdge)
	fmt.Println(time.Since(start))
	assert.Equal(t, 672, count)
}
