package day_22

const (
	RIGHT = iota
	DOWN
	LEFT
	UP
)

const (
	L = 76
	R = 82
)

type pos struct {
	rowIdx, columnIdx, facing int
}

type board [][]byte

func MonkeyMap(data *[]*[]byte) int {
	instrBytes := (*data)[(len(*data) - 1):][0]
	instructions := parseInstructions(instrBytes)

	gridData := (*data)[:len(*data)-1]
	b := padData(&gridData)

	s := b.start()

	for _, i := range instructions {
		switch i {
		case L:
			s.facing = (s.facing + 3) % 4
		case R:
			s.facing = (s.facing + 1) % 4
		default:
			b.move(s, i)
		}
	}

	return s.password()
}

func (b *board) start() *pos {
	p := &pos{rowIdx: 0, facing: RIGHT}
	for i, c := range (*b)[0] {
		if c == TILE {
			p.columnIdx = i
			break
		}
	}
	return p
}

func (b *board) move(p *pos, instr int) {
	for n := 0; n < instr; n++ {
		next := b.lookAhead(p)
		if (*b)[next.rowIdx][next.columnIdx] == TILE {
			p.rowIdx = next.rowIdx
			p.columnIdx = next.columnIdx
		}
	}
}

func (b *board) lookAhead(p *pos) *pos {
	np := pos{rowIdx: p.rowIdx, columnIdx: p.columnIdx, facing: p.facing}
	np.update()
	np.rowIdx = (np.rowIdx + len(*b)) % len(*b)
	np.columnIdx = (np.columnIdx + len((*b)[0])) % len((*b)[0])

	for (*b)[np.rowIdx][np.columnIdx] == SPACE {
		np.update()
		np.rowIdx = (np.rowIdx + len(*b)) % len(*b)
		np.columnIdx = (np.columnIdx + len((*b)[0])) % len((*b)[0])
	}

	return &np
}

func (p *pos) update() {
	switch p.facing {
	case RIGHT:
		p.columnIdx++
	case DOWN:
		p.rowIdx++
	case LEFT:
		p.columnIdx--
	case UP:
		p.rowIdx--
	}
}

func (p *pos) password() int {
	return ((p.rowIdx + 1) * 1000) + ((p.columnIdx + 1) * 4) + p.facing
}
