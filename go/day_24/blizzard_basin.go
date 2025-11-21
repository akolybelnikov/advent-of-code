package day_24

import (
	utils "github.com/akolybelnikov/advent-of-code"
)

const (
	N = iota
	E
	S
	W
)

type point struct {
	x int
	y int
}

type queue []*state

type state struct {
	minute int
	point  point
	stage  int
}

type basin struct {
	blizzards          *map[point]uint8
	cache              *map[state]struct{}
	directions         [][2]int
	height, width, lcm int
}

func FindPath(arr *[]*[]byte) int {
	b := newBasin(arr)

	start := point{0, -1}
	target := point{b.width - 1, b.height}
	q := newQueue()
	q.enqueue(
		state{
			minute: 0,
			point:  start,
			stage:  0,
		})

	return q.run(start, target, b)
}

func FindPath2(arr *[]*[]byte) int {
	b := newBasin(arr)

	gates := []point{{b.width - 1, b.height}, {0, -1}}
	q := newQueue()
	q.enqueue(state{
		minute: 0,
		point:  gates[1],
		stage:  0,
	})

	return q.run2(gates, b)
}

func newBasin(arr *[]*[]byte) *basin {
	cache := make(map[state]struct{})
	b := basin{
		blizzards: MakeBlizzards(arr),
		cache:     &cache,
		width:     len(*(*arr)[0]) - 2,
		height:    len(*arr) - 2,
		directions: [][2]int{
			{0, -1},
			{1, 0},
			{0, 1},
			{-1, 0},
			{0, 0},
		},
	}
	b.lcm = b.width * b.height / utils.Gcd(b.width, b.height)
	return &b
}

func newQueue() *queue {
	q := make(queue, 0)
	return &q
}

func (q *queue) enqueue(s state) {
	*q = append(*q, &s)
}

func (q *queue) dequeue() *state {
	if len(*q) == 0 {
		return nil
	}
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}

func (q *queue) run(start, target point, basin *basin) int {
	for len(*q) > 0 {
		s := q.dequeue()
		time := s.minute + 1
		for _, d := range basin.directions {
			np := point{s.point.x + d[0], s.point.y + d[1]}

			if np == target {
				return time
			}

			if (basin.containsPoint(&np) && np != start) || np == start {
				collides := false
				if np != start {
					collides = basin.collides(&np, time)
				}

				if !collides {
					cacheState := state{minute: time % basin.lcm, point: np}
					if _, ok := (*basin.cache)[cacheState]; ok {
						continue
					}
					(*basin.cache)[cacheState] = struct{}{}
					queueState := state{minute: time, point: np}
					q.enqueue(queueState)
				}
			}
		}

	}
	return 0
}

func (q *queue) run2(gates []point, basin *basin) int {
	for len(*q) > 0 {
		s := q.dequeue()
		time := s.minute + 1
		gate := gates[s.stage%2]
		for _, d := range basin.directions {
			np := point{s.point.x + d[0], s.point.y + d[1]}
			nextStage := s.stage

			if np == gate {
				if s.stage == 2 {
					return time
				}
				nextStage++
			}

			if (basin.containsPoint(&np) && np != gates[0] && np != gates[1]) ||
				np == gates[0] || np == gates[1] {
				collides := false
				if np != gates[0] && np != gates[1] {
					collides = basin.collides(&np, time)
				}

				if !collides {
					cacheState := state{minute: time % basin.lcm, point: np, stage: nextStage}
					if _, ok := (*basin.cache)[cacheState]; ok {
						continue
					}
					(*basin.cache)[cacheState] = struct{}{}
					queueState := state{minute: time, point: np, stage: nextStage}
					q.enqueue(queueState)
				}
			}
		}

	}

	return 0
}

func (b *basin) containsPoint(p *point) bool {
	return p.y >= 0 && p.y < b.height && p.x >= 0 && p.x < b.width
}

func (b *basin) collides(p *point, time int) bool {
	for i, dr := range b.directions[:4] {
		nx := ((p.x-dr[0]*time)%b.width + b.width) % b.width
		ny := ((p.y-dr[1]*time)%b.height + b.height) % b.height
		if blizzard, ok := (*b.blizzards)[point{nx, ny}]; ok {
			if blizzard == uint8(i) {
				return true
			}
		}
	}

	return false
}

func MakeBlizzards(arr *[]*[]byte) *map[point]uint8 {
	var ey, ex int
	res := make(map[point]uint8)
	for i, v := range *arr {
		for i2, bt := range *v {
			if bt != '#' {
				ex, ey = i2-1, i-1
				p := point{x: ex, y: ey}
				if bt != '.' {
					switch bt {
					case '^':
						res[p] = N
					case '>':
						res[p] = E
					case 'v':
						res[p] = S
					case '<':
						res[p] = W
					}
				}
			}
		}
	}

	return &res
}
