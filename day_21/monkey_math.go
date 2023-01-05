package day_21

type name [4]byte

type monkey struct {
	name        name
	op          byte
	left, right chan int
	val         int
	deps        []*chan int
	leftVal     []int
	rightVal    []int
	waiting     bool
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

	for {
		g.run()
		if root.val != 0 {
			break
		}
	}

	return root.val
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

func (g *graph) run() {
	for _, m := range *g {
		if m.op == 0 && len(m.deps) > 0 {
			go m.yell()
		}
		if len(m.leftVal) == 0 && len(m.rightVal) == 0 && !m.waiting {
			m.waiting = true
			go m.wait()
		}
	}
}
