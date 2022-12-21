package day_16_test

import (
	"fmt"
	utils "github.com/akolybelnikov/advent-of-code"
	"github.com/akolybelnikov/advent-of-code/day_16"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestValvesTestInput(t *testing.T) {
	data, _ := utils.ReadDataBytes("testdata/test_input.txt")
	arr, _ := utils.MakeBytesArray(&data)
	vs := day_16.ParseLines(arr)
	assert.Equal(t, 10, len(*vs))
	avs := vs.Active()
	assert.Equal(t, 7, len(*avs))
	start := (*vs)[[2]byte{65, 65}]
	flow := day_16.FindMaxFlow(avs, start.IID)
	assert.Equal(t, 1651, flow)
}

func TestValves(t *testing.T) {
	start := time.Now()
	data, _ := utils.ReadDataBytes("testdata/input.txt")
	arr, _ := utils.MakeBytesArray(&data)
	vs := day_16.ParseLines(arr)
	assert.Equal(t, 59, len(*vs))
	avs := vs.Active()
	assert.Equal(t, 16, len(*avs))
	s := (*vs)[[2]byte{65, 65}]
	flow := day_16.FindMaxFlow(avs, s.IID)
	assert.Equal(t, 1595, flow)
	fmt.Println(time.Since(start))
}
