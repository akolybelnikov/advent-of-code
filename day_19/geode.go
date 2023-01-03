package day_19

import (
	utils "github.com/akolybelnikov/advent-of-code"
	"log"
	"sync"
)

// Optimizations
// 1. Maximum spend-rate per recipe helps to limit what a recipe needs to produce over all cycles.
// 2. At each cycle, if you cannot possibly spend all accumulated resources, throw the surplus away.
//   This will keep more overlapping states with functionally the same values, and discard them.
// 3. Short-cut build-up cycles by calculating the needed amount of amounts over time

type state struct {
	robots  [4]int
	amounts [4]int
	cycle   int
}

type stateCache struct {
	lock   sync.RWMutex
	states map[state]struct{}
}

func initState(cycle int) *state {
	return &state{
		robots: [4]int{1, 0, 0, 0},
		cycle:  cycle,
	}
}

func Run(bps []*blueprint, cycles int) int {
	var wg sync.WaitGroup

	geodes := make(chan [2]int, len(bps))
	for _, bp := range bps {
		res := [2]int{bp.id}
		wg.Add(1)
		b := bp
		cache := &stateCache{states: make(map[state]struct{})}
		go func() {
			defer wg.Done()
			st := initState(cycles)
			res[1] = walk(b, st, cache)
			geodes <- res
		}()
	}
	go func() {
		wg.Wait()
		close(geodes)
	}()

	total := 0
	for res := range geodes {
		log.Printf("blueprint %d produce %d geodes\n", res[0], res[1])
		qLevel := res[0] * res[1]
		total += qLevel
	}

	return total
}

func Run2(bps []*blueprint, cycles int) int {
	var wg sync.WaitGroup

	geodes := make(chan [2]int, len(bps))
	for _, bp := range bps {
		res := [2]int{bp.id}
		wg.Add(1)
		b := bp
		cache := &stateCache{states: make(map[state]struct{})}
		go func() {
			defer wg.Done()
			st := initState(cycles)
			res[1] = walk(b, st, cache)
			geodes <- res
		}()
	}
	go func() {
		wg.Wait()
		close(geodes)
	}()

	total := 1
	for res := range geodes {
		log.Printf("blueprint %d produce %d geodes\n", res[0], res[1])
		total *= res[1]
	}

	return total
}

func walk(bp *blueprint, s *state, cache *stateCache) int {
	// if no more time left, return the geodes we already have.
	if s.cycle == 0 {
		return s.amounts[3]
	}
	// if current state exists in stateCache, return it.
	if ok := cache.load(*s); ok {
		return s.amounts[3]
	}
	// we have 5 options: build an ore botRecipe, build a clay botRecipe, build an obsidian botRecipe,
	//build a geode botRecipe, or do nothing.
	maxVal := s.amounts[3] + s.robots[3]*s.cycle

OUTER:
	for botType, botRecipe := range bp.robots {
		// if bot is geode bot, or we already have reached max needed botRecipe amount, skip.
		if botType != 3 && s.robots[botType] >= bp.maxSpend[botType] {
			continue
		}
		// try to calculate waiting time until we can build a bot.
		wait := 0
		for _, bot := range *botRecipe {
			// if we don't have any robots of particular type yet, no need to wait. skip to the next iteration.
			if s.robots[bot[1]] == 0 {
				continue OUTER
			}
			// waiting time is current amount of amt minus the cost of building divided by the number of
			// available bots. should we get a negative waiting time, we take 0 as the maximum.
			wait = utils.Max(wait, utils.Ceil(bot[0]-s.amounts[bot[1]], s.robots[bot[1]]))
		}
		// remaining cycles
		timeLeft := s.cycle - wait - 1
		if timeLeft <= 0 {
			continue
		}
		// next state values
		_bots := s.robots
		_amounts := [4]int{}
		for i, amt := range s.amounts {
			_amounts[i] = amt + s.robots[i]*(wait+1)
		}
		for _, _bot := range *botRecipe {
			_amounts[_bot[1]] -= _bot[0]
		}
		_bots[botType] += 1

		// Optimization 2 throws away the excess amounts
		// the amount of resources we need to hold on to, is the amount we consume pro round
		// the amount we consume is the maximum spend-rate over time left
		for i := 0; i < 3; i++ {
			_amounts[i] = utils.Min(_amounts[i], bp.maxSpend[i]*timeLeft)
		}
		// branch state
		newState := &state{
			robots:  _bots,
			amounts: _amounts,
			cycle:   timeLeft,
		}

		v := walk(bp, newState, cache)
		maxVal = utils.Max(maxVal, v)

	}

	cache.add(*s)

	return maxVal
}

func (c *stateCache) add(st state) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.states[st] = struct{}{}
}

func (c *stateCache) load(st state) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()
	_, exists := c.states[st]

	return exists
}
