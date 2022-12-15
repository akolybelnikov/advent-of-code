// Package advent Day 8
package advent

import "sort"

func MakeTreesGrid(data *[]byte) (*[][]byte, error) {
	grid := make([][]byte, 0)
	var endOfLine, prevIdx int
	var row []byte

	for byteIndex, b := range *data {
		if b == newline {
			endOfLine++
			if endOfLine == 1 && prevIdx != byteIndex {
				grid = append(grid, row)
				row = make([]byte, 0)
			}
			prevIdx = byteIndex + 1
		} else {
			row = append(row, b)
			endOfLine = 0
		}
	}

	return &grid, nil
}

func AllVisibleTrees(grid *[][]byte) int {
	visible := (len(*grid)-2)*2 + len((*grid)[0])*2

	for rowIndex := 1; rowIndex < len(*grid)-1; rowIndex++ {
		for colIndex := 1; colIndex < len((*grid)[rowIndex])-1; colIndex++ {
			if treeIsVisible(rowIndex, colIndex, grid) {
				visible++
			}
		}
	}

	return visible
}

func AllScenicScores(grid *[][]byte) *[]int {
	scores := make([]int, 0)

	for rowIndex := 1; rowIndex < len(*grid)-1; rowIndex++ {
		for colIndex := 1; colIndex < len((*grid)[rowIndex])-1; colIndex++ {
			scores = append(scores, treeScenicScore(rowIndex, colIndex, grid))
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(scores)))

	return &scores
}

func treeIsVisible(row, col int, grid *[][]byte) bool {
	tree := (*grid)[row][col]
	var left = true
	var top = true
	var right = true
	var bottom = true

	for _, t := range (*grid)[row][:col] {
		if t >= tree {
			left = false
		}
	}

	for _, t := range (*grid)[row][col+1:] {
		if t >= tree {
			right = false
		}
	}

	for _, r := range (*grid)[:row] {
		if r[col] >= tree {
			top = false
		}
	}

	for _, r := range (*grid)[row+1:] {
		if r[col] >= tree {
			bottom = false
		}
	}

	return left || top || right || bottom
}

func treeScenicScore(row, col int, grid *[][]byte) int {
	leftView := 1
	topView := 1
	rightView := 1
	btmView := 1
	tree := (*grid)[row][col]

	if (*grid)[row][col-1] < tree {
		for c := col - 2; c >= 0; c-- {
			leftView++
			if (*grid)[row][c] >= tree {
				break
			}
		}
	}

	if (*grid)[row][col+1] < tree {
		for c := col + 2; c <= len((*grid)[row])-1; c++ {
			rightView++
			if (*grid)[row][c] >= tree {
				break
			}
		}
	}

	if (*grid)[row-1][col] < tree {
		for r := row - 2; r >= 0; r-- {
			topView++
			if (*grid)[r][col] >= tree {
				break
			}
		}
	}

	if (*grid)[row+1][col] < tree {
		for r := row + 2; r <= len(*grid)-1; r++ {
			btmView++
			if (*grid)[r][col] >= tree {
				break
			}
		}
	}

	return leftView * topView * rightView * btmView
}
