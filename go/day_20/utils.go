package day_20

import utils "github.com/akolybelnikov/advent-of-code"

func encryptedFile(arr *[]*[]byte) []int {
	res := make([]int, 0)
	for _, line := range *arr {
		res = append(res, utils.BytesToInt(line))
	}

	return res
}
