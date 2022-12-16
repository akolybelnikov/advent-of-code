package advent_test

import (
	a "github.com/akolybelnikov/advent-of-code"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestCompareTestPackets(t *testing.T) {
	data, _ := a.ReadDataBytes("testdata/distress_signal/test_input.txt")
	result := a.HandlePacketsPart1(&data)

	assert.Equal(t, 13, result)
}

func TestComparePackets(t *testing.T) {
	data, _ := a.ReadDataBytes("testdata/distress_signal/input.txt")
	result := a.HandlePacketsPart1(&data)

	assert.Equal(t, 6101, result)
}

func TestCompareTestPackets2(t *testing.T) {
	data, _ := a.ReadDataBytes("testdata/distress_signal/test_input.txt")
	result := a.HandlePacketsPart2(&data)

	assert.Equal(t, 140, result)
}

func TestComparePackets2(t *testing.T) {
	data, _ := a.ReadDataBytes("testdata/distress_signal/input.txt")
	result := a.HandlePacketsPart2(&data)

	assert.Equal(t, 21909, result)
}

func BenchmarkComparePart1(b *testing.B) {
	data, _ := a.ReadDataBytes("testdata/distress_signal/input.txt")
	for i := 0; i < b.N; i++ {
		_ = a.HandlePacketsPart1(&data)
	}
}

func BenchmarkComparePart2(b *testing.B) {
	data, _ := a.ReadDataBytes("testdata/distress_signal/input.txt")
	for i := 0; i < b.N; i++ {
		_ = a.HandlePacketsPart2(&data)
	}
}
