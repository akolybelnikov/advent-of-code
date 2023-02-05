package day_10

import (
	utils "github.com/akolybelnikov/advent-of-code"
	"log"
	"strconv"
)

type cpuInstruction byte

type CPU struct {
	RegisterValue int
	Cycles        []int
}

const (
	addx cpuInstruction = 97
	hash                = 35
	dot                 = 46
)

func initCPU() *CPU {
	return &CPU{
		Cycles:        []int{1, 1},
		RegisterValue: 1,
	}
}

func (c *CPU) execInstruction(data *[]byte) error {
	if cpuInstruction((*data)[0]) == addx {
		val, err := strconv.Atoi(string((*data)[5:]))
		if err != nil {
			return err
		}
		c.Cycles = append(c.Cycles, c.RegisterValue)
		c.Cycles = append(c.Cycles, c.RegisterValue+val)
		c.RegisterValue += val
	} else {
		c.Cycles = append(c.Cycles, c.RegisterValue)
	}

	return nil
}

func SignalStrength(data *[]byte) (*CPU, error) {
	var line, prevIdx int
	c := initCPU()

	for byteIndex, b := range *data {
		if b == utils.NEWLINE {
			line++
			if line == 1 && prevIdx != byteIndex {
				nextInstruction := (*data)[prevIdx:byteIndex]
				err := c.execInstruction(&nextInstruction)
				if err != nil {
					return nil, err
				}
			}
			prevIdx = byteIndex + 1
		} else {
			line = 0
		}
	}

	return c, nil
}

func (c *CPU) ReadSignalStrength(cycles []int) ([]int, int) {
	var sumOfSignals int
	signals := make([]int, len(cycles))
	for i, n := range cycles {
		strength := c.Cycles[n] * n
		signals[i] = strength
		sumOfSignals += strength
	}

	return signals, sumOfSignals
}

func (c *CPU) RenderCRT() {
	pixels := make([]byte, len(c.Cycles))
	for i := 0; i < len(c.Cycles)-1; i++ {
		if i%40 >= c.Cycles[i+1]-1 && i%40 <= c.Cycles[i+1]+1 {
			pixels[i] = hash
		} else {
			pixels[i] = dot
		}
	}
	for i := 0; i < 6; i++ {
		start := i * 40
		log.Printf("%s", pixels[start:start+40])
	}
}
