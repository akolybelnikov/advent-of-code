package advent

import (
	"bufio"
	"bytes"
	"os"
)

const newline = 10

type handlerFunc func(data []byte) (int, error)
type groupHandlerFunction func(group ...[]byte) (int, error)

func ReadDataBytes(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)

	var buf bytes.Buffer
	for scanner.Scan() {
		buf.Write(scanner.Bytes())
	}
	return buf.Bytes(), nil
}

func HandleBytes(data []byte, fn handlerFunc) (int, error) {
	var total, nLine, prevIdx int

	for byteIndex, b := range data {
		if b == newline {
			nLine++
			if nLine == 1 && prevIdx != byteIndex {
				bs := data[prevIdx:byteIndex]
				res, err := fn(bs)
				if err != nil {
					return 0, err
				}
				total += res
			}
			prevIdx = byteIndex + 1
		} else {
			nLine = 0
		}
	}

	return total, nil
}

func handleByteGroups(data []byte, fn groupHandlerFunction, groupSize int) (int, error) {
	var total, nLine, prevIdx, count int
	group := make([][]byte, groupSize, groupSize)

	for byteIndex, b := range data {
		if b == newline {
			nLine++
			if nLine == 1 && prevIdx != byteIndex {
				bs := data[prevIdx:byteIndex]
				group[count] = bs
				count++
				if count == groupSize {
					res, err := fn(group...)
					if err != nil {
						return 0, err
					}
					total += res
					count = 0
				}
			}
			prevIdx = byteIndex + 1
			continue
		} else {
			nLine = 0
		}
	}

	return total, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
