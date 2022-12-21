package day_12

import (
	"fmt"
	utils "github.com/akolybelnikov/advent-of-code"
)

type Node struct {
	X, Y int
}

func (n *Node) StringId() string {
	return fmt.Sprintf("%d:%d", n.X, n.Y)
}

func (n *Node) Edges(grid *[]*[]byte) (edges []*Node, reverseEdges []*Node) {
	moves := make([]*Node, 0)
	if n.X > 0 {
		moves = append(moves, &Node{n.X - 1, n.Y})
	}
	if n.X < len(*grid)-1 {
		moves = append(moves, &Node{n.X + 1, n.Y})
	}
	if n.Y > 0 {
		moves = append(moves, &Node{n.X, n.Y - 1})
	}
	if n.Y < len(*(*grid)[n.X])-1 {
		moves = append(moves, &Node{n.X, n.Y + 1})
	}

	nodeByte := (*(*grid)[n.X])[n.Y]
	for _, m := range moves {
		moveByte := (*(*grid)[m.X])[m.Y]
		if moveByte <= nodeByte || moveByte-nodeByte == 1 {
			edges = append(edges, m)
		}
		if moveByte >= nodeByte || nodeByte-moveByte == 1 {
			reverseEdges = append(reverseEdges, m)
		}
	}
	return edges, reverseEdges
}

func MakeHeightMapGrid(data *[]byte) *[]*[]byte {
	var grid = make([]*[]byte, 0)
	var line, prevIdx int

	for byteIndex, b := range *data {
		if b == utils.NEWLINE {
			line++
			if line == 1 && prevIdx != byteIndex {
				row := (*data)[prevIdx:byteIndex]
				grid = append(grid, &row)
			}
			prevIdx = byteIndex + 1
		} else {
			line = 0
		}
	}

	return &grid
}

func CreateGraph(grid *[]*[]byte) (*utils.Graph, string, string, *[]string) {
	var start, end string
	graph := utils.NewGraph()
	startingPoints := make([]string, 0)

	for i, row := range *grid {
		for j, col := range *row {
			node := &Node{i, j}
			if col == 97 {
				sp := node.StringId()
				startingPoints = append(startingPoints, sp)
			}
			if col == 83 {
				(*(*grid)[i])[j] = 97
				start = node.StringId()
				startingPoints = append(startingPoints, start)
			}
			if col == 69 {
				(*(*grid)[i])[j] = 122
				end = node.StringId()
			}
			graph.Nodes[node.StringId()] = struct{}{}
			edges := make(map[string]any)
			e, r := node.Edges(grid)
			for _, n := range e {
				edges[n.StringId()] = struct{}{}
			}
			graph.Edges[node.StringId()] = edges
			reverseEdges := make(map[string]any)
			for _, n := range r {
				reverseEdges[n.StringId()] = struct{}{}
			}
			graph.ReverseEdges[node.StringId()] = reverseEdges
		}
	}

	return graph, start, end, &startingPoints
}
