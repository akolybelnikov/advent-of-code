package day_6_test

import (
	a "github.com/akolybelnikov/advent-of-code"
	c "github.com/akolybelnikov/advent-of-code/day_6"
	"testing"
)

func TestFindFirstMarker(t *testing.T) {

	t.Run("should find first markers in short input", func(t *testing.T) {
		data, err := a.ReadDataBytes("testdata/short_input.txt")
		if err != nil {
			t.Logf("error reading bytes: %v", err)
		}
		indices := c.HandleMarkers(&data, c.FindFirstMarker)
		for _, idx := range *indices {
			t.Logf("first marker: %d", idx)
		}
	})

	t.Run("should find first messages in short input", func(t *testing.T) {
		data, err := a.ReadDataBytes("testdata/short_input.txt")
		if err != nil {
			t.Logf("error reading bytes: %v", err)
		}
		indices := c.HandleMarkers(&data, c.FindFirstMessage)
		for _, idx := range *indices {
			t.Logf("first message: %d", idx)
		}
	})

	t.Run("should find first marker in default input", func(t *testing.T) {
		data, err := a.ReadDataBytes("testdata/input.txt")
		if err != nil {
			t.Logf("error reading bytes: %v", err)
		}
		idx := c.FindFirstMarker(&data)
		t.Logf("first marker: %d", idx)
	})

	t.Run("should find first message in default input", func(t *testing.T) {
		data, err := a.ReadDataBytes("testdata/input.txt")
		if err != nil {
			t.Logf("error reading bytes: %v", err)
		}
		idx := c.FindFirstMessage(&data)
		t.Logf("first message: %d", idx)
	})
}

func BenchmarkFindFirstMessage(b *testing.B) {
	data, err := a.ReadDataBytes("testdata/input.txt")
	if err != nil {
		b.Logf("error reading bytes: %v", err)
	}
	for i := 0; i < b.N; i++ {
		c.FindFirstMessage(&data)
	}
}
