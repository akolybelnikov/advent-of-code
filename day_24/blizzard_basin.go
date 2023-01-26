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

type Point struct {
	X int
	Y int
}

type Queue []*State

type State struct {
	minute int
	x      int
	y      int
}

func FindPath(arr *[]*[]byte) int {
	blizzards := MakeBlizzards(arr)
	width := len(*(*arr)[0]) - 2
	height := len(*arr) - 2
	targetX := width - 1
	targetY := height
	lcm := width * height / utils.Gcd(width, height)
	cache := make(map[State]struct{})
	q := NewQueue()
	q.enqueue(
		State{
			minute: 0,
			x:      0,
			y:      -1,
		})

	directions := [][2]int{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, 0},
	}

	start := Point{0, -1}

	for len(*q) > 0 {
		s := q.dequeue()
		time := s.minute + 1
		for _, d := range directions {
			np := Point{s.x + d[0], s.y + d[1]}

			if np.Y == targetY && np.X == targetX {
				return time
			}

			if (np.Y >= 0 && np.Y < height && np.X >= 0 && np.X < width && np != start) || np == start {
				collides := false
				if np != start {
					for i, dr := range directions[:4] {
						nx := ((np.X-dr[0]*time)%width + width) % width
						ny := ((np.Y-dr[1]*time)%height + height) % height
						if blizzard, ok := (*blizzards)[Point{nx, ny}]; ok {
							if blizzard == uint8(i) {
								collides = true
								break
							}
						}
					}
				}

				if !collides {
					cacheState := State{minute: time % lcm, x: np.X, y: np.Y}
					if _, ok := cache[cacheState]; ok {
						continue
					}
					cache[cacheState] = struct{}{}
					queueState := State{minute: time, x: np.X, y: np.Y}
					q.enqueue(queueState)
				}
			}
		}

	}

	return 0
}

func (q *Queue) enqueue(s State) {
	*q = append(*q, &s)
}

func (q *Queue) dequeue() *State {
	if len(*q) == 0 {
		return nil
	}
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}

func NewQueue() *Queue {
	q := make(Queue, 0)
	return &q
}

func MakeBlizzards(arr *[]*[]byte) *map[Point]uint8 {
	var ey, ex int
	res := make(map[Point]uint8)
	for i, v := range *arr {
		for i2, bt := range *v {
			if bt != '#' {
				ex, ey = i2-1, i-1
				p := Point{X: ex, Y: ey}
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
