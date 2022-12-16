// Package advent Day 3
package advent

import (
	"fmt"
	a "github.com/akolybelnikov/advent-of-code"
)

var priorities = map[int]int{
	97:  1,
	98:  2,
	99:  3,
	100: 4,
	101: 5,
	102: 6,
	103: 7,
	104: 8,
	105: 9,
	106: 10,
	107: 11,
	108: 12,
	109: 13,
	110: 14,
	111: 15,
	112: 16,
	113: 17,
	114: 18,
	115: 19,
	116: 20,
	117: 21,
	118: 22,
	119: 23,
	120: 24,
	121: 25,
	122: 26,
	65:  27,
	66:  28,
	67:  29,
	68:  30,
	69:  31,
	70:  32,
	71:  33,
	72:  34,
	73:  35,
	74:  36,
	75:  37,
	76:  38,
	77:  39,
	78:  40,
	79:  41,
	80:  42,
	81:  43,
	82:  44,
	83:  45,
	84:  46,
	85:  47,
	86:  48,
	87:  49,
	88:  50,
	89:  51,
	90:  52,
}

func findDuplicatePriority(data []byte) (int, error) {
	var item int

	left, right := data[:len(data)/2], data[len(data)/2:]

OUTER:
	for _, b := range left {
		for _, c := range right {
			if b == c {
				item = int(b)
				break OUTER
			}
		}
	}

	if priority, ok := priorities[item]; ok {
		return priority, nil
	} else {
		return 0, fmt.Errorf("item %c is not listed", rune(item))
	}
}

func findGroupBadge(group ...[]byte) (int, error) {
	badge, err := intersection(group...)
	if err != nil {
		fmt.Printf("encountered an error: %v\n", err)
		return 0, err
	}
	if priority, ok := priorities[int(badge)]; ok {
		return priority, nil
	} else {
		return 0, fmt.Errorf("item %c is not listed", rune(badge))
	}
}
func intersection(pS ...[]byte) (byte, error) {
	hashes := make([]map[byte]int, len(pS))
	for _, slice := range pS {
		hash := make(map[byte]int)
		for _, value := range slice {
			hash[value]++
		}
		hashes = append(hashes, hash)
	}

	resultHash := make(map[byte]int)
	for _, hash := range hashes {
		for k, _ := range hash {
			resultHash[k]++
		}
	}

	result := make([]byte, 0)
	for value, count := range resultHash {
		if count == len(pS) {
			result = append(result, value)
		}
	}

	if len(result) == 0 || len(result) > 1 {
		return 0, fmt.Errorf("no badge or too many %v", result)
	}

	return result[0], nil
}

func FindTotalPriorities(data []byte) (int, error) {
	return a.HandleBytes(data, findDuplicatePriority)
}

func FindTotalBadges(data []byte) (int, error) {
	return a.HandleByteGroups(data, findGroupBadge, 3)
}
