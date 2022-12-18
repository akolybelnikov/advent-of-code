package advent

import "strconv"

type direction byte
type position struct {
	x, y int
}

type state struct {
	knots   []position
	visited *map[position]int
}

type instruction struct {
	direction *direction
	steps     int
}

const (
	down  direction = 68
	left  direction = 76
	right direction = 82
	up    direction = 85
)

func initState(knots int) *state {
	return &state{
		knots:   make([]position, knots),
		visited: &map[position]int{position{0, 0}: 1},
	}
}

func parseInstruction(data *[]byte) (*instruction, error) {
	d := direction((*data)[0])
	steps, err := strconv.Atoi(string((*data)[2:]))
	if err != nil {
		return nil, err
	}

	return &instruction{
		direction: &d,
		steps:     steps,
	}, nil
}

func (s *state) move(d *direction) {
	switch *d {
	case up:
		s.knots[0].x++
	case down:
		s.knots[0].x--
	case left:
		s.knots[0].y--
	case right:
		s.knots[0].y++
	}
}

func (s *state) moveTail() {
	for i := 1; i < len(s.knots); i++ {
		delta := &position{s.knots[i-1].x - s.knots[i].x, s.knots[i-1].y - s.knots[i].y}
		if Abs(delta.x) <= 1 && Abs(delta.y) <= 1 {
			return
		}
		if delta.y > 0 {
			s.knots[i].y++
		} else if delta.y < 0 {
			s.knots[i].y--
		}
		if delta.x > 0 {
			s.knots[i].x++
		} else if delta.x < 0 {
			s.knots[i].x--
		}
	}
	(*s.visited)[s.knots[len(s.knots)-1]]++
}

func VisitedPositions(data *[]byte, numKnots int) (int, error) {
	var line, prevIdx int
	s := initState(numKnots)

	for byteIndex, b := range *data {
		if b == NEWLINE {
			line++
			if line == 1 && prevIdx != byteIndex {
				nextStep := (*data)[prevIdx:byteIndex]
				i, err := parseInstruction(&nextStep)
				if err != nil {
					return 0, err
				}
				for stepIdx := 0; stepIdx < i.steps; stepIdx++ {
					s.move(i.direction)
					s.moveTail()
				}
			}
			prevIdx = byteIndex + 1
		} else {
			line = 0
		}
	}

	return len(*s.visited), nil
}
