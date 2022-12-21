package day_16

import (
	utils "github.com/akolybelnikov/advent-of-code"
	"github.com/yourbasic/graph"
	"golang.org/x/exp/maps"
	"math"
	"sort"
)

type Valve struct {
	IID     int
	id      [2]byte
	leads   [][2]byte
	rate    int
	weights map[int]int
}

type System map[[2]byte]*Valve
type ActiveSystem map[int]*Valve

type State struct {
	id    int
	id2   int
	tick  int
	tick2 int
	flow  int
	path  [64]bool
}

func ParseLines(lines *[]*[]byte) *System {
	sys := make(System)
	iid := 0
	for _, line := range *lines {
		id := [2]byte{(*line)[6], (*line)[7]}
		var v = &Valve{
			id:  id,
			IID: iid,
		}
		iid++
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
				v.leads = append(v.leads, lead)
			}
		}

		sys[id] = v
	}

	return &sys
}

func (s *System) Active() *ActiveSystem {
	g := s.Graph()
	var sys = make(ActiveSystem)
	for id, v := range *s {
		if v.rate > 0 || id == [2]byte{65, 65} {
			active := &Valve{IID: v.IID, id: v.id, rate: v.rate, weights: make(map[int]int)}
			for bid, bv := range *s {
				if bid != id && bv.rate > 0 {
					_, dist := graph.ShortestPath(g, v.IID, bv.IID)
					active.weights[bv.IID] = int(dist) + 1
				}
			}
			sys[v.IID] = active
		}
	}

	rs := sys.remap()

	return rs
}

func (s *System) Graph() *graph.Mutable {
	g := graph.New(len(*s))
	for _, v := range *s {
		for _, lead := range v.leads {
			g.AddBothCost(v.IID, (*s)[lead].IID, 1)
		}
	}

	return g
}

func (s *ActiveSystem) versions(state *State) []*State {
	versions := make([]*State, 0)
	for id, v := range *s {
		cost := (*s)[state.id].weights[id]
		if v.rate > 0 && !state.path[v.IID] && state.tick+cost <= 30 {
			np := state.path
			np[v.IID] = true
			st := &State{
				id:   v.IID,
				tick: state.tick + cost,
				path: np,
				flow: state.flow + (30-(state.tick+cost))*v.rate,
			}
			versions = append(versions, st)
		}
	}

	return versions
}

func (s *ActiveSystem) remap() *ActiveSystem {
	vs := maps.Keys(*s)
	sort.Slice(vs, func(i, j int) bool {
		return (*s)[vs[i]].id[0] < (*s)[vs[j]].id[0] && (*s)[vs[i]].id[1] < (*s)[vs[j]].id[1]
	})
	np := make(map[int]int)
	ns := make(ActiveSystem)

	for idx, k := range vs {
		np[k] = idx
		(*s)[k].IID = idx
	}

	for k, v := range *s {
		newWeights := make(map[int]int)
		for k2, v2 := range v.weights {
			newWeights[np[k2]] = v2
		}
		v.weights = newWeights
		ns[np[k]] = v
	}

	return &ns
}

func (s *ActiveSystem) versions2(state *State, space *map[State]int) []*State {
	versions := make([]*State, 0)
	for id1 := 1; id1 < len(*s); id1++ {
		offset := 0
		if state.id == 0 && state.id2 == 0 {
			offset = id1
		}
		for id2 := 1 + offset; id2 < len(*s); id2++ {
			if id1 != id2 {
				v1 := (*s)[id1]
				v2 := (*s)[id2]
				cost1 := (*s)[state.id].weights[id1]
				cost2 := (*s)[state.id2].weights[id2]
				visited := state.path[v1.IID] || state.path[v2.IID]
				if !visited && state.tick+cost1 <= 26 && state.tick2+cost2 <= 26 {
					flow1 := (26 - (state.tick + cost1)) * v1.rate
					flow2 := (26 - (state.tick2 + cost2)) * v2.rate
					np := state.path
					np[v1.IID] = true
					np[v2.IID] = true
					flow := state.flow + flow1 + flow2
					sp := State{
						id:   v1.IID,
						id2:  v2.IID,
						path: np,
					}
					if prevFlow, ok := (*space)[sp]; !ok || prevFlow < flow {
						(*space)[sp] = flow
						newState := &State{
							id:    v1.IID,
							id2:   v2.IID,
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

func FindMaxFlow(s *ActiveSystem) int {
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

func FindMaxFlow2(s *ActiveSystem) int {
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
