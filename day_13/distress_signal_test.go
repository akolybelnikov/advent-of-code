package day_13_test

import (
	utils "github.com/akolybelnikov/advent-of-code"
	"github.com/akolybelnikov/advent-of-code/day_13"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestCompareTestPackets(t *testing.T) {
	data, _ := utils.ReadDataBytes("testdata/input_test.txt")
	result := day_13.HandlePacketsPart1(&data)

	assert.Equal(t, 13, result)
}

func TestComparePackets(t *testing.T) {
	data, _ := utils.ReadDataBytes("testdata/input.txt")
	result := day_13.HandlePacketsPart1(&data)

	assert.Equal(t, 6101, result)
}

func TestCompareTestPackets2(t *testing.T) {
	data, _ := utils.ReadDataBytes("testdata/input_test.txt")
	result := day_13.HandlePacketsPart2(&data)

	assert.Equal(t, 140, result)
}

func TestComparePackets2(t *testing.T) {
	data, _ := utils.ReadDataBytes("testdata/input.txt")
	result := day_13.HandlePacketsPart2(&data)

	assert.Equal(t, 21909, result)
}

func BenchmarkComparePart1(b *testing.B) {
	data, _ := utils.ReadDataBytes("testdata/input.txt")
	for i := 0; i < b.N; i++ {
		_ = day_13.HandlePacketsPart1(&data)
	}
}

func BenchmarkComparePart2(b *testing.B) {
	data, _ := utils.ReadDataBytes("testdata/input.txt")
	for i := 0; i < b.N; i++ {
		_ = day_13.HandlePacketsPart2(&data)
	}
}
