package day_15

import (
	utils "github.com/akolybelnikov/advent-of-code"
	"sync"
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

type CoverageMap struct {
	lock    sync.RWMutex
	covered map[int]int
}

func (m *CoverageMap) addPositions(pos *[]int) {
	m.lock.Lock()
	defer m.lock.Unlock()
	for _, p := range *pos {
		m.covered[p]++
	}
}

func ProcessData(data *[]*[]byte) *[]*Pair {
	pairs := make([]*Pair, 0)
	for _, line := range *data {
		pairs = append(pairs, parseLine(line))
	}

	return &pairs
}

func FindCoverageForRow(pairs *[]*Pair, y int) int {
	covered := make(map[int]int)
	for _, pair := range *pairs {
		c := pair.Coverage(y)
		for _, pos := range *c {
			covered[pos]++
		}
	}

	return len(covered)
}

func FindCoverageForRowSyncMap(pairs *[]*Pair, y int) int {
	var wg sync.WaitGroup
	var m sync.Map
	var res int

	for _, pair := range *pairs {
		wg.Add(1)
		pair := pair
		go func() {
			c := pair.Coverage(y)
			for _, pos := range *c {
				m.Store(pos, struct{}{})
			}
			wg.Done()
		}()
	}

	wg.Wait()

	m.Range(func(k, v any) bool {
		res++
		return true
	})

	return res
}

func FindCoverageForRowConcurrent(pairs *[]*Pair, y int) int {
	var wg sync.WaitGroup
	m := CoverageMap{covered: make(map[int]int)}

	for _, pair := range *pairs {
		wg.Add(1)
		pair := pair
		go func() {
			c := pair.Coverage(y)
			m.addPositions(c)
			wg.Done()
		}()
	}

	wg.Wait()

	return len(m.covered)
}

func (p *Pair) ManhattanDist() int {
	return utils.Abs(p[0].X-p[1].X) + utils.Abs(p[0].Y-p[1].Y)
}

func (p *Pair) Coverage(y int) *[]int {
	ps := make([]int, 0)
	md := p.ManhattanDist()
	// return early if index row not covered at all
	top, btm := p[0].Y-md, p[0].Y+md
	if y < top || y > btm {
		return &ps
	}
	var m int
	// middle row if sensor's Y coordinate is in the index row
	if p[0].Y == y {
		m = md
	}
	// index row in upper pyramid
	if y >= top && y < p[0].Y {
		m = y - p[0].Y + md
		// index row in lower pyramid
	} else {
		m = p[0].Y + md - y
	}

	if p[1].Y == y {
		for i := p[0].X - m; i < p[1].X; i++ {
			ps = append(ps, i)
		}
		for i := p[1].X + 1; i <= p[0].X+m; i++ {
			ps = append(ps, i)
		}
	} else {
		for i := p[0].X - m; i <= p[0].X+m; i++ {
			ps = append(ps, i)
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
