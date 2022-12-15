package advent

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Monkeys map[int]*monkey
type operator string

type op struct {
	action operator
	value  int
}

type opTest struct {
	divider int
	t       int
	f       int
}

type monkey struct {
	Count int
	Items *[]int
	op    *op
	test  *opTest
}

const (
	multiply operator = "*"
	add      operator = "+"
)

func newMonkey() *monkey {
	return &monkey{}
}

func (m *monkey) processLines(lines []*[]byte) (int, error) {
	id, err := findId(lines[0])
	if err != nil {
		return 0, err
	}

	startingItems, err := findStartingItems(lines[1])
	if err != nil {
		return 0, err
	}
	m.Items = startingItems

	operation, err := findOp(lines[2])
	if err != nil {
		return 0, err
	}
	m.op = operation

	test, err := findTest(lines[3:])
	if err != nil {
		return 0, err
	}
	m.test = test

	return id, nil
}

func (m *monkey) throwItem(item int, f int) (int, int, error) {
	worryLevel, err := findWorryLevel(item, m.op)
	if err != nil {
		return 0, 0, err
	}
	worryLevel = worryLevel % f
	to := m.test.testItem(worryLevel)

	return to, worryLevel, nil
}

func (o *opTest) testItem(item int) int {
	if item%o.divider == 0 {
		return o.t
	} else {
		return o.f
	}
}

func (m *Monkeys) Round(factor int) error {
	for i := 0; i < len(*m); i++ {
		mk := (*m)[i]
		for _, item := range *mk.Items {
			to, it, err := mk.throwItem(item, factor)
			if err != nil {
				return err
			}
			*(*m)[to].Items = append(*(*m)[to].Items, it)
			mk.Count++
		}

		items := make([]int, 0)
		mk.Items = &items
	}

	return nil
}

func (m *Monkeys) Level() int {
	counts := make([]int, len(*m))
	for _, mk := range *m {
		counts = append(counts, mk.Count)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	return counts[0] * counts[1]
}

func handleData(data *[]byte) (*Monkeys, error) {
	var m = make(Monkeys)
	var line, prevIdx int
	var curMonkey = newMonkey()
	monkeyLines := make([]*[]byte, 0)

	for byteIndex, b := range *data {
		if b == newline {
			line++
			if line == 1 && prevIdx != byteIndex {
				monkeyLine := (*data)[prevIdx:byteIndex]
				monkeyLines = append(monkeyLines, &monkeyLine)
			}
			prevIdx = byteIndex + 1
			if line == 2 {
				id, err := curMonkey.processLines(monkeyLines)
				if err != nil {
					return nil, err
				}
				m[id] = curMonkey
				curMonkey = newMonkey()
				monkeyLines = make([]*[]byte, 0)
			}
		} else {
			line = 0
		}
	}

	id, err := curMonkey.processLines(monkeyLines)
	if err != nil {
		return nil, err
	}
	m[id] = curMonkey

	return &m, nil
}

func findId(line *[]byte) (int, error) {
	id, err := strconv.Atoi(string((*line)[7 : len(*line)-1]))
	if err != nil {
		return 0, err
	}

	return id, nil
}

func findStartingItems(line *[]byte) (*[]int, error) {
	items := make([]int, 0)
	itemsString := strings.Split(string(*line), ":")[1]
	for _, entry := range strings.Split(itemsString, ",") {
		entry = strings.TrimSpace(entry)
		item, err := strconv.Atoi(entry)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return &items, nil
}

func findOp(line *[]byte) (*op, error) {
	opString := strings.Split(string(*line), "=")[1]
	opItems := strings.Split(strings.TrimSpace(opString), " ")
	value := 0
	if opItems[2] != "old" {
		val, err := strconv.Atoi(opItems[2])
		if err != nil {
			return nil, err
		}
		value = val
	}

	return &op{value: value, action: operator(opItems[1])}, nil
}

func findTest(lines []*[]byte) (*opTest, error) {
	intValues := [2]int{}
	dividerString := strings.Split(strings.TrimSpace(string(*lines[0])), " ")[3]
	divider, err := strconv.Atoi(dividerString)
	if err != nil {
		return nil, err
	}
	for i, line := range lines[1:] {
		str := strings.Split(strings.TrimSpace(string(*line)), " ")[5]
		val, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		intValues[i] = val
	}

	return &opTest{
		divider: divider,
		t:       intValues[0],
		f:       intValues[1],
	}, nil
}

func MonkeyBusiness(data *[]byte) (*Monkeys, int, error) {
	monkeys, err := handleData(data)
	var factor = 1
	for _, mk := range *monkeys {
		factor *= mk.test.divider
	}
	return monkeys, factor, err
}

func findWorryLevel(item int, op *op) (int, error) {
	var val = op.value
	if val == 0 {
		val = item
	}
	switch op.action {
	case add:
		return item + val, nil
	case multiply:
		return item * val, nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", op.action)
	}
}
