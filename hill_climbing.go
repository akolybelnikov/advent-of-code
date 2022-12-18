package advent

import (
	"fmt"
	dijkstrastructs "github.com/kirves/godijkstra/common/structs"
)

type Node struct {
	X, Y int
}

type Graph struct {
	Nodes        map[string]struct{}
	Edges        map[string]map[string]struct{}
	ReverseEdges map[string]map[string]struct{}
}

func NewGraph() *Graph {
	return &Graph{
		Nodes:        make(map[string]struct{}),
		Edges:        make(map[string]map[string]struct{}),
		ReverseEdges: make(map[string]map[string]struct{}),
	}
}

func (g *Graph) SuccessorsForNode(node string) []dijkstrastructs.Connection {
	successors := make([]dijkstrastructs.Connection, len(g.Edges[node]))
	i := 0
	for k, _ := range g.Edges[node] {
		successors[i] = dijkstrastructs.Connection{Destination: k, Weight: g.EdgeWeight(node, k)}
		i++
	}
	return successors

}

func (g *Graph) PredecessorsFromNode(node string) []dijkstrastructs.Connection {
	predecessors := make([]dijkstrastructs.Connection, len(g.ReverseEdges[node]))
	i := 0
	for k, _ := range g.ReverseEdges[node] {
		predecessors[i] = dijkstrastructs.Connection{Destination: k, Weight: g.EdgeWeight(k, node)}
		i++
	}
	return predecessors
}

func (g *Graph) EdgeWeight(_, _ string) int {
	return 1
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
		if b == NEWLINE {
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

func CreateGraph(grid *[]*[]byte) (*Graph, string, string, *[]string) {
	var start, end string
	graph := NewGraph()
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
			edges := make(map[string]struct{})
			e, r := node.Edges(grid)
			for _, n := range e {
				edges[n.StringId()] = struct{}{}
			}
			graph.Edges[node.StringId()] = edges
			reverseEdges := make(map[string]struct{})
			for _, n := range r {
				reverseEdges[n.StringId()] = struct{}{}
			}
			graph.ReverseEdges[node.StringId()] = reverseEdges
		}
	}

	return graph, start, end, &startingPoints
}
