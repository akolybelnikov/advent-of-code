package day_20

type list map[int]*node

type node struct {
	val  int
	next *node
	prev *node
}

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
	m := len(nums) - 1

	for nodeIdx, _ := range nums {
		cur := mix[nodeIdx]
		temp := cur
		if cur.val == 0 {
			continue
		}
		if cur.val > 0 {
			for i := 0; i < cur.val%m; i++ {
				temp = temp.next
			}
		} else {
			for i := 0; i >= cur.val%m; i-- {
				temp = temp.prev
			}
		}
		if temp == cur {
			continue
		}
		update(cur, temp)
	}

	return mix.coordinates(idx0)
}

func (m *list) link(nums *[]int) {
	for i, _ := range *nums {
		(*m)[i].next = (*m)[(i+1)%len(*m)]
		(*m)[i].prev = (*m)[(i-1+len(*m))%len(*m)]
	}
}

func update(n *node, t *node) {
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
