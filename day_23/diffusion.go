package day_23

import (
	"github.com/mitchellh/hashstructure/v2"
	"github.com/puzpuzpuz/xsync"
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
}

type Point struct {
	X int32
	Y int32
}

type Grid struct {
	Cells xsync.MapOf[Point, *Cell]
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

func NewCell() *Cell {
	return &Cell{}
}

func NewGrid(arr *[]*[]byte) *Grid {
	g := &Grid{
		Cells: *xsync.NewTypedMapOf[Point, *Cell](getHash),
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

func getHash(p Point) uint64 {
	hash, _ := hashstructure.Hash(p, hashstructure.FormatV2, nil)
	return hash
}

func (g *Grid) makeCells(arr *[]*[]byte) {
	for i, v := range *arr {
		for i2, b := range *v {
			p := Point{X: int32(i2), Y: int32(i)}
			c := NewCell()
			if b == '#' {
				c.State = ELF
			}
			g.Cells.Store(p, c)
		}
	}
}

func (g *Grid) getMinRectangular() [2]Point {
	var left, right, top, bottom int32
	g.Cells.Range(func(p Point, c *Cell) bool {
		if c.State == ELF {
			if p.X < left {
				left = p.X
			}
			if p.X > right {
				right = p.X
			}
			if p.Y < top {
				top = p.Y
			}
			if p.Y > bottom {
				bottom = p.Y
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
			if c, ok := g.Cells.Load(p); ok && c.State == EMPTY {
				count++
			} else if !ok {
				count++
			}
		}
	}
	return count
}

func (g *Grid) hasNeighborElves(ps *[8]Point) byte {
	var bt byte
	for i, p := range *ps {
		if c, ok := g.Cells.Load(p); ok {
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

func (g *Grid) proposeOne(p *Point) *[2]Point {
	nc := p.getNeighborCells()
	hn := g.hasNeighborElves(&nc)
	if hn == 0 {
		return nil
	} else {
		d := g.cur
		for i := 0; i < 4; i++ {
			rules := g.rules[d]
			if hn&byte(1<<rules[0]) == 0 && hn&byte(1<<rules[1]) == 0 && hn&byte(1<<rules[2]) == 0 {
				cell := NewCell()
				if c, ok := g.Cells.LoadOrStore(nc[d], cell); ok {
					cell.Proposed.Add(c.Proposed.Value())
				} else {
					cell.Proposed.Inc()
				}

				return &[2]Point{*p, nc[d]}
			}
			d = (d + 1) % 4
		}
	}

	return nil
}

func (g *Grid) Propose(iteration int) []*[2]Point {
	canMove := make([]*[2]Point, 0)
	g.cur = iteration % 4
	g.Cells.Range(func(p Point, c *Cell) bool {
		if c.State == ELF {
			nc := p.getNeighborCells()
			hn := g.hasNeighborElves(&nc)
			if hn > 0 {
				d := g.cur
				for i := 0; i < 4; i++ {
					rules := g.rules[d]
					if hn&byte(1<<rules[0]) == 0 && hn&byte(1<<rules[1]) == 0 && hn&byte(1<<rules[2]) == 0 {
						if cell, ok := g.Cells.Load(nc[d]); ok {
							cell.Proposed.Inc()
						} else {
							cl := NewCell()
							cl.Proposed.Inc()
							g.Cells.Store(nc[d], cl)
						}
						canMove = append(canMove, &[2]Point{p, nc[d]})
						return true
					}
					d = (d + 1) % 4
				}
			}
		}
		return true
	})

	return canMove
}

func (g *Grid) moveOne(move *[2]Point) {
	if c, ok := g.Cells.Load(move[1]); ok {
		if c.Proposed.Value() == 1 {
			c.State = ELF
			e, _ := g.Cells.Load(move[0])
			e.State = EMPTY
		}
	}
}

func (g *Grid) Move(moves []*[2]Point) {
	for _, move := range moves {
		if c, ok := g.Cells.Load(move[1]); ok {
			if c.Proposed.Value() == 1 {
				c.State = ELF
				e, _ := g.Cells.Load(move[0])
				e.State = EMPTY
			}
			c.Proposed.Reset()
		}
	}
}
