// Package advent Day 5
package advent

import (
	"fmt"
	"strconv"
	"strings"
)

const space = 32

type makeMove func(*[]byte, *[][]byte) error

// ParseInput splits the input file into the stacks and moves.
// Stacks are collected as rows and need to be converted to columns.
func ParseInput(dataBytes *[]byte) (int, [][]byte) {
	var numNewlines, valueIndex, movesIndex int
	var lines [][]byte
	line := make([]byte, 0)
	for byteIndex, dataByte := range *dataBytes {
		if dataByte == newline {
			if len(line) > 0 {
				lines = append(lines, line)
			}
			line = make([]byte, 0)
			numNewlines++
			valueIndex = 0
			if numNewlines == 2 {
				movesIndex = byteIndex + 1
				break
			}
		} else {
			if valueIndex%4 == 0 {
				line = append(line, (*dataBytes)[byteIndex+1])
			}
			valueIndex++
			numNewlines = 0
		}
	}

	return movesIndex, lines[:len(lines)-1]
}

// CreateStacks converts stacks rows to columns.
func CreateStacks(lines [][]byte) *[][]byte {
	var stacks [][]byte
	for lineIndex := len(lines) - 1; lineIndex >= 0; lineIndex-- {
		for columnIndex := 0; columnIndex < len(lines[lineIndex]); columnIndex++ {
			if len(stacks) < columnIndex+1 {
				stack := make([]byte, 0)
				stacks = append(stacks, stack)
			}
			if lines[lineIndex][columnIndex] != space {
				stacks[columnIndex] = append(stacks[columnIndex], lines[lineIndex][columnIndex])
			}
		}
	}

	return &stacks
}

func makeStacksAndMoves(data *[]byte) (*[][]byte, *[]byte) {
	movesIdx, stackLines := ParseInput(data)
	stacks := CreateStacks(stackLines)
	moves := (*data)[movesIdx:]

	return stacks, &moves
}

func makeMove9000(data *[]byte, stack *[][]byte) error {
	line := strings.Split(string(*data), " ")
	if len(line) != 6 {
		return fmt.Errorf("wrong line: %v", &line)
	}

	num, origin, dest, err := toIntegers(&line)
	if err != nil {
		return err
	}

	for i := 0; i < num; i++ {
		crate := (*stack)[origin-1][len((*stack)[origin-1])-1]
		(*stack)[origin-1][len((*stack)[origin-1])-1] = 0
		(*stack)[origin-1] = (*stack)[origin-1][:len((*stack)[origin-1])-1]
		(*stack)[dest-1] = append((*stack)[dest-1], crate)
	}

	return nil
}

func makeMove9001(data *[]byte, stack *[][]byte) error {
	line := strings.Split(string(*data), " ")
	if len(line) != 6 {
		return fmt.Errorf("wrong line: %v", &line)
	}

	num, origin, dest, err := toIntegers(&line)
	if err != nil {
		return err
	}

	for i := 0; i < num; i++ {
		crate := (*stack)[origin-1][len((*stack)[origin-1])-num+i]
		(*stack)[origin-1][len((*stack)[origin-1])-num+i] = 0
		(*stack)[dest-1] = append((*stack)[dest-1], crate)
	}
	(*stack)[origin-1] = (*stack)[origin-1][:len((*stack)[origin-1])-num]

	return nil
}

func toIntegers(line *[]string) (int, int, int, error) {
	num, err := strconv.Atoi((*line)[1])
	if err != nil {
		return 0, 0, 0, err
	}
	origin, err := strconv.Atoi((*line)[3])
	if err != nil {
		return 0, 0, 0, err
	}
	dest, err := strconv.Atoi((*line)[5])
	if err != nil {
		return 0, 0, 0, err
	}

	return num, origin, dest, nil
}

func FindTopCrates9000(data *[]byte) (string, error) {
	stacks, moves := makeStacksAndMoves(data)
	err := moveCrates(moves, makeMove9000, stacks)
	if err != nil {
		return "", err
	}

	crates, err := findTopCrates(stacks)
	if err != nil {
		return "", err
	}

	return crates, nil
}

func FindTopCrates9001(data *[]byte) (string, error) {
	stacks, moves := makeStacksAndMoves(data)
	err := moveCrates(moves, makeMove9001, stacks)
	if err != nil {
		return "", err
	}

	crates, err := findTopCrates(stacks)
	if err != nil {
		return "", err
	}

	return crates, nil
}

func findTopCrates(stacks *[][]byte) (string, error) {
	crates := new(strings.Builder)
	for _, stack := range *stacks {
		err := crates.WriteByte(stack[len(stack)-1])
		if err != nil {
			return "", err
		}
	}

	return crates.String(), nil
}

func moveCrates(moves *[]byte, fn makeMove, stacks *[][]byte) error {
	var nLine, prevIdx int
	for byteIndex, b := range *moves {
		if b == newline {
			nLine++
			if nLine == 1 && prevIdx != byteIndex {
				bs := (*moves)[prevIdx:byteIndex]
				err := fn(&bs, stacks)
				if err != nil {
					return err
				}
			}
			prevIdx = byteIndex + 1
			continue
		} else {
			nLine = 0
		}
	}

	return nil
}
