package advent_test

import (
	a "scripts/advent"
	"testing"
)

func TestFindFirstMarker(t *testing.T) {

	t.Run("should find first markers in short input", func(t *testing.T) {
		data, err := a.ReadDataBytes("testdata/tuning/short_input.txt")
		if err != nil {
			t.Logf("error reading bytes: %v", err)
		}
		indices := a.HandleMarkers(&data, a.FindFirstMarker)
		for _, idx := range *indices {
			t.Logf("first marker: %d", idx)
		}
	})

	t.Run("should find first messages in short input", func(t *testing.T) {
		data, err := a.ReadDataBytes("testdata/tuning/short_input.txt")
		if err != nil {
			t.Logf("error reading bytes: %v", err)
		}
		indices := a.HandleMarkers(&data, a.FindFirstMessage)
		for _, idx := range *indices {
			t.Logf("first message: %d", idx)
		}
	})

	t.Run("should find first marker in default input", func(t *testing.T) {
		data, err := a.ReadDataBytes("testdata/tuning/input.txt")
		if err != nil {
			t.Logf("error reading bytes: %v", err)
		}
		idx := a.FindFirstMarker(&data)
		t.Logf("first marker: %d", idx)
	})

	t.Run("should find first message in default input", func(t *testing.T) {
		data, err := a.ReadDataBytes("testdata/tuning/input.txt")
		if err != nil {
			t.Logf("error reading bytes: %v", err)
		}
		idx := a.FindFirstMessage(&data)
		t.Logf("first message: %d", idx)
	})
}
