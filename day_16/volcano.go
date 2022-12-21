package day_16

import (
	utils "github.com/akolybelnikov/advent-of-code"
	"github.com/yourbasic/graph"
	"golang.org/x/exp/maps"
	"math"
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
	id   int
	tick int
	flow int
	path map[int]bool
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
					//active.leads = append(active.leads, bid)
					_, dist := graph.ShortestPath(g, v.IID, bv.IID)
					//log.Printf("path from v %s to w %s is %v with weight %d", id, bid, path, dist)
					active.weights[bv.IID] = int(dist) + 1
				}
			}
			sys[v.IID] = active
		}
	}

	return &sys
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
			np := make(map[int]bool)
			maps.Copy(np, state.path)
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

func FindMaxFlow(s *ActiveSystem, start int) int {
	state := &State{
		id:   start,
		tick: 0,
		flow: 0,
		path: make(map[int]bool),
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
