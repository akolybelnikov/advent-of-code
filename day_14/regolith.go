package day_14

import (
	utils "github.com/akolybelnikov/advent-of-code"
	"log"
)

const (
	HOR  = 0
	VER  = 1
	AIR  = 46
	ROCK = 35
	SAND = 111
)

type Grid []*[]byte
type step [2]int
type line struct {
	from *step
	to   *step
	is   int
}

func MakeGrid(data *[]*[]byte) (*Grid, int, int) {
	grid := make(Grid, 0)
	var leftEdge, rightEdge int

	for i, dataRow := range *data {
		p := path(dataRow)
		lines := makeLines(p)
		l, r := grid.addPath(lines)
		if i == 0 {
			leftEdge = l
		}
		if l < leftEdge {
			leftEdge = l
		}
		if r > rightEdge {
			rightEdge = r
		}
	}

	return &grid, leftEdge, rightEdge
}

func (g *Grid) addRows(s *step) {
	diff := len(*g) - (s[0] + 1)
	for i := 0; i <= utils.Abs(diff); i++ {
		row := make([]byte, 0)
		*g = append(*g, &row)
	}
}

func (g *Grid) addCols(s *step) {
	row := (*g)[s[0]]
	diff := len(*row) - s[1]
	for i := 0; i <= utils.Abs(diff); i++ {
		*row = append(*row, AIR)
	}
}

func (g *Grid) addPath(lines []*line) (int, int) {
	left := lines[0].from[0]
	right := left
	for _, l := range lines {
		if l.is == HOR {
			s, e := l.from, l.to
			if l.to[1] < l.from[1] {
				s, e = l.to, l.from
			}
			if s[0] < left {
				left = s[0]
			}
			if len(*g) <= s[0] {
				g.addRows(s)
			}
			if len(*(*g)[e[0]]) < e[1] {
				g.addCols(e)
			}
			for i := s[1]; i <= e[1]; i++ {
				(*(*g)[s[0]])[i] = ROCK
			}
		} else {
			s, e := l.from, l.to
			if l.to[0] < l.from[0] {
				s, e = l.to, l.from
			}
			if s[0] < left {
				left = s[0]
			}
			if e[0] > right {
				right = e[0]
			}
			if len(*g) <= e[0] {
				g.addRows(e)
			}

			for ri := s[0]; ri <= e[0]; ri++ {
				if len(*(*g)[ri]) < s[1] {
					st := &step{ri, s[1]}
					g.addCols(st)
				}
				(*(*g)[ri])[s[1]] = ROCK
			}
		}
	}

	return left, right
}

func (g *Grid) DropSand(leftEdge, rightEdge int) int {
	var count, idx2 int
	var idx1 = 100

	for idx1 >= leftEdge && idx1 <= rightEdge {
		count++
		idx1 = 100
		idx2 = 0
		cur := (*(*g)[idx1])[idx2]
		for cur != SAND && (idx1 >= leftEdge && idx1 <= rightEdge) {
			next := (*(*g)[idx1])[idx2+1]
			switch {
			case next == AIR:
				idx2++
			case next == ROCK || next == SAND:
				if idx1 == leftEdge {
					idx1--
					break
				}
				nextLeft := (*(*g)[idx1-1])[idx2+1]
				if nextLeft == AIR {
					idx1--
					idx2++
					continue
				}
				if idx1 == rightEdge {
					idx1++
					break
				}
				nextRight := (*(*g)[idx1+1])[idx2+1]
				if nextRight == AIR {
					idx1++
					idx2++
					continue
				}
				cur = SAND
				(*(*g)[idx1])[idx2] = SAND
			}
		}
	}

	return count - 1
}

func (g *Grid) Render() {
	cols := make([][]byte, 0)
	for _, row := range *g {
		if len(*row) > 0 {
			if len(*row) < len(cols) {
				for i := len(*row); i < len(cols); i++ {
					*row = append(*row, AIR)
				}
			}
			for ci, col := range *row {
				if len(cols) == ci {
					cols = append(cols, make([]byte, 0))
				}
				cols[ci] = append(cols[ci], col)
			}
		}
	}
	for _, l := range cols {
		log.Printf("%s\n", l)
	}
}

func path(line *[]byte) *[]step {
	res := make([]step, 0)
	var s *step
	var cur int
	for i := 0; i < len(*line); i++ {
		if (*line)[i] == 32 && cur < i {
			s = parseStep((*line)[cur:i])
			res = append(res, *s)
			cur = i + 4
		}
	}

	s = parseStep((*line)[cur:])
	res = append(res, *s)

	return &res
}

func parseStep(data []byte) *step {
	var cur, val int
	var res step
	for data[cur] != 44 {
		val = val*10 + int(data[cur]-48)
		cur++
	}
	res[0] = val % 400
	cur++
	val = 0
	for cur < len(data) {
		val = val*10 + int(data[cur]-48)
		cur++
	}
	res[1] = val

	return &res
}

func makeLines(path *[]step) []*line {
	lines := make([]*line, len(*path)-1)
	for i := 0; i < len(*path)-1; i++ {
		from := (*path)[i]
		to := (*path)[i+1]
		l := &line{
			from: &from,
			to:   &to,
			is:   0,
		}
		if from[0] == to[0] {
			l.is = HOR
		} else {
			l.is = VER
		}
		lines[i] = l
	}

	return lines
}
