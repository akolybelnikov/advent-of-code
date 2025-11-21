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

type wrap int

const (
	wrap2D wrap = iota
	wrap3D
)

type pos struct {
	y, x, facing int
}

type board struct {
	tiles [][]byte
	faces map[int]*face
	cur   int
}

type face struct {
	ID       int
	size     int
	vertex   pos
	adjacent map[int]vector
}

type vector [2]int

func MonkeyMap(data *[]*[]byte) int {
	instrBytes := (*data)[(len(*data) - 1):][0]
	instructions := parseInstructions(instrBytes)

	gridData := (*data)[:len(*data)-1]
	b := board{
		tiles: *padData(&gridData),
	}

	pswd := b.walk(&instructions, wrap2D)

	return pswd
}

func MonkeyMap2(data *[]*[]byte) int {
	instrBytes := (*data)[(len(*data) - 1):][0]
	instructions := parseInstructions(instrBytes)

	gridData := (*data)[:len(*data)-1]
	tiles := padData(&gridData)
	b := board{
		tiles: *tiles,
		cur:   0,
		faces: make(map[int]*face),
	}

	createFaces(tiles, &b.faces)

	pswd := b.walk(&instructions, wrap3D)

	return pswd
}

func (b *board) walk(instructions *[]int, w wrap) int {
	s := b.start()

	for _, i := range *instructions {
		switch i {
		case L:
			s.facing = (s.facing + 3) % 4
		case R:
			s.facing = (s.facing + 1) % 4
		default:
			b.move(s, i, w)
		}
	}

	return s.password()
}

func (b *board) start() *pos {
	p := &pos{y: 0, facing: RIGHT}
	for i, c := range b.tiles[0] {
		if c == TILE {
			p.x = i
			break
		}
	}
	return p
}

func (b *board) move(p *pos, instr int, w wrap) {
	for n := 0; n < instr; n++ {
		var nextPos *pos
		var nextFace int
		if w == wrap2D {
			nextPos = b.lookAhead2D(p)
		} else {
			nextPos, nextFace = b.lookAhead3D(p)
		}

		if b.tiles[nextPos.y][nextPos.x] == TILE {
			p.y = nextPos.y
			p.x = nextPos.x
			p.facing = nextPos.facing
			b.cur = nextFace
		} else {
			break
		}
	}
}

func (b *board) lookAhead2D(p *pos) *pos {
	np := pos{y: p.y, x: p.x, facing: p.facing}
	np.update()
	np.y = (np.y + len(b.tiles)) % len(b.tiles)
	np.x = (np.x + len((b.tiles)[0])) % len((b.tiles)[0])

	for (b.tiles)[np.y][np.x] == SPACE {
		np.update()
		np.y = (np.y + len(b.tiles)) % len(b.tiles)
		np.x = (np.x + len((b.tiles)[0])) % len((b.tiles)[0])
	}

	return &np
}

func (b *board) lookAhead3D(p *pos) (*pos, int) {
	np := pos{y: p.y, x: p.x, facing: p.facing}
	np.update()

	if !b.faces[b.cur].within(&np) {
		return b.swap(&np)
	}

	return &np, b.cur
}

func (b *board) swap(p *pos) (*pos, int) {
	adj := b.faces[b.cur].adjacent[p.facing]
	f := b.faces[adj[0]]
	sp := f.makeSwap(p, adj[1])

	return sp, f.ID
}

func (p *pos) update() {
	switch p.facing {
	case RIGHT:
		p.x++
	case DOWN:
		p.y++
	case LEFT:
		p.x--
	case UP:
		p.y--
	}
}

func (p *pos) password() int {
	return ((p.y + 1) * 1000) + ((p.x + 1) * 4) + p.facing
}

func (f *face) within(np *pos) bool {
	return np.y >= f.vertex.y && np.x >= f.vertex.x &&
		np.y <= f.vertex.y+(f.size-1) && np.x <= f.vertex.x+(f.size-1)
}

func (f *face) makeSwap(p *pos, newDir int) *pos {
	np := &pos{y: p.y, x: p.x, facing: p.facing}

	switch {
	case p.facing == UP && newDir == DOWN:
		np.y = f.vertex.y
		np.x = f.vertex.x + f.size - np.x%f.size - 1
		np.facing = DOWN
	case p.facing == DOWN && newDir == UP:
		np.y = f.vertex.y + f.size - 1
		np.x = f.vertex.x + f.size - np.x%f.size - 1
		np.facing = UP
	case p.facing == LEFT && newDir == RIGHT:
		np.y = f.vertex.y + f.size - np.y%f.size - 1
		np.x = f.vertex.x
		np.facing = RIGHT
	case p.facing == RIGHT && newDir == LEFT:
		np.y = f.vertex.y + f.size - np.y%f.size - 1
		np.x = f.vertex.x + f.size - 1
		np.facing = LEFT
	case p.facing == LEFT && newDir == DOWN:
		np.y = f.vertex.y
		np.x = f.vertex.x + p.y%f.size
		np.facing = DOWN
	case p.facing == DOWN && newDir == LEFT:
		np.y = f.vertex.y + p.x%f.size
		np.x = f.vertex.x + f.size - 1
		np.facing = LEFT
	case p.facing == LEFT && newDir == UP:
		np.y = f.vertex.y
		np.x = f.vertex.x + p.y%f.size
		np.facing = UP
	case p.facing == UP && newDir == LEFT:
		np.y = f.vertex.y + f.size - p.x%f.size - 1
		np.x = f.vertex.x + f.size - 1
		np.facing = LEFT
	case p.facing == RIGHT && newDir == DOWN:
		np.y = f.vertex.y
		np.x = f.vertex.x + f.size - p.y%f.size - 1
		np.facing = DOWN
	case p.facing == DOWN && newDir == RIGHT:
		np.y = f.vertex.y + f.size - p.x%f.size - 1
		np.x = f.vertex.x
		np.facing = RIGHT
	case p.facing == RIGHT && newDir == UP:
		np.y = f.vertex.y + f.size - 1
		np.x = f.vertex.x + p.y%f.size
		np.facing = UP
	case p.facing == UP && newDir == RIGHT:
		np.y = f.vertex.y + p.x%f.size
		np.x = f.vertex.x
		np.facing = RIGHT
	case p.facing == UP && newDir == UP:
		np.y = f.vertex.y + f.size - 1
		np.x = f.vertex.x + p.x%f.size
		np.facing = UP
	case p.facing == DOWN && newDir == DOWN:
		np.y = f.vertex.y
		np.x = f.vertex.x + p.x%f.size
		np.facing = DOWN
	case p.facing == LEFT && newDir == LEFT:
		np.y = f.vertex.y + p.y%f.size
		np.x = f.vertex.x + f.size - 1
		np.facing = LEFT
	case p.facing == RIGHT && newDir == RIGHT:
		np.y = f.vertex.y + p.y%f.size
		np.x = f.vertex.x
		np.facing = RIGHT
	}

	return np
}
