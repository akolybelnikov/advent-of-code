package day_19

import utils "github.com/akolybelnikov/advent-of-code"

type Blueprint struct {
	ID            int
	OreRobot      Cost
	ClayRobot     Cost
	ObsidianRobot Cost
	GeodeRobot    Cost
}

type Cost struct {
	Ore      int
	Clay     int
	Obsidian int
}

func MakeBlueprints(data *[]*[]byte) []*Blueprint {
	res := make([]*Blueprint, 0)
	for _, line := range *data {
		res = append(res, blueprint(line))
	}

	return res
}

func blueprint(line *[]byte) *Blueprint {
	res := &Blueprint{}
	*line = (*line)[10:]
	var prev int
	for cur, b := range *line {
		// find ID
		if b == 58 {
			id := (*line)[prev:cur]
			res.ID = utils.BytesToInt(&id)
			prev = cur + 2
			break
		}
	}
	*line = (*line)[prev:]
	// find costs
	prev = 0
	costs := make([]*[]byte, 0)
	for cur, b := range *line {
		if b == 46 {
			costBytes := (*line)[prev:cur]
			costs = append(costs, &costBytes)
			prev = cur + 2
		}
	}
	// find the costs of one ore robot
	oreRobotCost := Cost{}
	prev = 0
	space := 0
	cost0 := *costs[0]
	for cur, b := range cost0 {
		if b == 32 {
			space++
		} else {
			continue
		}
		if space == 4 {
			prev = cur + 1
			continue
		}
		if space == 5 {
			cb := cost0[prev:cur]
			oreRobotCost.Ore = utils.BytesToInt(&cb)
			res.OreRobot = oreRobotCost
			break
		}
	}
	// find the costs of one clay robot
	clayRobotCost := Cost{}
	prev = 0
	space = 0
	cost1 := *costs[1]
	for cur, b := range cost1 {
		if b == 32 {
			space++
		} else {
			continue
		}
		if space == 4 {
			prev = cur + 1
			continue
		}
		if space == 5 {
			cb := cost1[prev:cur]
			clayRobotCost.Ore = utils.BytesToInt(&cb)
			res.ClayRobot = clayRobotCost
			break
		}
	}
	// find the costs of one obsidian robot
	obsidianRobotCost := Cost{}
	prev = 0
	space = 0
	cost2 := *costs[2]
	for cur, b := range cost2 {
		if b == 32 {
			space++
		} else {
			continue
		}
		if space == 4 {
			prev = cur + 1
			continue
		}
		if space == 5 {
			cb := cost2[prev:cur]
			obsidianRobotCost.Ore = utils.BytesToInt(&cb)
			continue
		}
		if space == 7 {
			prev = cur + 1
			continue
		}
		if space == 8 {
			cb := cost2[prev:cur]
			obsidianRobotCost.Clay = utils.BytesToInt(&cb)
			res.ObsidianRobot = obsidianRobotCost
			break
		}
	}
	// find the costs of one geode robot
	geodeRobotCost := Cost{}
	prev = 0
	space = 0
	cost3 := *costs[3]
	for cur, b := range cost3 {
		if b == 32 {
			space++
		} else {
			continue
		}
		if space == 4 {
			prev = cur + 1
			continue
		}
		if space == 5 {
			cb := cost3[prev:cur]
			geodeRobotCost.Ore = utils.BytesToInt(&cb)
			continue
		}
		if space == 7 {
			prev = cur + 1
			continue
		}
		if space == 8 {
			cb := cost3[prev:cur]
			geodeRobotCost.Obsidian = utils.BytesToInt(&cb)
			res.GeodeRobot = geodeRobotCost
			break
		}
	}

	return res
}
