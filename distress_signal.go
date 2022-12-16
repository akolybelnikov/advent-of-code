package advent

import "sort"

const (
	VALUE = 0
	LIST  = 1
)

type element struct {
	kind   int
	value  int
	nested []*element
}

func HandlePacketsPart1(data *[]byte) int {
	var lineNumber, prevIdx, packetIdx int
	packetLines := make([]*[]byte, 0)
	var res int

	for byteIndex, b := range *data {
		if b == newline {
			lineNumber++
			if lineNumber == 1 && prevIdx != byteIndex {
				packetLine := (*data)[prevIdx:byteIndex]
				packetLines = append(packetLines, &packetLine)
			}
			prevIdx = byteIndex + 1
			if lineNumber == 2 {
				packetIdx++
				correct := Compare(&packetLines)
				if correct {
					res += packetIdx
				}
				packetLines = packetLines[:0]
			}
		} else {
			lineNumber = 0
		}
	}

	packetIdx++
	correct := Compare(&packetLines)
	if correct {
		res += packetIdx
	}

	return res
}

func HandlePacketsPart2(data *[]byte) int {
	var lineNumber, prevIdx int
	elements := make([]*element, 0)

	for byteIndex, b := range *data {
		if b == newline {
			lineNumber++
			if lineNumber == 1 && prevIdx != byteIndex {
				packetLine := (*data)[prevIdx:byteIndex]
				el, _ := parseList(&packetLine, 0)
				elements = append(elements, el)
			}
			prevIdx = byteIndex + 1
		} else {
			lineNumber = 0
		}
	}

	sort.Slice(elements, func(i, j int) bool { return compare(elements[i], elements[j]) > 0 })
	div1, _ := parseList(&[]byte{91, 91, 50, 93, 93}, 0)
	div2, _ := parseList(&[]byte{91, 91, 54, 93, 93}, 0)

	idx1, _ := sort.Find(len(elements), func(i int) int { return compare(elements[i], div1) })
	idx2, _ := sort.Find(len(elements), func(i int) int { return compare(elements[i], div2) })

	return (idx1 + 1) * (idx2 + 2)
}

func Compare(packet *[]*[]byte) bool {
	l, r := (*packet)[0], (*packet)[1]
	lel, _ := parseList(l, 0)
	rel, _ := parseList(r, 0)

	return compare(lel, rel) > 0
}

func parseList(bytes *[]byte, i int) (*element, int) {
	var nested []*element
	i++
	for (*bytes)[i] != 93 {
		if (*bytes)[i] == 91 {
			el, j := parseList(bytes, i)
			nested = append(nested, el)
			i = j
		} else if (*bytes)[i] == 44 {
			i++
		} else {
			val, j := parseValue(bytes, i)
			nested = append(nested, &element{kind: VALUE, value: val})
			i = j
		}
	}
	i++

	return &element{kind: LIST, nested: nested}, i
}

func parseValue(bytes *[]byte, i int) (int, int) {
	var val int
	for (*bytes)[i] != 44 && (*bytes)[i] != 93 {
		val = val*10 + int((*bytes)[i]-48)
		i++
	}

	return val, i
}

func compare(l, r *element) int {
	if l.kind == VALUE && r.kind == VALUE {
		return r.value - l.value
	}

	if l.kind == LIST && r.kind == LIST {
		for i := 0; i < len(l.nested) && i < len(r.nested); i++ {
			res := compare(l.nested[i], r.nested[i])
			if res != 0 {
				return res
			}
		}
		return len(r.nested) - len(l.nested)
	}

	if l.kind == VALUE && r.kind == LIST {
		return compare(&element{kind: LIST, nested: []*element{l}}, r)
	}

	if l.kind == LIST && r.kind == VALUE {
		return compare(l, &element{kind: LIST, nested: []*element{r}})
	}

	return 0
}
