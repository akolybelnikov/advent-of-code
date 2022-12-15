package advent_test

import (
	a "scripts/advent"
	"testing"
)

func TestMakeGrid(t *testing.T) {
	t.Run("make a grid from input", func(t *testing.T) {
		data, _ := a.ReadDataBytes("testdata/treehouse/input.txt")
		grid, err := a.MakeTreesGrid(&data)
		if err != nil {
			t.Errorf("received an error while making grid: %v\n", err)
		}
		t.Logf("grid size is %dx%d", len(*grid), len((*grid)[0]))
		t.Logf("num of trees visible: %d", a.AllVisibleTrees(grid))
		scores := a.AllScenicScores(grid)
		t.Logf("out of %d scores the highest is %d", len(*scores), (*scores)[0])
	})
}
