package day_15

import (
	utils "github.com/akolybelnikov/advent-of-code"
)

// For each line:
// Parse input line and calculate the manhattan distance from the sensor to the beacon
// Find all # positions by applying the manhattan length (Coverage area) and search row index
// Store X coordinates of all # positions for all pairs in a map
const (
	colon = 58
)

type Pos struct {
	X, Y int
}
type Pair [2]*Pos

func ProcessData(data *[]*[]byte) *[]*Pair {
	pairs := make([]*Pair, 0)
	for _, line := range *data {
		pairs = append(pairs, parseLine(line))
	}

	return &pairs
}

func FindCoverageForRow(pairs *[]*Pair, y int) int {
	covered := make(map[int]int, 0)
	for _, pair := range *pairs {
		c := pair.Coverage(y)
		for _, pos := range *c {
			covered[pos]++
		}
	}

	return len(covered)
}

func (p *Pair) ManhattanDist() int {
	return utils.Abs(p[0].X-p[1].X) + utils.Abs(p[0].Y-p[1].Y)
}

func (p *Pair) Coverage(y int) *[]int {
	ps := make([]int, 0)
	md := p.ManhattanDist()
	// middle row only if sensor's Y coordinate is in the index row
	if p[0].Y == y {
		for i := p[0].X - md; i <= p[0].X+md; i++ {
			if p[1].X == i && p[1].Y == y {
				continue
			}
			ps = append(ps, i)
		}
		return &ps
	}

	top, btm := p[0].Y-md, p[0].Y+md
	if y < top || y > btm {
		return &ps
	}

	if y >= top && y < p[0].Y {
		// row in upper pyramid
		m := y - p[0].Y + md
		for j := p[0].X - m; j <= p[0].X+m; j++ {
			if p[1].X == j && p[1].Y == y {
				continue
			}
			ps = append(ps, j)
		}
	} else {
		// row in lower pyramid
		m := p[0].Y + md - y
		for j := p[0].X - m; j <= p[0].X+m; j++ {
			if p[1].X == j && p[1].Y == y {
				continue
			}
			ps = append(ps, j)
		}
	}

	return &ps
}

func parseLine(bytes *[]byte) *Pair {
	var sensor, beacon []byte
	var idx int
	var cur byte
	for cur != colon {
		cur = (*bytes)[idx]
		idx++
	}
	sensor, beacon = (*bytes)[:idx-1], (*bytes)[idx+1:]

	return &Pair{findPos(&sensor), findPos(&beacon)}
}

func findPos(bytes *[]byte) *Pos {
	var startIdx int
	var x, y []byte
	for idx, b := range *bytes {
		if b == 120 {
			startIdx = idx + 2
		}
		if b == 44 {
			x = (*bytes)[startIdx:idx]
		}
		if b == 121 {
			startIdx = idx + 2
		}
		y = (*bytes)[startIdx:]
	}

	return &Pos{
		X: utils.BytesToInt(&x),
		Y: utils.BytesToInt(&y),
	}
}
