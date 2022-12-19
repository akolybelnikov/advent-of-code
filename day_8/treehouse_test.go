package day_8_test

import (
	utils "github.com/akolybelnikov/advent-of-code"
	"github.com/akolybelnikov/advent-of-code/day_8"
	"testing"
)

func TestMakeGrid(t *testing.T) {
	t.Run("make a grid from input", func(t *testing.T) {
		data, _ := utils.ReadDataBytes("testdata/input.txt")
		grid, err := day_8.MakeTreesGrid(&data)
		if err != nil {
			t.Errorf("received an error while making grid: %v\n", err)
		}
		t.Logf("grid size is %dx%d", len(*grid), len((*grid)[0]))
		t.Logf("num of trees visible: %d", day_8.AllVisibleTrees(grid))
		scores := day_8.AllScenicScores(grid)
		t.Logf("out of %d scores the highest is %d", len(*scores), (*scores)[0])
	})
}
