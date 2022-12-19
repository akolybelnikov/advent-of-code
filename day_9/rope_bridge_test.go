package day_9_test

import (
	utils "github.com/akolybelnikov/advent-of-code"
	"github.com/akolybelnikov/advent-of-code/day_9"
	"testing"
)

func TestVisitedPositions(t *testing.T) {
	t.Run("count visited positions with the opTest input", func(t *testing.T) {
		data, _ := utils.ReadDataBytes("testdata/rope_bridge/short_input.txt")
		visited, err := day_9.VisitedPositions(&data, 2)
		if err != nil {
			t.Errorf("encountered an error: %v", err)
		}
		t.Logf("tail visited %d positions", visited)
	})

	t.Run("count visited positions with the long input", func(t *testing.T) {
		data, _ := utils.ReadDataBytes("testdata/rope_bridge/input.txt")
		visited, err := day_9.VisitedPositions(&data, 2)
		if err != nil {
			t.Errorf("encountered an error: %v", err)
		}
		t.Logf("tail visited %d positions", visited)
	})
}

func TestVisitedPositionsLastKnot(t *testing.T) {
	t.Run("count visited positions with the opTest input", func(t *testing.T) {
		data, _ := utils.ReadDataBytes("testdata/rope_bridge/short_input.txt")
		visited, err := day_9.VisitedPositions(&data, 10)
		if err != nil {
			t.Errorf("encountered an error: %v", err)
		}
		t.Logf("tail visited %d positions", visited)
	})

	t.Run("count visited positions with the larger opTest input", func(t *testing.T) {
		data, _ := utils.ReadDataBytes("testdata/rope_bridge/large_input.txt")
		visited, err := day_9.VisitedPositions(&data, 10)
		if err != nil {
			t.Errorf("encountered an error: %v", err)
		}
		t.Logf("tail visited %d positions", visited)
	})

	t.Run("count visited positions with the full opTest input", func(t *testing.T) {
		data, _ := utils.ReadDataBytes("testdata/rope_bridge/input.txt")
		visited, err := day_9.VisitedPositions(&data, 10)
		if err != nil {
			t.Errorf("encountered an error: %v", err)
		}
		t.Logf("tail visited %d positions", visited)
	})
}
