package advent

import (
	"bufio"
	"bytes"
	"golang.org/x/exp/constraints"
	"math"
	"os"
	"sync"
)

const NEWLINE = 10

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

func MakeBytesArray(data *[]byte) (*[]*[]byte, error) {
	var nLine, prevIdx int
	res := make([]*[]byte, 0)

	for byteIndex, b := range *data {
		if b == NEWLINE {
			nLine++
			if nLine == 1 && prevIdx != byteIndex {
				bs := (*data)[prevIdx:byteIndex]
				res = append(res, &bs)
			}
			prevIdx = byteIndex + 1
		} else {
			nLine = 0
		}
	}

	return &res, nil
}

func HandleBytes(data []byte, fn handlerFunc) (int, error) {
	var total, nLine, prevIdx int

	for byteIndex, b := range data {
		if b == NEWLINE {
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

func HandleByteGroups(data []byte, fn groupHandlerFunction, groupSize int) (int, error) {
	var total, nLine, prevIdx, count int
	group := make([][]byte, groupSize, groupSize)

	for byteIndex, b := range data {
		if b == NEWLINE {
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

func Abs[T constraints.Float | constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func BytesToInt(bytes *[]byte) int {
	var val int
	neg := (*bytes)[0] == 45
	if neg {
		*bytes = (*bytes)[1:]
	}
	for _, b := range *bytes {
		val = val*10 + int(b-48)
	}

	if neg {
		val = -val
	}

	return val
}

func Ceil[T constraints.Float | constraints.Integer](a, b T) T {
	af := float64(a)
	bf := float64(b)

	return T(math.Ceil(af / bf))
}

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

func LenSyncMap(m *sync.Map) int {
	var i int
	m.Range(func(k, v interface{}) bool {
		i++
		return true
	})
	return i
}
