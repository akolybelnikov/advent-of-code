package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func findHouse(target int, multiplier int, maxVisits int, workers int) int {
	size := 100_000

	for {
		fmt.Printf("Trying with size: %d\n", size)
		presents := make([]int, size)
		ctx, cancel := context.WithCancel(context.Background())
		var wg sync.WaitGroup

		chunk := size / workers
		for i := 0; i < workers; i++ {
			start := i * chunk
			end := start + chunk
			if i == workers-1 {
				end = size
			}

			wg.Add(1)
			go func(start, end int) {
				defer wg.Done()
				for elf := start; elf < end; elf++ {
					visits := 0
					for house := elf; house < size; house += elf {
						if maxVisits > 0 && visits >= maxVisits {
							break
						}
						select {
						case <-ctx.Done():
							return
						default:
							presents[house] += elf * multiplier
							visits++
						}
					}
				}
			}(start+1, end)
		}

		// Wait for all workers
		wg.Wait()
		cancel()

		for house, total := range presents {
			if total >= target {
				return house
			}
		}

		size *= 2
	}
}

func main() {
	data, err := os.ReadFile("inputs/day20.txt")
	if err != nil {
		fmt.Println("Failed to read input file.", err)
		os.Exit(0)
	}
	input := strings.TrimSpace(string(data))
	target, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Failed to parse input.", err)
	}
	start := time.Now()

	// Part 1: elves deliver 10 * elf# to every house
	house1 := findHouse(target, 10, 0, 8)
	fmt.Printf("Part 1 result: %d (in %v)\n", house1, time.Since(start))

	// Part 2: elves deliver 11 * elf#, but only to 50 houses each
	start = time.Now()
	house2 := findHouse(target, 11, 50, 8)
	fmt.Printf("Part 2 result: %d (in %v)\n", house2, time.Since(start))
}
