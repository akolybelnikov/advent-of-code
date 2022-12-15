package advent_test

import (
	a "scripts/advent"
	"testing"
)

func TestVisitedPositions(t *testing.T) {
	t.Run("count visited positions with the opTest input", func(t *testing.T) {
		data, _ := a.ReadDataBytes("testdata/rope_bridge/short_input.txt")
		visited, err := a.VisitedPositions(&data, 2)
		if err != nil {
			t.Errorf("encountered an error: %v", err)
		}
		t.Logf("tail visited %d positions", visited)
	})

	t.Run("count visited positions with the long input", func(t *testing.T) {
		data, _ := a.ReadDataBytes("testdata/rope_bridge/input.txt")
		visited, err := a.VisitedPositions(&data, 2)
		if err != nil {
			t.Errorf("encountered an error: %v", err)
		}
		t.Logf("tail visited %d positions", visited)
	})
}

func TestVisitedPositionsLastKnot(t *testing.T) {
	t.Run("count visited positions with the opTest input", func(t *testing.T) {
		data, _ := a.ReadDataBytes("testdata/rope_bridge/short_input.txt")
		visited, err := a.VisitedPositions(&data, 10)
		if err != nil {
			t.Errorf("encountered an error: %v", err)
		}
		t.Logf("tail visited %d positions", visited)
	})

	t.Run("count visited positions with the larger opTest input", func(t *testing.T) {
		data, _ := a.ReadDataBytes("testdata/rope_bridge/large_input.txt")
		visited, err := a.VisitedPositions(&data, 10)
		if err != nil {
			t.Errorf("encountered an error: %v", err)
		}
		t.Logf("tail visited %d positions", visited)
	})

	t.Run("count visited positions with the full opTest input", func(t *testing.T) {
		data, _ := a.ReadDataBytes("testdata/rope_bridge/input.txt")
		visited, err := a.VisitedPositions(&data, 10)
		if err != nil {
			t.Errorf("encountered an error: %v", err)
		}
		t.Logf("tail visited %d positions", visited)
	})
}
