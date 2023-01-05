package day_21

import utils "github.com/akolybelnikov/advent-of-code"

func parseMonkeyData(data *[]*[]byte, g *graph) {
	for _, line := range *data {
		n := name{(*line)[0], (*line)[1], (*line)[2], (*line)[3]}
		var mk *monkey
		// find out if monkey has been already created
		if m, ok := (*g)[n]; !ok {
			mk = &monkey{
				name: n,
				deps: make([]*chan int, 0),
			}
			(*g)[n] = mk
		} else {
			mk = m
		}
		// parse and assign value or operation
		if (*line)[6] > 47 && (*line)[6] < 58 {
			var b = make([]byte, 0)
			for _, c := range (*line)[6:len(*line)] {
				b = append(b, c)
			}
			mk.val = utils.BytesToInt(&b)
		} else {
			var lm, rm *monkey
			mk.left = make(chan int, 1)
			mk.right = make(chan int, 1)
			mk.leftVal = make([]int, 0)
			mk.rightVal = make([]int, 0)
			mk.waiting = false
			ln := name{(*line)[6], (*line)[7], (*line)[8], (*line)[9]}
			if mkl, ok := (*g)[ln]; !ok {
				lm = &monkey{
					name: ln,
					deps: []*chan int{&mk.left},
				}
				(*g)[ln] = lm
			} else {
				lm = mkl
				lm.deps = append(lm.deps, &mk.left)
			}

			mk.op = (*line)[11]

			rn := name{(*line)[13], (*line)[14], (*line)[15], (*line)[16]}
			if mkr, ok := (*g)[rn]; !ok {
				rm = &monkey{
					name: rn,
					deps: []*chan int{&mk.right},
				}
				(*g)[rn] = rm
			} else {
				rm = mkr
				rm.deps = append(rm.deps, &mk.right)
			}
		}
	}
}
