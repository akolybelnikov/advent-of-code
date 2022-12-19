package day_15

import (
	utils "github.com/akolybelnikov/advent-of-code"
	"sync"
)

// For each line:
// Parse input line and calculate the manhattan distance from the sensor to the beacon
// Find all # positions by applying the manhattan length (CoverageForRow area) and search row index
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
type Interval struct {
	Min, Max int
}

type Coverage []*Interval

func (c *Coverage) Add(interval *Interval) {
	newSet := make([]*Interval, 0)
	finalSet := make([]*Interval, 0)
	i := 0

	for i < len(*c) && (*c)[i].Max < interval.Min {
		newSet = append(newSet, (*c)[i])
		i++
	}

	newSet = append(newSet, interval)

	for i < len(*c) {
		var last = newSet[len(newSet)-1]
		if (*c)[i].Min <= last.Max {
			newInterval := &Interval{
				Min: utils.Min(last.Min, (*c)[i].Min),
				Max: utils.Max(last.Max, (*c)[i].Max),
			}
			newSet[len(newSet)-1] = newInterval
		} else {
			finalSet = append(finalSet, (*c)[i])
		}
		i++
	}

	*c = append(newSet, finalSet...)
}

func (c *Coverage) Merge() {
	newSet := make([]*Interval, 0)
	var last = (*c)[0]
	i := 1

	for i < len(*c) {
		if last.Max+1 >= (*c)[i].Min {
			last.Max = (*c)[i].Max
		} else {
			newSet = append(newSet, last)
			last = (*c)[i]
		}
		i++
	}

	newSet = append(newSet, last)
	*c = newSet
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
		c := pair.CoverageForRow(y)
		for _, pos := range *c {
			covered[pos]++
		}
	}

	return len(covered)
}

func FindCoverageWithLimit(pairs *[]*Pair, limit int) int {
	for i := 0; i <= limit; i++ {
		c := make(Coverage, 0)
		for _, pair := range *pairs {
			interval := pair.CoverageForRowWithLimit(i, limit)
			if interval != nil {
				c.Add(interval)
			}
		}
		c.Merge()
		if len(c) > 1 {
			return ((c[0].Max + 1) * 4000000) + i
		}
	}

	return 0
}

func FindCoverageWithLimitConcurrently(pairs *[]*Pair, limit int) int {
	fqs := make(chan int)
	var res int

	for i := 0; i <= limit; i += 1000 {
		go func(fq chan int, min, max int) {
			for j := min; j <= max; j++ {
				c := make(Coverage, 0)
				for _, pair := range *pairs {
					interval := pair.CoverageForRowWithLimit(j, limit)
					if interval != nil {
						c.Add(interval)
					}
				}
				c.Merge()
				if len(c) > 1 {
					fq <- ((c[0].Max + 1) * 4000000) + j
				}
			}
		}(fqs, i, i+999)
	}

	res = <-fqs
	return res
}

func FindCoverageForRowSyncMap(pairs *[]*Pair, y int) int {
	var wg sync.WaitGroup
	var m sync.Map
	var res int

	for _, pair := range *pairs {
		wg.Add(1)
		pair := pair
		go func() {
			c := pair.CoverageForRow(y)
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
			c := pair.CoverageForRow(y)
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

func (p *Pair) CoverageForRow(y int) *[]int {
	ps := make([]int, 0)
	md := p.ManhattanDist()

	m := md - utils.Abs(p[0].Y-y)
	if m >= 0 {
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
	}

	return &ps
}

func (p *Pair) CoverageForRowWithLimit(y int, limit int) *Interval {
	md := p.ManhattanDist()
	m := md - utils.Abs(p[0].Y-y)
	iMin, iMax := p[0].X-m, p[0].X+m
	if iMin > limit || iMax < 0 {
		return nil
	}
	if m >= 0 {
		return &Interval{
			Min: utils.Max(0, iMin),
			Max: utils.Min(limit, iMax),
		}
	}

	return nil
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
