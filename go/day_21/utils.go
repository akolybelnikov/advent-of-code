package day_21

import utils "github.com/akolybelnikov/advent-of-code"

func parseMonkeyData(data *[]*[]byte, g *graph) {
	humn := name{'h', 'u', 'm', 'n'}
	for _, line := range *data {
		n := name{(*line)[0], (*line)[1], (*line)[2], (*line)[3]}
		var mk *monkey
		// find out if monkey has been already created
		if m, ok := (*g)[n]; !ok {
			mk = &monkey{
				name:      n,
				leftDeps:  make([]*monkey, 0),
				rightDeps: make([]*monkey, 0),
			}
			(*g)[n] = mk
		} else {
			mk = m
		}

		if n == humn {
			mk.depOnHuman = true
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

			ln := name{(*line)[6], (*line)[7], (*line)[8], (*line)[9]}
			mk.leftName = ln
			if mkl, ok := (*g)[ln]; !ok {
				lm = &monkey{
					name:      ln,
					leftDeps:  []*monkey{mk},
					rightDeps: make([]*monkey, 0),
				}
				(*g)[ln] = lm
			} else {
				lm = mkl
				lm.leftDeps = append(lm.leftDeps, mk)
			}

			mk.op = (*line)[11]

			rn := name{(*line)[13], (*line)[14], (*line)[15], (*line)[16]}
			mk.rightName = rn
			if mkr, ok := (*g)[rn]; !ok {
				rm = &monkey{
					name:      rn,
					leftDeps:  make([]*monkey, 0),
					rightDeps: []*monkey{mk},
				}
				(*g)[rn] = rm
			} else {
				rm = mkr
				rm.rightDeps = append(rm.rightDeps, mk)
			}

			go mk.wait()
		}
	}
}
