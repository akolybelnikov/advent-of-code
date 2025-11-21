package main

import (
	"fmt"
	"github.com/akolybelnikov/aoc-2015/internal/utils"
	"os"
	"strconv"
	"strings"
)

type opType int

const (
	NOOP   opType = 0
	AND           = 'A'
	OR            = 'O'
	NOT           = 'N'
	LSHIFT        = 'L'
	RSHIFT        = 'R'
)

type instruction struct {
	op   opType
	lval uint16
	rval uint16
	wire string
}

type wireSet struct {
	circuit map[string]uint16
	queue   []string
}

func main() {
	data, err := os.ReadFile("inputs/day07.txt")
	if err != nil {
		fmt.Println("Failed to read input file.", err)
		os.Exit(0)
	}
	input := string(data)

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input))

	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input))

	os.Exit(0)
}

// part one
func part1(input string) uint16 {
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)

	ws := runCircuit(lines, 0)

	return ws.circuit["a"]
}

// part two
func part2(input string) uint16 {
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)

	ws := runCircuit(lines, 0)
	a := ws.circuit["a"]

	wsr := runCircuit(lines, a)

	return wsr.circuit["a"]
}

func runCircuit(lines []string, repl uint16) *wireSet {
	ws := wireSet{
		circuit: make(map[string]uint16),
		queue:   make([]string, 0),
	}

	for _, line := range lines {
		if line == "1674 -> b" && repl != 0 {
			line = fmt.Sprintf("%d -> b", repl)
		}
		i, err := ws.parseInstruction(line)
		if err != nil {
			ws.queue = append(ws.queue, line)
			continue
		}
		ws.exec(i)
	}

	for len(ws.queue) > 0 {
		line := ws.queue[0]
		ws.queue = ws.queue[1:]
		i, err := ws.parseInstruction(line)
		if err != nil {
			ws.queue = append(ws.queue, line)
			continue
		}
		ws.exec(i)
	}

	return &ws
}

func (w *wireSet) exec(instr *instruction) {
	switch instr.op {
	case NOOP:
		w.circuit[instr.wire] = instr.rval
	case AND:
		w.circuit[instr.wire] = instr.lval & instr.rval
	case OR:
		w.circuit[instr.wire] = instr.lval | instr.rval
	case NOT:
		w.circuit[instr.wire] = ^instr.rval
	case LSHIFT:
		w.circuit[instr.wire] = instr.lval << instr.rval
	case RSHIFT:
		w.circuit[instr.wire] = instr.lval >> instr.rval
	default:

	}
}

func (w *wireSet) parseInstruction(line string) (*instruction, error) {
	i := instruction{}
	fields := strings.Fields(line)
	switch len(fields) {
	case 3:
		rval, err := w.parseValue(fields[0])
		if err != nil {
			return nil, fmt.Errorf("invalid instruction: %s", line)
		}
		i.op = NOOP
		i.rval = rval
		i.wire = fields[2]
	case 4:
		rval, err := w.parseValue(fields[1])
		if err != nil {
			return nil, fmt.Errorf("invalid instruction: %s", line)
		}
		i.op = NOT
		i.rval = rval
		i.wire = fields[3]
	default:
		lval, err := w.parseValue(fields[0])
		if err != nil {
			return nil, fmt.Errorf("invalid instruction: %s", line)
		}
		rval, err := w.parseValue(fields[2])
		if err != nil {
			return nil, fmt.Errorf("invalid instruction: %s", line)
		}
		i.op = opType(fields[1][0])
		i.lval = lval
		i.rval = rval
		i.wire = fields[4]
	}

	return &i, nil
}

// Helper function to parse a value or retrieve it from the circuit.
func (w *wireSet) parseValue(field string) (uint16, error) {
	val, err := strconv.Atoi(field)
	if err != nil {
		rval, ok := w.circuit[field]
		if !ok {
			return 0, fmt.Errorf("could not resolve value: %s", field)
		}
		return rval, nil
	}
	return uint16(val), nil
}
