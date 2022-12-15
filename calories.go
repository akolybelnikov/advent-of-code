// Package advent Day 1
package advent

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Elf struct {
	Calories int
}

func (e *Elf) AddCalories(num []byte) error {
	cal, err := strconv.Atoi(string(num))
	if err != nil {
		return err
	}
	e.Calories += cal

	return nil
}

func FindCaloriesBytes(data []byte) (int, error) {
	var nLine, prevIdx int
	var elves []*Elf
	var elf = new(Elf)

	for byteIndex, b := range data {
		if b == newline {
			nLine++
			if nLine == 1 && prevIdx != byteIndex {
				num := data[prevIdx:byteIndex]
				err := elf.AddCalories(num)
				if err != nil {
					return 0, err
				}
			}
			prevIdx = byteIndex + 1
			continue
		} else {
			if nLine > 1 {
				elves = append(elves, elf)
				elf = new(Elf)
			}
			nLine = 0
		}
	}

	return TopThreeTotal(elves), nil
}

func FindCaloriesStrings(filename string) int {
	var elves []*Elf

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var currPos int
	var emptyLine = -1
	var curr *Elf
	for scanner.Scan() {
		num := strings.Trim(scanner.Text(), " \n\t")
		if num != "" {
			if curr == nil {
				curr = new(Elf)
				emptyLine = 0
			}

			qty, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalln(err)
			}
			curr.Calories += qty
		} else {
			if emptyLine == 0 {
				elves = append(elves, curr)
				curr = nil
				currPos++
				emptyLine++
			}
		}
	}

	return TopThreeTotal(elves)
}

func sortElves(elves []*Elf) {
	sort.SliceStable(elves, func(i, j int) bool {
		return elves[i].Calories > elves[j].Calories
	})
}

func TopThreeTotal(elves []*Elf) (total int) {
	sortElves(elves)
	if len(elves) >= 3 {
		for i := 0; i < 3; i++ {
			total += elves[i].Calories
		}
	}

	return
}
