package day_16

import (
	utils "github.com/akolybelnikov/advent-of-code"
	"github.com/yourbasic/graph"
	"math"
	"sort"
)

type Valve struct {
	ID        int
	byteID    [2]byte
	byteLeads [][2]byte
	leads     []int
	rate      int
	weights   map[int]int
}

type System map[int]*Valve

//type ActiveSystem map[int]*Valve

type State struct {
	id    int
	id2   int
	tick  int
	tick2 int
	flow  int
	path  [64]bool
}

func ParseLines(lines *[]*[]byte) *System {
	vs := make([]*Valve, 0)
	sys := make(System)
	temp := make(map[[2]byte]int)

	for _, line := range *lines {
		id := [2]byte{(*line)[6], (*line)[7]}
		var v = &Valve{
			byteID: id,
		}

		i := 23
		b := (*line)[i]
		for b != 59 {
			i++
			b = (*line)[i]
		}
		vi := (*line)[23:i]
		v.rate = utils.BytesToInt(&vi)
		e := 0
		for e != 5 {
			if (*line)[i] == 32 {
				e++
			}
			i++
		}

		leads := (*line)[i:]
		if len(leads) > 0 {
			for idx := 0; idx < len(leads); idx += 4 {
				lead := [2]byte{leads[idx], leads[idx+1]}
				v.byteLeads = append(v.byteLeads, lead)
			}
		}
		vs = append(vs, v)
	}

	sort.Slice(vs, func(i, j int) bool {
		left, right := vs[i], vs[j]
		return left.byteID[0] < right.byteID[0]
	})

	sort.Slice(vs, func(i, j int) bool {
		left, right := vs[i], vs[j]
		return left.byteID[1] < right.byteID[1]
	})

	for i, v := range vs {
		v.ID = i
		sys[i] = v
		temp[v.byteID] = i
	}

	for _, v := range sys {
		v.leads = make([]int, len(v.byteLeads))
		for i, lead := range v.byteLeads {
			v.leads[i] = temp[lead]
		}
	}

	g := sys.Graph()
	for id, v := range sys {
		if v.rate > 0 || id == 0 {
			v.weights = make(map[int]int)
			for bid, bv := range sys {
				if bid != id && bv.rate > 0 {
					_, dist := graph.ShortestPath(g, v.ID, bv.ID)
					v.weights[bv.ID] = int(dist) + 1
				}
			}
		}
	}

	return &sys
}

func (s *System) Graph() *graph.Mutable {
	g := graph.New(len(*s))
	for _, v := range *s {
		for _, lead := range v.leads {
			g.AddBothCost(v.ID, (*s)[lead].ID, 1)
		}
	}

	return g
}

func (s *System) versions(state *State) []*State {
	versions := make([]*State, 0)
	for id, v := range *s {
		if v.rate > 0 && !state.path[id] {
			cost := (*s)[state.id].weights[id]
			if state.tick+cost <= 30 {
				np := state.path
				np[id] = true
				st := &State{
					id:   id,
					tick: state.tick + cost,
					path: np,
					flow: state.flow + (30-(state.tick+cost))*(*s)[id].rate,
				}
				versions = append(versions, st)
			}
		}
	}

	return versions
}

func (s *System) versions2(state *State, space *map[State]int) []*State {
	versions := make([]*State, 0)
	for id1 := 1; id1 < len(*s); id1++ {
		if (*s)[id1].rate == 0 {
			continue
		}
		offset := 0
		if state.id == 0 && state.id2 == 0 {
			offset = id1
		}
		for id2 := 1 + offset; id2 < len(*s); id2++ {
			if (*s)[id2].rate == 0 {
				continue
			}
			if id1 != id2 {
				v1 := (*s)[id1]
				v2 := (*s)[id2]
				cost1 := (*s)[state.id].weights[id1]
				cost2 := (*s)[state.id2].weights[id2]
				visited := state.path[v1.ID] || state.path[v2.ID]
				if !visited && state.tick+cost1 <= 26 && state.tick2+cost2 <= 26 {
					flow1 := (26 - (state.tick + cost1)) * v1.rate
					flow2 := (26 - (state.tick2 + cost2)) * v2.rate
					np := state.path
					np[v1.ID] = true
					np[v2.ID] = true
					flow := state.flow + flow1 + flow2
					sp := State{
						id:   v1.ID,
						id2:  v2.ID,
						path: np,
					}
					if prevFlow, ok := (*space)[sp]; !ok || prevFlow < flow {
						(*space)[sp] = flow
						newState := &State{
							id:    v1.ID,
							id2:   v2.ID,
							tick:  state.tick + cost1,
							tick2: state.tick2 + cost2,
							flow:  flow,
							path:  np,
						}
						versions = append(versions, newState)
					}
				}
			}
		}
	}

	return versions
}

func FindMaxFlow(s *System) int {
	state := &State{
		id:   0,
		tick: 0,
		flow: 0,
		path: [64]bool{},
	}

	queue := []*State{state}
	maxFlow := math.MinInt
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		maxFlow = utils.Max(maxFlow, cur.flow)
		vs := s.versions(cur)
		for _, v := range vs {
			queue = append(queue, v)
		}
	}

	return maxFlow
}

func FindMaxFlow2(s *System) int {
	space := make(map[State]int)
	state := &State{
		id:    0,
		id2:   0,
		tick:  0,
		tick2: 0,
		flow:  0,
		path:  [64]bool{},
	}

	queue := []*State{state}
	maxFlow := math.MinInt
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		maxFlow = utils.Max(maxFlow, cur.flow)
		vs := s.versions2(cur, &space)
		for _, v := range vs {
			queue = append(queue, v)
		}
	}

	return maxFlow
}
