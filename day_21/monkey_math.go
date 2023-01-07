package day_21

type name [4]byte

type monkey struct {
	deps        []*chan int
	left, right chan int
	leftVal     []int
	name        name
	op          byte
	rightVal    []int
	val         int
}

type graph map[name]*monkey

const (
	MUL = 42
	ADD = 43
	SUB = 45
	DIV = 47
)

func MonkeyMath(data *[]*[]byte) int {
	g := make(graph)
	parseMonkeyData(data, &g)
	r := name{'r', 'o', 'o', 't'}
	root := g[r]

	for _, m := range g {
		if m.op == 0 {
			go m.yell()
		}
	}

	res := <-root.signal()

	return res
}

func (m *monkey) signal() <-chan int {
	c := make(chan int, 1)

	go func() {
		defer close(c)
		for {
			if m.val != 0 {
				c <- m.val
				break
			}
		}

	}()

	return c
}

func (m *monkey) yell() {
	for _, dep := range m.deps {
		*dep <- m.val
	}
	m.deps = nil
}

func (m *monkey) wait() {
	for {
		select {
		case v := <-m.left:
			m.leftVal = append(m.leftVal, v)
		case v := <-m.right:
			m.rightVal = append(m.rightVal, v)
		}
		if len(m.leftVal) > 0 && len(m.rightVal) > 0 {
			switch m.op {
			case MUL:
				m.val = m.leftVal[0] * m.rightVal[0]
			case ADD:
				m.val = m.leftVal[0] + m.rightVal[0]
			case SUB:
				m.val = m.leftVal[0] - m.rightVal[0]
			case DIV:
				m.val = m.leftVal[0] / m.rightVal[0]
			}
			go m.yell()
			break
		}
	}
}
