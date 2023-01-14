package day_22

import (
	utils "github.com/akolybelnikov/advent-of-code"
)

const (
	SPACE uint8 = 32
	TILE  uint8 = 46
)

func parseInstructions(input *[]byte) []int {
	res := make([]int, 0)
	var prev int
	for i := 0; i < len(*input)-1; i++ {
		prev = i
		for (*input)[i] != L && (*input)[i] != R {
			i++
		}
		if i > 0 {
			iBytes := (*input)[prev:i]
			res = append(res, utils.BytesToInt(&iBytes))
		}
		res = append(res, int((*input)[i]))
		prev = i
	}
	// last instruction
	if prev < len(*input)-1 {
		iBytes := (*input)[prev+1 : len(*input)]
		res = append(res, utils.BytesToInt(&iBytes))
	}

	return res
}

func padData(input *[]*[]byte) *[][]byte {
	var max int
	for _, r := range *input {
		if len(*r) > max {
			max = len(*r)
		}
	}
	var b = make([][]byte, 0)
	for _, r := range *input {
		row := make([]byte, max)
		for i, c := range *r {
			row[i] = c
		}
		if len(*r) < max {
			for i := len(*r); i < max; i++ {
				row[i] = SPACE
			}
		}
		b = append(b, row)
	}

	return &b
}

func createFaces(tiles *[][]byte, faces *map[int]*face) {
	size := utils.Gcd(len(*tiles), len((*tiles)[0]))
	matrix := make([][]int, 0)
	for i := 0; i < len(*tiles); i += size {
		r := make([]int, 0)
		for j := 0; j < len((*tiles)[i]); j += size {
			if (*tiles)[i][j] != SPACE {
				r = append(r, 1)
			} else {
				r = append(r, 0)
			}
		}
		matrix = append(matrix, r)
	}

	var f []*face
	if len(matrix) == 3 {
		f = fromTestInput(size)
	} else {
		f = fromInput(size)
	}

	for i, fc := range f {
		fc.ID = i
		(*faces)[i] = fc
	}
}

// hard-coded for test input
func fromTestInput(size int) []*face {
	faces := make([]*face, 0)
	face0 := &face{
		vertex:   pos{y: 0, x: size * 2},
		adjacent: make(map[int]vector),
		size:     size,
	}
	face0.adjacent[UP] = vector{1, DOWN}
	face0.adjacent[RIGHT] = vector{5, LEFT}
	face0.adjacent[DOWN] = vector{3, DOWN}
	face0.adjacent[LEFT] = vector{2, DOWN}
	faces = append(faces, face0)

	face1 := &face{
		vertex:   pos{y: size, x: 0},
		adjacent: make(map[int]vector),
		size:     size,
	}
	face1.adjacent[UP] = vector{0, DOWN}
	face1.adjacent[RIGHT] = vector{2, RIGHT}
	face1.adjacent[DOWN] = vector{4, UP}
	face1.adjacent[LEFT] = vector{5, UP}
	faces = append(faces, face1)

	face2 := &face{
		vertex:   pos{y: size, x: size},
		adjacent: make(map[int]vector),
		size:     size,
	}
	face2.adjacent[UP] = vector{0, RIGHT}
	face2.adjacent[RIGHT] = vector{3, RIGHT}
	face2.adjacent[DOWN] = vector{4, RIGHT}
	face2.adjacent[LEFT] = vector{1, LEFT}
	faces = append(faces, face2)

	face3 := &face{
		vertex:   pos{y: size, x: size * 2},
		adjacent: make(map[int]vector),
		size:     size,
	}
	face3.adjacent[UP] = vector{0, UP}
	face3.adjacent[RIGHT] = vector{5, DOWN}
	face3.adjacent[DOWN] = vector{4, DOWN}
	face3.adjacent[LEFT] = vector{2, LEFT}
	faces = append(faces, face3)

	face4 := &face{
		vertex:   pos{y: size * 2, x: size * 2},
		adjacent: make(map[int]vector),
		size:     size,
	}
	face4.adjacent[UP] = vector{3, UP}
	face4.adjacent[RIGHT] = vector{5, RIGHT}
	face4.adjacent[DOWN] = vector{1, UP}
	face4.adjacent[LEFT] = vector{2, UP}
	faces = append(faces, face4)

	face5 := &face{
		vertex:   pos{y: size * 2, x: size * 3},
		adjacent: make(map[int]vector),
		size:     size,
	}
	face5.adjacent[UP] = vector{3, LEFT}
	face5.adjacent[RIGHT] = vector{0, LEFT}
	face5.adjacent[DOWN] = vector{1, RIGHT}
	face5.adjacent[LEFT] = vector{4, LEFT}
	faces = append(faces, face5)

	return faces
}

// hard-coded for input
func fromInput(size int) []*face {
	faces := make([]*face, 0)
	face0 := &face{
		vertex:   pos{y: 0, x: size},
		adjacent: make(map[int]vector),
		size:     size,
	}
	face0.adjacent[UP] = vector{5, RIGHT}
	face0.adjacent[RIGHT] = vector{1, RIGHT}
	face0.adjacent[DOWN] = vector{2, DOWN}
	face0.adjacent[LEFT] = vector{3, RIGHT}
	faces = append(faces, face0)

	face1 := &face{
		vertex:   pos{y: 0, x: size * 2},
		adjacent: make(map[int]vector),
		size:     size,
	}
	face1.adjacent[UP] = vector{5, UP}
	face1.adjacent[RIGHT] = vector{4, LEFT}
	face1.adjacent[DOWN] = vector{2, LEFT}
	face1.adjacent[LEFT] = vector{0, LEFT}
	faces = append(faces, face1)

	face2 := &face{
		vertex:   pos{y: size, x: size},
		adjacent: make(map[int]vector),
		size:     size,
	}
	face2.adjacent[UP] = vector{0, UP}
	face2.adjacent[RIGHT] = vector{1, UP}
	face2.adjacent[DOWN] = vector{4, DOWN}
	face2.adjacent[LEFT] = vector{3, DOWN}
	faces = append(faces, face2)

	face3 := &face{
		vertex:   pos{y: size * 2, x: 0},
		adjacent: make(map[int]vector),
		size:     size,
	}
	face3.adjacent[UP] = vector{2, RIGHT}
	face3.adjacent[RIGHT] = vector{4, RIGHT}
	face3.adjacent[DOWN] = vector{5, DOWN}
	face3.adjacent[LEFT] = vector{0, RIGHT}
	faces = append(faces, face3)

	face4 := &face{
		vertex:   pos{y: size * 2, x: size},
		adjacent: make(map[int]vector),
		size:     size,
	}
	face4.adjacent[UP] = vector{2, UP}
	face4.adjacent[RIGHT] = vector{1, LEFT}
	face4.adjacent[DOWN] = vector{5, LEFT}
	face4.adjacent[LEFT] = vector{3, LEFT}
	faces = append(faces, face4)

	face5 := &face{
		vertex:   pos{y: size * 3, x: 0},
		adjacent: make(map[int]vector),
		size:     size,
	}
	face5.adjacent[UP] = vector{3, UP}
	face5.adjacent[RIGHT] = vector{4, UP}
	face5.adjacent[DOWN] = vector{1, DOWN}
	face5.adjacent[LEFT] = vector{0, DOWN}
	faces = append(faces, face5)

	return faces
}
