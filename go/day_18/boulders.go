package day_18

import (
	utils "github.com/akolybelnikov/advent-of-code"
	"math"
)

type surface [4]point

type point struct {
	x, y, z int
}

type points map[point]struct{}

type grid map[surface]int

func (ps *points) bounds() [2]point {
	var bounds [2]point
	bounds[0].x, bounds[0].y, bounds[0].z = math.MaxInt, math.MaxInt, math.MaxInt
	bounds[1].x, bounds[1].y, bounds[1].z = math.MinInt, math.MinInt, math.MinInt
	for p, _ := range *ps {
		bounds[0].x = utils.Min(bounds[0].x, p.x)
		bounds[0].y = utils.Min(bounds[0].y, p.y)
		bounds[0].z = utils.Min(bounds[0].z, p.z)
		bounds[1].x = utils.Max(bounds[1].x, p.x)
		bounds[1].y = utils.Max(bounds[1].y, p.y)
		bounds[1].z = utils.Max(bounds[1].z, p.z)

	}

	return bounds
}

func (ps *points) pop() *point {
	for el := range *ps {
		delete(*ps, el)
		return &el
	}
	return nil
}

func FindNotConnected(cubeBytes *[]*[]byte) int {
	var total = len(*cubeBytes) * 6
	g := make(grid)
	for _, cb := range *cubeBytes {
		s := makeCubeSurfaces(cb)
		for _, sfc := range s {
			g[sfc]++
		}
	}

	for _, s := range g {
		if s > 1 {
			total -= 2
		}
	}

	return total
}

func FindNotConnected2(cubeBytes *[]*[]byte) int {
	var total int
	var ps = make(points)
	for _, cb := range *cubeBytes {
		makeCubePoints(cb, &ps)
	}

	var airCubes = make(points)
	bounds := ps.bounds()
	findAirCubes(&ps, bounds, &airCubes)
	for p := range ps {
		for _, q := range p.neighbors() {
			_, contains := ps[q]
			_, ok := airCubes[q]
			if !contains && ok {
				total++
			}
		}
	}

	return total
}

func findAirCubes(ps *points, bounds [2]point, airCubes *points) {
	temp := make(points)
	p := findAirCube(ps, bounds)
	if p != nil {
		temp[*p] = struct{}{}
	}
	for len(temp) > 0 {
		p = temp.pop()
		if _, exists := (*ps)[*p]; !exists {
			if _, ok := (*airCubes)[*p]; !ok {
				if p.withinBounds(bounds) {
					(*airCubes)[*p] = struct{}{}
					for _, n := range p.neighbors() {
						temp[n] = struct{}{}
					}
				}
			}
		}
	}
}

func findAirCube(ps *points, bounds [2]point) *point {
	for x := bounds[0].x; x <= bounds[1].x; x++ {
		for y := bounds[0].y; y <= bounds[1].y; y++ {
			for z := bounds[0].z; z <= bounds[1].z; z++ {
				p := point{x, y, z}
				if _, exists := (*ps)[p]; !exists {
					return &p
				}
			}
		}
	}

	return nil
}

func (p *point) neighbors() []point {
	return []point{
		{p.x + 1, p.y, p.z},
		{p.x - 1, p.y, p.z},
		{p.x, p.y + 1, p.z},
		{p.x, p.y - 1, p.z},
		{p.x, p.y, p.z + 1},
		{p.x, p.y, p.z - 1},
	}
}

func (p *point) withinBounds(bounds [2]point) bool {
	return p.x >= bounds[0].x-1 && p.x <= bounds[1].x+1 &&
		p.y >= bounds[0].y-1 && p.y <= bounds[1].y+1 &&
		p.z >= bounds[0].z-1 && p.z <= bounds[1].z+1
}

func makeCubeSurfaces(bytes *[]byte) [6]surface {
	var cur, commaCount int
	var p7 = point{}
	for idx, b := range *bytes {
		if b == 44 {
			if commaCount == 0 {
				toInt := (*bytes)[cur:idx]
				p7.x = utils.BytesToInt(&toInt)
			} else {
				toInt := (*bytes)[cur:idx]
				p7.y = utils.BytesToInt(&toInt)
			}
			commaCount++
			cur = idx + 1
		}
	}
	toInt := (*bytes)[cur:]
	p7.z = utils.BytesToInt(&toInt)

	return findSurfaces(&p7)
}

func findSurfaces(p7 *point) [6]surface {
	var c [6]surface
	p0 := point{x: p7.x - 1, y: p7.y - 1, z: p7.z - 1}
	p1 := point{x: p7.x, y: p7.y - 1, z: p7.z - 1}
	p2 := point{x: p7.x, y: p7.y - 1, z: p7.z}
	p3 := point{x: p7.x - 1, y: p7.y - 1, z: p7.z}
	p4 := point{x: p7.x - 1, y: p7.y, z: p7.z}
	p5 := point{x: p7.x - 1, y: p7.y, z: p7.z - 1}
	p6 := point{x: p7.x, y: p7.y, z: p7.z - 1}

	bottom := surface{p0, p1, p2, p3}
	c[0] = bottom
	left := surface{p0, p5, p4, p3}
	c[1] = left
	back := surface{p0, p5, p6, p1}
	c[2] = back
	right := surface{p1, p6, *p7, p2}
	c[3] = right
	front := surface{p3, p4, *p7, p2}
	c[4] = front
	top := surface{p5, p6, *p7, p4}
	c[5] = top

	return c
}

func makeCubePoints(bytes *[]byte, ps *points) {
	var cur, commaCount int
	var p = point{}
	for idx, b := range *bytes {
		if b == 44 {
			if commaCount == 0 {
				toInt := (*bytes)[cur:idx]
				p.x = utils.BytesToInt(&toInt)
			} else {
				toInt := (*bytes)[cur:idx]
				p.y = utils.BytesToInt(&toInt)
			}
			commaCount++
			cur = idx + 1
		}
	}
	toInt := (*bytes)[cur:]
	p.z = utils.BytesToInt(&toInt)
	(*ps)[p] = struct{}{}
}
