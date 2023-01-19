package day_23_test

import (
	_ "embed"
	utils "github.com/akolybelnikov/advent-of-code"
	"github.com/akolybelnikov/advent-of-code/day_23"
	"github.com/go-playground/assert/v2"
	"sort"
	"testing"
	"time"
)

//go:embed testdata/small_input_test.txt
var smallTestInput []byte

//go:embed testdata/input_test.txt
var testInput []byte

//go:embed testdata/input.txt
var input []byte

type testCase struct {
	point              day_23.Point
	proposedBeforeMove int64
	proposedAfterMove  int64
	stateBeforeMove    day_23.State
	stateAfterMove     day_23.State
}

func TestDiffusion(t *testing.T) {
	t.Run("Part 1 test input", func(t *testing.T) {
		start := time.Now()
		arr, _ := utils.MakeBytesArray(&testInput)
		res := day_23.UnstableDiffusion(arr)
		t.Logf("Took %s", time.Since(start))

		assert.Equal(t, 110, res)
	})

	t.Run("Part 1 long input", func(t *testing.T) {
		start := time.Now()
		arr, _ := utils.MakeBytesArray(&input)
		res := day_23.UnstableDiffusion(arr)
		t.Logf("Took %s", time.Since(start))

		assert.Equal(t, 3684, res)
	})
}

func TestDiffusion2(t *testing.T) {
	t.Run("Part 2 test input", func(t *testing.T) {
		start := time.Now()
		arr, _ := utils.MakeBytesArray(&testInput)
		res := day_23.UnstableDiffusion2(arr)
		t.Logf("Took %s", time.Since(start))

		assert.Equal(t, 20, res)
	})

	t.Run("Part 2 long input", func(t *testing.T) {
		start := time.Now()
		arr, _ := utils.MakeBytesArray(&input)
		res := day_23.UnstableDiffusion2(arr)
		t.Logf("Took %s", time.Since(start))

		assert.Equal(t, 862, res)
	})
}

