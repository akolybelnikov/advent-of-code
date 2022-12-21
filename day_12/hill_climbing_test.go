package day_12_test

import (
	utils "github.com/akolybelnikov/advent-of-code"
	"github.com/akolybelnikov/advent-of-code/day_12"
	dijkstrapath "github.com/kirves/godijkstra/common/path"
	"github.com/kirves/godijkstra/dijkstra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortInput(t *testing.T) {
	bytes, err := utils.ReadDataBytes("testdata/short_input.txt")
	assert.NoError(t, err)
	grid := day_12.MakeHeightMapGrid(&bytes)
	assert.Equal(t, 5, len(*grid))
	assert.Equal(t, 8, len(*(*grid)[0]))
	graph, s, e, sp := day_12.CreateGraph(grid)
	assert.Equal(t, 40, len(graph.Nodes))
	assert.Equal(t, "0:0", s)
	assert.Equal(t, "2:5", e)

	t.Run("Dijkstra bi-directional solution for the test graph", func(t *testing.T) {
		path, valid := dijkstra.SearchPath(graph, s, e, dijkstra.BIDIR)
		assert.Equal(t, true, valid)
		assert.Equal(t, 31, path.Weight)
	})

	t.Run("Multi-start solution for the test graph", func(t *testing.T) {
		paths := make([]dijkstrapath.DijkstraPath, 0)
		for _, p := range *sp {
			if path, valid := dijkstra.SearchPath(graph, p, e, dijkstra.BIDIR); valid {
				paths = append(paths, path)
			}
		}
		assert.Equal(t, 6, len(paths))
		var shortest int
		for i, p := range paths {
			if i == 0 || p.Weight < shortest {
				shortest = p.Weight
			}
		}
		assert.Equal(t, 29, shortest)
	})
}

func TestLongInput(t *testing.T) {
	bytes, err := utils.ReadDataBytes("testdata/input.txt")
	assert.NoError(t, err)
	grid := day_12.MakeHeightMapGrid(&bytes)
	assert.Equal(t, 41, len(*grid))
	assert.Equal(t, 67, len(*(*grid)[0]))
	graph, s, e, sp := day_12.CreateGraph(grid)
	assert.Equal(t, 2747, len(graph.Nodes))
	assert.Equal(t, "20:0", s)
	assert.Equal(t, "20:43", e)

	t.Run("Dijkstra vanilla solution for the test graph", func(t *testing.T) {
		path, valid := dijkstra.SearchPath(graph, s, e, dijkstra.VANILLA)
		assert.Equal(t, true, valid)
		assert.Equal(t, 381, path.Weight)
	})

	t.Run("Dijkstra bi-directional solution for the test graph", func(t *testing.T) {
		path, valid := dijkstra.SearchPath(graph, s, e, dijkstra.BIDIR)
		assert.Equal(t, true, valid)
		assert.Equal(t, 383, path.Weight)
	})

	t.Run("Multi-start solution for the test graph", func(t *testing.T) {
		paths := make([]dijkstrapath.DijkstraPath, 0)
		for _, p := range *sp {
			if path, valid := dijkstra.SearchPath(graph, p, e, dijkstra.BIDIR); valid {
				paths = append(paths, path)
			}

		}
		assert.Equal(t, 168, len(paths))
		var shortest int
		for i, p := range paths {
			if i == 0 || p.Weight < shortest {
				shortest = p.Weight
			}
		}
		assert.Equal(t, 377, shortest)
	})
}
