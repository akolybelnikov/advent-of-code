package day_23

import (
	"github.com/mitchellh/hashstructure/v2"
	"github.com/puzpuzpuz/xsync"
	"log"
)

const (
	EMPTY state = iota
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

type state uint8

type cell struct {
	state    state
	dir      int
	proposed int
}

type point struct {
	x int32
	y int32
}

type grid struct {
	cells xsync.MapOf[point, *cell]
	edges [2]point
	cur   int
	rules map[int][]int
}

func UnstableDiffusion(arr *[]*[]byte) int {
	g := newGrid()
	g.makeCells(arr)

	log.Printf("Initial grid size: %d x %d", g.getSizeX(), g.getSizeY())
	ps := g.getMinRectangular()
	log.Printf("Minimal containing rectangular: %v %v", ps[0], ps[1])
	log.Printf("Empty cells count: %d", g.getEmptyCellCount())

	p := point{x: 4, y: 1}
	if c, ok := g.cells.Load(p); ok {
		log.Printf("Cell %v: %v", p, c)
	}

	ns := p.getNeighborCells()
	log.Printf("Neighbors of %v: %v", p, ns)

	hn := g.hasNeighborElves(&ns)
	log.Printf("Has neighbors: %08b = %d", hn, hn)
	log.Printf("Has neighbors: %v", hn > 0)

	for i := 0; i < 8; i++ {
		log.Printf("neighbor %d: %v", i, hn&byte(1<<i) > 0)
	}

	for i := 0; i < 10; i++ {
		g.cur = i % 4
		g.cells.Range(func(p point, c *cell) bool {
			if c.state == ELF {
				cm := g.propose(&p)
				log.Printf("Can move %v | %v | %d ", p, cm, c.dir)
			}
			return true
		})
		log.Printf("Current direction: %d", g.cur)
	}

	return 0
}

func newCell() *cell {
	return &cell{
		state:    EMPTY,
		proposed: 0,
	}
}

func newGrid() *grid {
	g := &grid{
		edges: [2]point{
			{x: 0, y: 0},
			{x: 0, y: 0},
		},
	}

	g.cells = *xsync.NewTypedMapOf[point, *cell](getHash)

	g.rules = make(map[int][]int)
	g.rules[N] = []int{NW, N, NE}
	g.rules[S] = []int{SW, S, SE}
	g.rules[W] = []int{NW, W, SW}
	g.rules[E] = []int{NE, E, SE}

	return g
}

func getHash(p point) uint64 {
	hash, _ := hashstructure.Hash(p, hashstructure.FormatV2, nil)
	return hash
}

func (g *grid) makeCells(arr *[]*[]byte) {
	for i, v := range *arr {
		for i2, b := range *v {
			p := point{x: int32(i2), y: int32(i)}
			c := &cell{}
			if b == '#' {
				c.state = ELF
			}
			g.cells.Store(p, c)
			g.updateEdges(p)
		}
	}
}

func (g *grid) updateEdges(p point) {
	if p.x < g.edges[0].x {
		g.edges[0].x = p.x
	}
	if p.y < g.edges[0].y {
		g.edges[0].y = p.y
	}
	if p.x > g.edges[1].x {
		g.edges[1].x = p.x
	}
	if p.y > g.edges[1].y {
		g.edges[1].y = p.y
	}
}

func (g *grid) getSizeX() int32 {
	return g.edges[1].x - g.edges[0].x + 1
}

func (g *grid) getSizeY() int32 {
	return g.edges[1].y - g.edges[0].y + 1
}

func (g *grid) getMinRectangular() [2]point {
	var left, right, top, bottom int32
	g.cells.Range(func(p point, c *cell) bool {
		if c.state == ELF {
			if p.x < left {
				left = p.x
			}
			if p.x > right {
				right = p.x
			}
			if p.y < top {
				top = p.y
			}
			if p.y > bottom {
				bottom = p.y
			}
		}
		return true
	})

	return [2]point{
		{x: left, y: top},
		{x: right, y: bottom},
	}
}

func (g *grid) getEmptyCellCount() int {
	var count int
	rect := g.getMinRectangular()
	for i := rect[0].y; i <= rect[1].y; i++ {
		for j := rect[0].x; j <= rect[1].x; j++ {
			p := point{x: j, y: i}
			if c, ok := g.cells.Load(p); ok && c.state == EMPTY {
				count++
			} else if !ok {
				count++
			}
		}
	}
	return count
}

func (g *grid) hasNeighborElves(ps *[8]point) byte {
	var bt byte
	for i, p := range *ps {
		if c, ok := g.cells.Load(p); ok {
			if c.state == ELF {
				bt |= byte(1 << i)
			}
		}
	}

	return bt
}

func (p *point) getNeighborCells() [8]point {
	return [8]point{
		{x: p.x, y: p.y - 1},     //N
		{x: p.x, y: p.y + 1},     //S
		{x: p.x - 1, y: p.y},     //W
		{x: p.x + 1, y: p.y},     //E
		{x: p.x - 1, y: p.y - 1}, //NW
		{x: p.x + 1, y: p.y - 1}, //NE
		{x: p.x + 1, y: p.y + 1}, //SE
		{x: p.x - 1, y: p.y + 1}, //SW

	}
}

func (g *grid) propose(p *point) bool {
	nc := p.getNeighborCells()
	hn := g.hasNeighborElves(&nc)
	if hn == 0 {
		return false
	} else {
		d := g.cur
		for i := 0; i < 4; i++ {
			rules := g.rules[d]
			if hn&byte(1<<rules[0]) == 0 && hn&byte(1<<rules[1]) == 0 && hn&byte(1<<rules[2]) == 0 {
				if elf, ok := g.cells.Load(*p); ok {
					elf.dir = d
				}
				if c, ok := g.cells.Load(nc[d]); ok {
					c.proposed++
				} else {
					c = &cell{}
					c.proposed++
					g.cells.Store(nc[d], c)
				}

				return true
			}
			d = (d + 1) % 4
		}
	}

	return false
}