func TestSmallInput(t *testing.T) {
	arr, _ := utils.MakeBytesArray(&smallTestInput)
	g := day_23.NewGrid(arr)

	var elves []day_23.Point

	t.Run("grid has cells with elves on the right places", func(t *testing.T) {
		assert.Equal(t, g.Cells.Size(), 30)

		g.Cells.Range(func(p day_23.Point, c *day_23.Cell) bool {
			if c.State == day_23.ELF {
				elves = append(elves, p)
			}
			return true
		})
		assert.Equal(t, len(elves), 5)
		sort.Slice(elves, func(i, j int) bool {
			if elves[i].Y == elves[j].Y {
				return elves[i].X < elves[j].X
			} else {
				return elves[i].Y < elves[j].Y
			}
		})
		assert.Equal(t, elves[0].X, int32(2))
		assert.Equal(t, elves[0].Y, int32(1))
		assert.Equal(t, elves[1].X, int32(3))
		assert.Equal(t, elves[1].Y, int32(1))
		assert.Equal(t, elves[2].X, int32(2))
		assert.Equal(t, elves[2].Y, int32(2))
		assert.Equal(t, elves[3].X, int32(2))
		assert.Equal(t, elves[3].Y, int32(4))
		assert.Equal(t, elves[4].X, int32(3))
		assert.Equal(t, elves[4].Y, int32(4))

		moves := g.Propose(0)
		assert.Equal(t, len(moves), 5)

		testCells := []*testCase{
			{
				point:              day_23.Point{X: 2, Y: 0},
				proposedBeforeMove: 1,
				proposedAfterMove:  0,
				stateBeforeMove:    day_23.EMPTY,
				stateAfterMove:     day_23.ELF,
			},
			{point: day_23.Point{X: 3, Y: 0},
				proposedBeforeMove: 1,
				proposedAfterMove:  0,
				stateBeforeMove:    day_23.EMPTY,
				stateAfterMove:     day_23.ELF,
			},
			{point: day_23.Point{X: 2, Y: 3},
				proposedBeforeMove: 2,
				proposedAfterMove:  0,
				stateBeforeMove:    day_23.EMPTY,
				stateAfterMove:     day_23.EMPTY,
			},
			{point: day_23.Point{X: 3, Y: 3},
				proposedBeforeMove: 1,
				proposedAfterMove:  0,
				stateBeforeMove:    day_23.EMPTY,
				stateAfterMove:     day_23.ELF,
			},
		}

		for _, cell := range testCells {
			if c, ok := g.Cells.Load(cell.point); ok {
				assert.Equal(t, c.Proposed.Value(), cell.proposedBeforeMove)
				assert.Equal(t, c.State, cell.stateBeforeMove)
			}
		}

		g.Move(moves)

		for _, cell := range testCells {
			if c, ok := g.Cells.Load(cell.point); ok {
				assert.Equal(t, c.Proposed.Value(), cell.proposedAfterMove)
				assert.Equal(t, c.State, cell.stateAfterMove)
			}
		}

		moves1 := g.Propose(1)
		assert.Equal(t, len(moves1), 5)

		testCells = []*testCase{
			{
				point:              day_23.Point{X: 2, Y: 1},
				proposedBeforeMove: 1,
				proposedAfterMove:  0,
				stateBeforeMove:    day_23.EMPTY,
				stateAfterMove:     day_23.ELF,
			},
			{
				point:              day_23.Point{X: 3, Y: 1},
				proposedBeforeMove: 1,
				proposedAfterMove:  0,
				stateBeforeMove:    day_23.EMPTY,
				stateAfterMove:     day_23.ELF,
			},
			{
				point:              day_23.Point{X: 1, Y: 2},
				proposedBeforeMove: 1,
				proposedAfterMove:  0,
				stateBeforeMove:    day_23.EMPTY,
				stateAfterMove:     day_23.ELF,
			},
			{
				point:              day_23.Point{X: 4, Y: 3},
				proposedBeforeMove: 1,
				proposedAfterMove:  0,
				stateBeforeMove:    day_23.EMPTY,
				stateAfterMove:     day_23.ELF,
			},
			{
				point:              day_23.Point{X: 2, Y: 5},
				proposedBeforeMove: 1,
				proposedAfterMove:  0,
				stateBeforeMove:    day_23.EMPTY,
				stateAfterMove:     day_23.ELF,
			},
		}

		for _, cell := range testCells {
			if c, ok := g.Cells.Load(cell.point); ok {
				assert.Equal(t, c.Proposed.Value(), cell.proposedBeforeMove)
				assert.Equal(t, c.State, cell.stateBeforeMove)
			}
		}

		g.Move(moves1)

		for _, cell := range testCells {
			if c, ok := g.Cells.Load(cell.point); ok {
				assert.Equal(t, c.Proposed.Value(), cell.proposedAfterMove)
				assert.Equal(t, c.State, cell.stateAfterMove)
			}
		}

		moves2 := g.Propose(2)
		assert.Equal(t, len(moves2), 3)

		testCells = []*testCase{
			{
				point:              day_23.Point{X: 0, Y: 2},
				proposedBeforeMove: 1,
				proposedAfterMove:  0,
				stateBeforeMove:    day_23.EMPTY,
				stateAfterMove:     day_23.ELF,
			},
			{
				point:              day_23.Point{X: 4, Y: 1},
				proposedBeforeMove: 1,
				proposedAfterMove:  0,
				stateBeforeMove:    day_23.EMPTY,
				stateAfterMove:     day_23.ELF,
			},
			{
				point:              day_23.Point{X: 2, Y: 0},
				proposedBeforeMove: 1,
				proposedAfterMove:  0,
				stateBeforeMove:    day_23.EMPTY,
				stateAfterMove:     day_23.ELF,
			},
		}

		for _, cell := range testCells {
			if c, ok := g.Cells.Load(cell.point); ok {
				assert.Equal(t, c.Proposed.Value(), cell.proposedBeforeMove)
				assert.Equal(t, c.State, cell.stateBeforeMove)
			}
		}

		g.Move(moves2)

		for _, cell := range testCells {
			if c, ok := g.Cells.Load(cell.point); ok {
				assert.Equal(t, c.Proposed.Value(), cell.proposedAfterMove)
				assert.Equal(t, c.State, cell.stateAfterMove)
			}
		}
	})
}
