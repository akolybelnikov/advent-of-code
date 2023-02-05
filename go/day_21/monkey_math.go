package day_21

type name [4]byte

type monkey struct {
	leftDeps, rightDeps       []*monkey
	left, right               chan int
	leftVal, rightVal         []int
	name, leftName, rightName name
	op                        byte
	val                       int
	depOnHuman                bool
}

type graph map[name]*monkey

const (
	MUL = 42
	ADD = 43
	SUB = 45
	DIV = 47
	EQL = 61
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

func MonkeyMath2(data *[]*[]byte) int {
	g := make(graph)
	parseMonkeyData(data, &g)
	r := name{'r', 'o', 'o', 't'}
	root := g[r]

	h := name{'h', 'u', 'm', 'n'}
	human := g[h]

	for _, m := range g {
		if m.op == 0 {
			go m.yell()
		}
	}

	<-root.signal()
	root.op = EQL

	left := g[root.leftName]
	right := g[root.rightName]

	if left.depOnHuman {
		root.val = right.val
	} else {
		root.val = left.val
	}

	for m := root; m.name != h; {
		leftM := g[m.leftName]
		rightM := g[m.rightName]

		if leftM.depOnHuman {
			leftM.val, _ = reverseOp(m.op, nil, &rightM.val, m.val)
			m = leftM
		} else {
			_, rightM.val = reverseOp(m.op, &leftM.val, nil, m.val)
			m = rightM
		}
	}

	return human.val
}

func reverseOp(op byte, left, right *int, target int) (int, int) {
	switch op {
	case ADD:
		if left == nil {
			return target - *right, *right
		}
		return *left, target - *left
	case SUB:
		if left == nil {
			return *right + target, *right
		}
		return *left, *left - target
	case MUL:
		if left == nil {
			return target / *right, *right
		}
		return *left, target / *left
	case DIV:
		if left == nil {
			return *right * target, *right
		}
		return *left, *left / target
	case EQL:
		if left == nil {
			return *right, *right
		}
		return *left, *left
	default:
		panic("unknown op")
	}
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
	for _, lm := range m.leftDeps {
		if m.depOnHuman {
			lm.depOnHuman = m.depOnHuman
		}
		lm.left <- m.val
	}

	for _, rm := range m.rightDeps {
		if m.depOnHuman {
			rm.depOnHuman = m.depOnHuman
		}
		rm.right <- m.val
	}
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
