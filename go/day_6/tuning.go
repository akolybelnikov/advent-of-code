package day_6

const newline = 10

type markerHandler func(*[]byte) int

type void struct{}

type marker interface {
	[4]byte | [14]byte
}

var member void

func FindFirstMarker(data *[]byte) int {
	var marker [4]byte
	return findMarker(data, &marker)
}

func FindFirstMessage(data *[]byte) int {
	var marker [14]byte
	return findMarker(data, &marker)
}

func findMarker[M marker](data *[]byte, marker *M) int {
	idx := 0
	l := len(*marker)
	for i, b := range *data {
		idx++
		if i < l {
			(*marker)[i] = b
		} else {
			for j := 0; j < l-1; j++ {
				(*marker)[j] = (*marker)[j+1]
			}
			(*marker)[l-1] = b
		}

		if i > l-1 {
			if ok := isMarker(marker, l); ok {
				break
			}
		}
	}
	return idx
}

func isMarker[M marker](marker *M, l int) bool {
	set := make(map[byte]void)
	for i := 0; i < l; i++ {
		set[(*marker)[i]] = member
	}

	return len(set) == l
}

func HandleMarkers(data *[]byte, fn markerHandler) *[]int {
	var nLine, prevIdx int
	var indices []int

	for byteIndex, b := range *data {
		if b == newline {
			nLine++
			if nLine == 1 && prevIdx != byteIndex {
				line := (*data)[prevIdx:byteIndex]
				idx := fn(&line)
				indices = append(indices, idx)
			}
			prevIdx = byteIndex + 1
		} else {
			nLine = 0
		}
	}

	return &indices
}
