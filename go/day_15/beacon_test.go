package day_15_test

import (
	"fmt"
	utils "github.com/akolybelnikov/advent-of-code"
	"github.com/akolybelnikov/advent-of-code/day_15"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInputBeacons(t *testing.T) {
	data, _ := utils.ReadDataBytes("testdata/input_test.txt")
	arr, _ := utils.MakeBytesArray(&data)
	pairs := day_15.ProcessData(arr)

	t.Run("test number of pairs created", func(t *testing.T) {
		assert.Equal(t, 14, len(*pairs))
	})

	t.Run("test coverage generated for a sensor", func(t *testing.T) {
		p := day_15.Pair{
			&day_15.Pos{X: 8, Y: 7},  // sensor
			&day_15.Pos{X: 2, Y: 10}, // beacon
		}
		assert.Equal(t, 9, p.ManhattanDist())
		assert.Equal(t, 12, len(*p.CoverageForRow(10)))
		assert.Equal(t, 19, len(*p.CoverageForRow(7)))
		assert.Equal(t, 1, len(*p.CoverageForRow(16)))
		assert.Equal(t, 1, len(*p.CoverageForRow(-2)))
		assert.Equal(t, 3, len(*p.CoverageForRow(-1)))
		assert.Equal(t, 0, len(*p.CoverageForRow(-3)))
	})

	t.Run("test coverage generated for a pair with negative area", func(t *testing.T) {
		p := day_15.Pair{
			&day_15.Pos{X: -6, Y: -6}, // sensor
			&day_15.Pos{X: -3, Y: -4}, // beacon
		}
		assert.Equal(t, 5, p.ManhattanDist())
		assert.Equal(t, 1, len(*p.CoverageForRow(-1)))
		assert.Equal(t, 11, len(*p.CoverageForRow(-6)))
		assert.Equal(t, 1, len(*p.CoverageForRow(-11)))
		assert.Equal(t, 0, len(*p.CoverageForRow(0)))
	})

	t.Run("test for coverage by all sensors in the index row", func(t *testing.T) {
		covered := day_15.FindCoverageForRow(pairs, 10)
		assert.Equal(t, 26, covered)
	})

	t.Run("test for coverage concurrently", func(t *testing.T) {
		covered := day_15.FindCoverageForRowConcurrent(pairs, 10)
		assert.Equal(t, 26, covered)
	})
}

func TestFindCoverageForRow(t *testing.T) {
	data, _ := utils.ReadDataBytes("testdata/input.txt")
	arr, _ := utils.MakeBytesArray(&data)
	pairs := day_15.ProcessData(arr)

	t.Run("test number of pairs created", func(t *testing.T) {
		assert.Equal(t, 35, len(*pairs))
	})

	t.Run("test for coverage by all sensors in the index row", func(t *testing.T) {
		start := time.Now()
		covered := day_15.FindCoverageForRow(pairs, 2000000)
		fmt.Println(time.Since(start))
		assert.Equal(t, 5367037, covered)
		//})
		//
		//t.Run("test for coverage by all sensors in the index row", func(t *testing.T) {
		start = time.Now()
		covered = day_15.FindCoverageForRowConcurrent(pairs, 2000000)
		fmt.Println(time.Since(start))
		assert.Equal(t, 5367037, covered)
		//})
		//
		//t.Run("test for coverage by all sensors in the index row", func(t *testing.T) {
		start = time.Now()
		covered = day_15.FindCoverageForRowSyncMap(pairs, 2000000)
		fmt.Println(time.Since(start))
		assert.Equal(t, 5367037, covered)
	})
}

func TestPair_CoverageForRowWithLimit(t *testing.T) {
	p := day_15.Pair{
		&day_15.Pos{X: 8, Y: 7},
		&day_15.Pos{X: 2, Y: 10},
	}
	i := p.CoverageForRowWithLimit(7, 20)
	assert.EqualValues(t, day_15.Interval{Min: 0, Max: 17}, *i)
}

func TestFindCoverageWithLimit_Test(t *testing.T) {
	data, _ := utils.ReadDataBytes("testdata/input_test.txt")
	arr, _ := utils.MakeBytesArray(&data)
	pairs := day_15.ProcessData(arr)
	fq := day_15.FindCoverageWithLimit(pairs, 20)
	assert.Equal(t, 56000011, fq)
}

func TestFindCoverageWithLimit(t *testing.T) {
	t.Run("single thread", func(t *testing.T) {
		start := time.Now()
		data, _ := utils.ReadDataBytes("testdata/input.txt")
		arr, _ := utils.MakeBytesArray(&data)
		pairs := day_15.ProcessData(arr)
		fq := day_15.FindCoverageWithLimit(pairs, 4000000)
		fmt.Println(time.Since(start))
		assert.Equal(t, 11914583249288, fq)
	})

	t.Run("concurrently", func(t *testing.T) {
		start := time.Now()
		data, _ := utils.ReadDataBytes("testdata/input.txt")
		arr, _ := utils.MakeBytesArray(&data)
		pairs := day_15.ProcessData(arr)
		fq := day_15.FindCoverageWithLimitConcurrently(pairs, 4000000)
		fmt.Println(time.Since(start))
		assert.Equal(t, 11914583249288, fq)
	})
}
