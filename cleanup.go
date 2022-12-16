// Package advent Day 4
package advent

import "strconv"

const (
	comma  = 44
	hyphen = 45
)

func contain(data []byte) (int, error) {
	pairs, err := findPairs(data)
	if err != nil {
		return 0, err
	}
	switch {
	case pairs[0] >= pairs[2] && pairs[1] <= pairs[3]:
		return 1, nil
	case pairs[2] >= pairs[0] && pairs[3] <= pairs[1]:
		return 1, nil
	default:
		return 0, nil
	}
}

func overlap(data []byte) (int, error) {
	pairs, err := findPairs(data)
	if err != nil {
		return 0, err
	}
	switch {
	case pairs[2] >= pairs[0] && pairs[2] <= pairs[1]:
		return 1, nil
	case pairs[0] >= pairs[2] && pairs[0] <= pairs[3]:
		return 1, nil
	default:
		return 0, nil
	}
}

func findPairs(data []byte) ([4]int, error) {
	var curIdx, prevIdx int
	nums := [4]int{}
	for i := 0; i < 4; i++ {
		for data[curIdx] != hyphen && data[curIdx] != comma && curIdx < len(data)-1 {
			curIdx++
		}
		if curIdx == len(data)-1 {
			curIdx++
		}
		num, err := strconv.Atoi(string(data[prevIdx:curIdx]))
		if err != nil {
			return nums, err
		}
		nums[i] = num
		curIdx++
		prevIdx = curIdx
	}

	return nums, nil
}

func FindContainedPairs(data []byte) (int, error) {
	return HandleBytes(data, contain)
}

func FindOverlappingPairs(data []byte) (int, error) {
	return HandleBytes(data, overlap)
}
