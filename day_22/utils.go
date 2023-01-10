package day_22

import utils "github.com/akolybelnikov/advent-of-code"

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

func padData(input *[]*[]byte) *board {
	var max int
	for _, r := range *input {
		if len(*r) > max {
			max = len(*r)
		}
	}
	var b board
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
