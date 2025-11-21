package advent

import dijkstrastructs "github.com/kirves/godijkstra/common/structs"

type edgeWeight func(string, string) int

type Graph struct {
	Nodes        map[string]any
	Edges        map[string]map[string]any
	ReverseEdges map[string]map[string]any
}

func NewGraph() *Graph {
	return &Graph{
		Nodes:        make(map[string]any),
		Edges:        make(map[string]map[string]any),
		ReverseEdges: make(map[string]map[string]any),
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
