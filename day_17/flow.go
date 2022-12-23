package day_17

const (
	LEFT byte = 60
)

type group map[int]int

type state struct {
	shape  map[int]struct{}
	buffer group
	cache  map[int]struct{}
	left   int
	right  int
	top    int
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

func initState() *state {
	return &state{
		buffer: map[int]int{
			0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0,
		},
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

func (g *group) overlap(other *group) bool {
	for k, v := range *g {
		if (*other)[k] == v {
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
