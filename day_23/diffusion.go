package day_23

import (
	"github.com/mitchellh/hashstructure/v2"
	"github.com/puzpuzpuz/xsync"
	"sync"
)

const (
	EMPTY State = iota
	ELF
)

const (
	N = iota
	S
	W
	E
	NW
	NE
	SE
	SW
)

type State uint8

type Cell struct {
	State    State
	Proposed xsync.Counter
	*Point
}

type Point struct {
	X int32
	Y int32
}

type Grid struct {
	Cells sync.Map
	cur   int
	rules map[int][]int
}

func UnstableDiffusion(arr *[]*[]byte) int {
	g := NewGrid(arr)

	for i := 0; i < 10; i++ {
		elvesCanMove := g.Propose(i)

		if len(elvesCanMove) == 0 {
			break
		}

		g.Move(elvesCanMove)
	}

	return g.getEmptyCellCount()
}

func UnstableDiffusion2(arr *[]*[]byte) int {
	g := NewGrid(arr)
	var rounds int

	for {
		elvesCanMove := g.Propose(rounds)
		rounds++

		if len(elvesCanMove) == 0 {
			break
		}

		g.Move(elvesCanMove)
	}

	return rounds
}

func NewCell(p *Point) *Cell {
	return &Cell{
		Point: p,
	}
}

func NewGrid(arr *[]*[]byte) *Grid {
	g := &Grid{
		rules: map[int][]int{
			N: {N, NW, NE},
			S: {S, SW, SE},
			W: {W, NW, SW},
			E: {E, NE, SE},
		},
	}

	g.makeCells(arr)

	return g
}

func GetHash(p Point) uint64 {
	hash, _ := hashstructure.Hash(p, hashstructure.FormatV2, nil)
	return hash
}

func (g *Grid) makeCells(arr *[]*[]byte) {
	var wg sync.WaitGroup

	for i, v := range *arr {
		for i2, b := range *v {
			wg.Add(1)
			go func(idx, idx2 int, bt byte) {
				defer wg.Done()
				p := Point{X: int32(idx2), Y: int32(idx)}
				c := NewCell(&p)
				if bt == '#' {
					c.State = ELF
				}
				h := GetHash(p)
				g.Cells.Store(h, c)
			}(i, i2, b)
		}
	}

	wg.Wait()
}

func (g *Grid) getMinRectangular() [2]Point {
	var left, right, top, bottom int32
	g.Cells.Range(func(k, v any) bool {
		c := v.(*Cell)
		if c.State == ELF {
			if c.X < left {
				left = c.X
			}
			if c.X > right {
				right = c.X
			}
			if c.Y < top {
				top = c.Y
			}
			if c.Y > bottom {
				bottom = c.Y
			}
		}
		return true
	})

	return [2]Point{
		{X: left, Y: top},
		{X: right, Y: bottom},
	}
}

func (g *Grid) getEmptyCellCount() int {
	var count int
	rect := g.getMinRectangular()
	for i := rect[0].Y; i <= rect[1].Y; i++ {
		for j := rect[0].X; j <= rect[1].X; j++ {
			p := Point{X: j, Y: i}
			h := GetHash(p)
			if v, ok := g.Cells.Load(h); ok {
				c := v.(*Cell)
				if c.State == EMPTY {
					count++
				}
			} else {
				count++
			}
		}
	}

	return count
}

func (g *Grid) hasNeighborElves(ps *[8]Point) byte {
	var bt byte
	for i, p := range *ps {
		h := GetHash(p)
		if v, ok := g.Cells.Load(h); ok {
			c := v.(*Cell)
			if c.State == ELF {
				bt |= byte(1 << i)
			}
		}
	}

	return bt
}

func (p *Point) getNeighborCells() [8]Point {
	return [8]Point{
		{X: p.X, Y: p.Y - 1},     //N
		{X: p.X, Y: p.Y + 1},     //S
		{X: p.X - 1, Y: p.Y},     //W
		{X: p.X + 1, Y: p.Y},     //E
		{X: p.X - 1, Y: p.Y - 1}, //NW
		{X: p.X + 1, Y: p.Y - 1}, //NE
		{X: p.X + 1, Y: p.Y + 1}, //SE
		{X: p.X - 1, Y: p.Y + 1}, //SW

	}
}

func (g *Grid) proposeOne(p Point) *[2]Point {
	nc := p.getNeighborCells()
	hn := g.hasNeighborElves(&nc)
	if hn > 0 {
		d := g.cur
		for i := 0; i < 4; i++ {
			rules := g.rules[d]
			if hn&byte(1<<rules[0]) == 0 && hn&byte(1<<rules[1]) == 0 && hn&byte(1<<rules[2]) == 0 {
				cell := NewCell(&nc[d])
				h := GetHash(nc[d])
				v, _ := g.Cells.LoadOrStore(h, cell)
				c := v.(*Cell)
				c.Proposed.Inc()
				return &[2]Point{p, nc[d]}
			}
			d = (d + 1) % 4
		}
	}

	return nil
}

func (g *Grid) Propose(iteration int) []*[2]Point {
	proposed := make(chan *[2]Point)
	wg := sync.WaitGroup{}

	g.cur = iteration % 4
	g.Cells.Range(func(k, v any) bool {
		c := v.(*Cell)
		if c.State == ELF {
			wg.Add(1)
			go func(p *Point) {
				defer wg.Done()
				move := g.proposeOne(*p)
				if move != nil {
					proposed <- move
				}
			}(c.Point)
		}
		return true
	})

	go func() {
		wg.Wait()
		close(proposed)
	}()

	elvesCanMove := make([]*[2]Point, 0)
	for move := range proposed {
		elvesCanMove = append(elvesCanMove, move)
	}

	return elvesCanMove
}

func (g *Grid) moveOne(move *[2]Point) {
	h1 := GetHash(move[1])
	if v, ok := g.Cells.Load(h1); ok {
		c := v.(*Cell)
		if c.Proposed.Value() == 1 {
			c.State = ELF
			h0 := GetHash(move[0])
			ve, _ := g.Cells.Load(h0)
			e := ve.(*Cell)
			e.State = EMPTY
		}
	}
}

func (g *Grid) Move(moves []*[2]Point) {
	wg := sync.WaitGroup{}
	for _, move := range moves {
		wg.Add(1)
		go func(m *[2]Point) {
			defer wg.Done()
			g.moveOne(m)
		}(move)
	}
	wg.Wait()

	for _, move := range moves {
		h := GetHash(move[1])
		if v, ok := g.Cells.Load(h); ok {
			c := v.(*Cell)
			c.Proposed.Reset()
		}
	}
}
