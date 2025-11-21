package main

import (
	"fmt"
	"github.com/akolybelnikov/aoc-2015/internal/utils"
	"os"
	"strconv"
	"strings"
)

type graph [][]int
type node struct {
	name string
	idx  int
	adj  []node
}

const inputFilePath = "inputs/day09.txt"

func main() {
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("Failed to read input file.", err)
		os.Exit(0)
	}
	input := string(data)

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input))

	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input))

	os.Exit(0)
}

// Extracted function to process input and return parsed graph and nodes
func processInput(input string) (map[int]string, graph, error) {
	lines, err := utils.ParseLines(input)
	if err != nil {
		return nil, nil, err
	}
	return newGraph(lines)
}

func part1(input string) int {
	nodes, graph, err := processInput(input)
	utils.HandleErr(err)

	minCost, shortestPath, _, _ := tsp(graph)
	printPath(nodes, shortestPath)
	return minCost
}

func part2(input string) int {
	nodes, graph, err := processInput(input)
	utils.HandleErr(err)

	_, _, maxCost, longestPath := tsp(graph)
	printPath(nodes, longestPath)
	return maxCost
}

func newGraph(lines []string) (map[int]string, graph, error) {
	nodeIndices := make(map[string]int)
	edges := [][3]int{}
	currentIndex := 0

	for _, line := range lines {
		fields := strings.Fields(line)
		var fromIndex, toIndex int

		// Assign indices for nodes if not already mapped
		if idx, exists := nodeIndices[fields[0]]; !exists {
			fromIndex = currentIndex
			currentIndex++
		} else {
			fromIndex = idx
		}

		if idx, exists := nodeIndices[fields[2]]; !exists {
			toIndex = currentIndex
			currentIndex++
		} else {
			toIndex = idx
		}

		dist, err := strconv.Atoi(fields[4])
		if err != nil {
			return nil, nil, err
		}

		edges = append(edges, [3]int{fromIndex, toIndex, dist})
		nodeIndices[fields[0]] = fromIndex
		nodeIndices[fields[2]] = toIndex
	}

	// Build adjacency graph
	adjGraph := make(graph, len(nodeIndices))
	for _, edge := range edges {
		fromIndex, toIndex, dist := edge[0], edge[1], edge[2]
		if adjGraph[fromIndex] == nil {
			adjGraph[fromIndex] = make([]int, len(nodeIndices))
		}
		adjGraph[fromIndex][toIndex] = dist

		if adjGraph[toIndex] == nil {
			adjGraph[toIndex] = make([]int, len(nodeIndices))
		}
		adjGraph[toIndex][fromIndex] = dist
	}

	// Invert the nodeIndices map to create node names map
	nodeNames := make(map[int]string)
	for name, idx := range nodeIndices {
		nodeNames[idx] = name
	}
	return nodeNames, adjGraph, nil
}

func tsp(g [][]int) (int, []int, int, []int) {
	n := len(g)
	const Inf = int(1e9)

	// DP tables for shortest and longest paths
	minDp := make([][]int, 1<<n)
	maxDp := make([][]int, 1<<n)
	parent := make([][]int, 1<<n)
	parentMax := make([][]int, 1<<n)

	for i := 0; i < (1 << n); i++ {
		minDp[i] = make([]int, n)
		maxDp[i] = make([]int, n)
		parent[i] = make([]int, n)
		parentMax[i] = make([]int, n)
		for j := 0; j < n; j++ {
			minDp[i][j] = Inf
			maxDp[i][j] = -Inf
			parent[i][j] = -1
			parentMax[i][j] = -1
		}
	}

	// Initialize DP for each starting node
	for i := 0; i < n; i++ {
		minDp[1<<i][i] = 0
		maxDp[1<<i][i] = 0
	}

	// Main DP loop
	for mask := 1; mask < (1 << n); mask++ {
		for u := 0; u < n; u++ {
			if mask&(1<<u) == 0 {
				continue
			}
			for v := 0; v < n; v++ {
				if mask&(1<<v) != 0 || g[u][v] == 0 {
					continue
				}
				nextMask := mask | (1 << v)
				// Shortest Path DP Update
				newCost := minDp[mask][u] + g[u][v]
				if newCost < minDp[nextMask][v] {
					minDp[nextMask][v] = newCost
					parent[nextMask][v] = u
				}
				// Longest Path DP Update
				newCost = maxDp[mask][u] + g[u][v]
				if newCost > maxDp[nextMask][v] {
					maxDp[nextMask][v] = newCost
					parentMax[nextMask][v] = u
				}
			}
		}
	}

	// Identify final solutions
	allVisited := (1 << n) - 1
	minCost, maxCost := Inf, -Inf
	lastMin, lastMax := -1, -1

	for i := 0; i < n; i++ {
		if minDp[allVisited][i] < minCost {
			minCost = minDp[allVisited][i]
			lastMin = i
		}
		if maxDp[allVisited][i] > maxCost {
			maxCost = maxDp[allVisited][i]
			lastMax = i
		}
	}

	shortestPath := reconstructPath(allVisited, lastMin, parent)
	longestPath := reconstructPath(allVisited, lastMax, parentMax)

	return minCost, shortestPath, maxCost, longestPath
}

// Helper function to reconstruct path
func reconstructPath(finalMask, lastNode int, parent [][]int) []int {
	var path []int
	for current := lastNode; current != -1; {
		path = append([]int{current}, path...)
		next := parent[finalMask][current]
		finalMask ^= 1 << current
		current = next
	}
	return path
}

// Helper function to print the path
func printPath(nodes map[int]string, path []int) {
	cities := make([]string, len(path))
	for i, idx := range path {
		cities[i] = nodes[idx]
	}
	fmt.Println(strings.Join(cities, " -> "))
}
