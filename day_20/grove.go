package day_20

type list map[int]*node

type node struct {
	val  int
	next *node
	prev *node
}

const dk = 811589153

func GrovePositioning(arr *[]*[]byte) int {
	var idx0 int
	mix := make(list)
	nums := encryptedFile(arr)
	for i, num := range nums {
		if num == 0 {
			idx0 = i
		}
		mix[i] = &node{val: num}
	}

	mix.link(&nums)
	mix.move(len(nums))

	return mix.coordinates(idx0)
}

func GrovePositioning2(arr *[]*[]byte) int {
	var idx0 int
	mix := make(list)
	nums := encryptedFile(arr)
	for i, num := range nums {
		if num == 0 {
			idx0 = i
		}
		mix[i] = &node{val: num * dk}
	}

	mix.link(&nums)
	for i := 0; i < 10; i++ {
		mix.move(len(nums))
	}

	return mix.coordinates(idx0)
}

func (m *list) link(nums *[]int) {
	for i := range *nums {
		(*m)[i].next = (*m)[(i+1)%len(*m)]
		(*m)[i].prev = (*m)[(i-1+len(*m))%len(*m)]
	}
}

func (m *list) move(l int) {
	n := l - 1

	for nodeIdx := 0; nodeIdx < l; nodeIdx++ {
		cur := (*m)[nodeIdx]
		temp := cur
		if cur.val == 0 {
			continue
		}
		if cur.val > 0 {
			for i := 0; i < cur.val%n; i++ {
				temp = temp.next
			}
		} else {
			for i := 0; i >= cur.val%n; i-- {
				temp = temp.prev
			}
		}
		if temp == cur {
			continue
		}

		cur.update(temp)
	}
}

func (n *node) update(t *node) {
	n.next.prev = n.prev
	n.prev.next = n.next
	n.next = t.next
	t.next = n
	n.prev = t
	n.next.prev = n
}

func (m *list) coordinates(idx int) int {
	var res int
	cur := (*m)[idx]
	for i := 0; i < 1000%len(*m); i++ {
		cur = cur.next
	}
	res += cur.val
	cur = (*m)[idx]
	for i := 0; i < 2000%len(*m); i++ {
		cur = cur.next
	}
	res += cur.val
	cur = (*m)[idx]
	for i := 0; i < 3000%len(*m); i++ {
		cur = cur.next
	}
	res += cur.val

	return res
}
