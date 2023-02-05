package day_17

import "reflect"

const (
	LEFT byte = 60
)

type state struct {
	shape map[int]struct{}
	cache map[int]struct{}
	left  int
	right int
	top   int
}

func Run(pushes *[]byte, rocks int) int {
	var count int
	var next bool
	s := initState()
	shapeIdx := shape()
	jet := push(pushes)

	for count < rocks {
		next = s.spawn(shapeIdx())
		for next {
			s.shift(jet())
			next = s.sink()
		}
		count++
	}

	return s.top
}

func Run2(pushes *[]byte, rocks int) int {
	prefix, cycle, sumStart, sumCycle, _, cycleValues := findCycle(pushes)
	quotient := (rocks - prefix) / cycle
	remainder := (rocks - prefix) % cycle
	maxY := sumStart + quotient*sumCycle
	for i := 0; i < remainder; i++ {
		maxY += cycleValues[i]
	}

	return maxY
}

func findCycle(input *[]byte) (int, int, int, int, []int, []int) {
	values := make([]int, 0)
	curTop := 0
	var next bool
	s := initState()
	shapeIdx := shape()
	jet := push(input)
	for i := 0; i < 3*len(*input); i++ {
		next = s.spawn(shapeIdx())
		for next {
			s.shift(jet())
			next = s.sink()
		}
		prevTop := curTop
		curTop = s.top
		values = append(values, curTop-prevTop)
	}
	i, j := findRecurringElement(prefixes(values, len(*input)))
	sumStart := 0
	sumCycle := 0
	for k := 0; k < i; k++ {
		sumStart += values[k]
	}
	for k := i; k < j; k++ {
		sumCycle += values[k]
	}
	return i, j - i, sumStart, sumCycle, values[:i], values[i:j]
}

func findRecurringElement(l [][]int) (int, int) {
	for i := 0; i < len(l); i++ {
		for j := i + 1; j < len(l); j++ {
			if reflect.DeepEqual(l[i], l[j]) {
				return i, j
			}
		}
	}
	return -1, -1
}

func prefixes(l []int, n int) [][]int {
	var res [][]int
	for i := 0; i < len(l)-n; i++ {
		res = append(res, l[i:i+n])
	}
	return res
}

func initState() *state {
	return &state{
		cache: make(map[int]struct{}),
	}
}

func (s *state) spawn(idx int) bool {
	ns := make(map[int]struct{})
	row := s.top * 10
	switch idx {
	case 0:
		ns[row+43] = struct{}{}
		ns[row+44] = struct{}{}
		ns[row+45] = struct{}{}
		ns[row+46] = struct{}{}
		s.right = 5
	case 1:
		ns[row+53] = struct{}{}
		ns[row+44] = struct{}{}
		ns[row+54] = struct{}{}
		ns[row+64] = struct{}{}
		ns[row+55] = struct{}{}
		s.right = 4
	case 2:
		ns[row+43] = struct{}{}
		ns[row+44] = struct{}{}
		ns[row+45] = struct{}{}
		ns[row+55] = struct{}{}
		ns[row+65] = struct{}{}
		s.right = 4
	case 3:
		ns[row+43] = struct{}{}
		ns[row+53] = struct{}{}
		ns[row+63] = struct{}{}
		ns[row+73] = struct{}{}
		s.right = 2
	case 4:
		ns[row+43] = struct{}{}
		ns[row+53] = struct{}{}
		ns[row+44] = struct{}{}
		ns[row+54] = struct{}{}
		s.right = 3
	default:
		return false

	}
	s.left = 2
	s.shape = ns

	return true
}

func (s *state) sink() bool {
	lookahead := make(map[int]struct{})
	for k, _ := range s.shape {
		lookahead[k-10] = struct{}{}
	}
	if s.collide(&lookahead) {
		s.cacheShape()
		return false
	} else {
		s.shape = lookahead
	}

	return true
}

func (s *state) cacheShape() {
	for k, _ := range s.shape {
		s.cache[k] = struct{}{}
		if k/10 > s.top {
			s.top = k / 10
		}
	}
}

func (s *state) shift(jet *byte) {
	if *jet == LEFT {
		s.updateLeft()
	} else {
		s.updateRight()
	}
}

func (s *state) updateLeft() {
	if s.left > 0 {
		lookahead := make(map[int]struct{})
		for k, v := range s.shape {
			lookahead[k-1] = v
		}
		if !s.collide(&lookahead) {
			s.shape = lookahead
			s.left--
			s.right--
		}
	}
}

func (s *state) updateRight() {
	if s.right < 6 {
		lookahead := make(map[int]struct{})
		for k, v := range s.shape {
			lookahead[k+1] = v
		}
		if !s.collide(&lookahead) {
			s.shape = lookahead
			s.left++
			s.right++
		}
	}
}

func (s *state) collide(lookahead *map[int]struct{}) bool {
	for k, _ := range *lookahead {
		if _, ok := s.cache[k]; ok || k < 10 {
			return true
		}
	}
	return false
}

func push(pushes *[]byte) func() *byte {
	cur := 0
	return func() *byte {
		p := (*pushes)[cur]
		cur++
		if cur > len(*pushes)-1 {
			cur = 0
		}

		return &p
	}
}

func shape() func() int {
	cur := 0
	return func() int {
		idx := cur
		cur++
		if cur > 4 {
			cur = 0
		}
		return idx
	}
}
