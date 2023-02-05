package day_19

import utils "github.com/akolybelnikov/advent-of-code"

type recipe [][2]int
type blueprint struct {
	id       int
	robots   []*recipe
	maxSpend [3]int
}

const (
	ORE = iota
	CLAY
	OBSIDIAN
)

func MakeBlueprints(data *[]*[]byte) []*blueprint {
	res := make([]*blueprint, 0)
	for _, line := range *data {
		res = append(res, parseBytes(line))
	}

	return res
}

func parseBytes(line *[]byte) *blueprint {
	res := &blueprint{robots: make([]*recipe, 4)}
	*line = (*line)[10:]
	var prev int
	for cur, b := range *line {
		// find ID
		if b == 58 {
			id := (*line)[prev:cur]
			res.id = utils.BytesToInt(&id)
			prev = cur + 2
			break
		}
	}
	*line = (*line)[prev:]
	// find the costs
	prev = 0
	robots := make([]*[]byte, 0)
	for cur, b := range *line {
		if b == 46 {
			robotBytes := (*line)[prev:cur]
			robots = append(robots, &robotBytes)
			prev = cur + 2
		}
	}
	// find the costs of one ore robot
	oreRobot := recipe{}
	r := [2]int{0, ORE}
	prev = 0
	space := 0
	oreRobotBytes := *robots[0]
	for cur, b := range oreRobotBytes {
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
			cb := oreRobotBytes[prev:cur]
			r[0] = utils.BytesToInt(&cb)
			oreRobot = append(oreRobot, r)
			res.robots[0] = &oreRobot
			res.maxSpend[ORE] = utils.Max(res.maxSpend[ORE], r[0])
			break
		}
	}
	// find the costs of one clay robot
	clayRobot := recipe{}
	cr := [2]int{0, ORE}
	prev = 0
	space = 0
	clayRobotBytes := *robots[1]
	for cur, b := range clayRobotBytes {
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
			cb := clayRobotBytes[prev:cur]
			cr[0] = utils.BytesToInt(&cb)
			clayRobot = append(clayRobot, cr)
			res.robots[1] = &clayRobot
			res.maxSpend[ORE] = utils.Max(res.maxSpend[ORE], cr[0])
			break
		}
	}
	// find the costs of one obsidian robot
	obsidianRobot := recipe{}
	oro := [2]int{0, ORE}
	orc := [2]int{0, CLAY}
	prev = 0
	space = 0
	obsidianRobotBytes := *robots[2]
	for cur, b := range obsidianRobotBytes {
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
			cb := obsidianRobotBytes[prev:cur]
			oro[0] = utils.BytesToInt(&cb)
			obsidianRobot = append(obsidianRobot, oro)
			res.maxSpend[ORE] = utils.Max(res.maxSpend[ORE], oro[0])
			continue
		}
		if space == 7 {
			prev = cur + 1
			continue
		}
		if space == 8 {
			cb := obsidianRobotBytes[prev:cur]
			orc[0] = utils.BytesToInt(&cb)
			obsidianRobot = append(obsidianRobot, orc)
			res.robots[2] = &obsidianRobot
			res.maxSpend[CLAY] = utils.Max(res.maxSpend[CLAY], orc[0])
			break
		}
	}
	// find the costs of one geode robot
	geodeRobot := recipe{}
	gro := [2]int{0, ORE}
	grob := [2]int{0, OBSIDIAN}
	prev = 0
	space = 0
	geodeRobotBytes := *robots[3]
	for cur, b := range geodeRobotBytes {
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
			cb := geodeRobotBytes[prev:cur]
			gro[0] = utils.BytesToInt(&cb)
			geodeRobot = append(geodeRobot, gro)
			res.maxSpend[ORE] = utils.Max(res.maxSpend[ORE], gro[0])
			continue
		}
		if space == 7 {
			prev = cur + 1
			continue
		}
		if space == 8 {
			cb := geodeRobotBytes[prev:cur]
			grob[0] = utils.BytesToInt(&cb)
			geodeRobot = append(geodeRobot, grob)
			res.robots[3] = &geodeRobot
			res.maxSpend[OBSIDIAN] = utils.Max(res.maxSpend[OBSIDIAN], grob[0])
			break
		}
	}

	return res
}
