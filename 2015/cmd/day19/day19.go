package main

import (
	"fmt"
	"github.com/akolybelnikov/aoc-2015/internal/utils"
	"os"
	"regexp"
	"sort"
	"strings"
)

// element represents a pair of runes, which can include a second optional rune for a two-character element.
type element [2]rune

// plant represents the nuclear fusion/fission with replacement rules, a molecule, and a collection of unique generated molecules.
type plant struct {
	replacements map[element][][]rune
	molecules    map[string]int
}

func main() {
	data, err := os.ReadFile("inputs/day19.txt")
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
func part1(input string) int {
	blocks, err := utils.ParseBlocksOfLines(input)
	utils.HandleErr(err)
	p := plant{
		replacements: replacements(blocks[0]),
		molecules:    make(map[string]int),
	}
	p.replace([]rune(blocks[1][0]))
	return len(p.molecules)
}

// replacements parses input strings into a map where keys are elements and values are lists of rune replacements.
func replacements(input []string) map[element][][]rune {
	res := make(map[element][][]rune)

	// Extracted function to parse element
	parseElement := func(token string) element {
		var el element
		runes := []rune(token)
		el[0] = runes[0]
		if len(runes) > 1 {
			el[1] = runes[1]
		}
		return el
	}

	for _, line := range input {
		parts := strings.Fields(line)
		if len(parts) < 3 {
			continue // Skip invalid input lines
		}

		fromElement := parseElement(parts[0])
		toRunes := []rune(parts[2])

		// Simplify map management by always appending
		res[fromElement] = append(res[fromElement], toRunes)
	}

	return res
}

// replace traverses the molecule and applies replacements based on the defined replacement rules.
func (p *plant) replace(molecule []rune) {
	for e, r := range p.replacements {
		for i, m := range molecule {
			if e[1] == 0 && m == e[0] {
				p.processVariants(r, i, i+1, molecule)
			} else if e[1] != 0 && m == e[0] && i < len(molecule)-1 && molecule[i+1] == e[1] {
				p.processVariants(r, i, i+2, molecule)
			}
		}
	}
}

// processVariants handles the replacement variants for a given molecule segment.
func (p *plant) processVariants(variants [][]rune, start, next int, molecule []rune) {
	for _, variant := range variants {
		p.buildAndAddMolecule(variant, start, next, molecule)
	}
}

// buildAndAddMolecule constructs a new molecule and adds it to the molecules map.
func (p *plant) buildAndAddMolecule(variant []rune, start, next int, molecule []rune) {
	var b strings.Builder
	for _, n := range molecule[:start] {
		b.WriteRune(n)
	}
	for _, n := range variant {
		b.WriteRune(n)
	}
	for _, n := range molecule[next:] {
		b.WriteRune(n)
	}
	p.molecules[b.String()]++
}

func part2(input string) int {
	// Parse the input into rules and molecule
	blocks, err := utils.ParseBlocksOfLines(input) // Assuming ParseBlocksOfLines parses replacement rules and molecule
	utils.HandleErr(err)                           // Handle errors
	rules := blocks[0]                             // First block is rules
	molecule := blocks[1][0]                       // Second block is the target molecule

	// Get parsed replacement rules (returns map[element][][]rune)
	replacementRules := replacements(rules)

	// Reverse the rules
	reversedRules := make(map[string]string)
	for from, toList := range replacementRules {
		// Convert the `from` key to string removing the 0 bytes and reverse it
		reversedFrom := reverseString(strings.Trim(string(from[:]), "\x00"))
		for _, to := range toList {
			// Reverse each `to` and store it as a reversed rule
			reversedTo := reverseRuneSlice(to)
			reversedRules[reversedTo] = reversedFrom
		}
	}

	// Reverse the molecule
	reversedMolecule := reverseString(molecule)

	// Build a regex for all reversed "From" replacements
	keys := make([]string, 0, len(reversedRules))
	for k := range reversedRules {
		keys = append(keys, k)
	}

	// Sort keys by length (longest replacements are matched first)
	sort.Slice(keys, func(i, j int) bool {
		return len(keys[i]) > len(keys[j])
	})

	// Combine into a single regex pattern
	regexReplacements := strings.Join(keys, "|")
	replacementRegex := regexp.MustCompile(regexReplacements)

	// Start processing the reversed molecule
	replacementCount := 0
	for reversedMolecule != "e" {
		// Replace one occurrence at a time using the regex
		newReversedMolecule := replacementRegex.ReplaceAllStringFunc(reversedMolecule, func(match string) string {
			replacementCount++ // Increment the count
			return reversedRules[match]
		})

		// Check if the molecule has changed (to prevent infinite loop)
		if newReversedMolecule == reversedMolecule {
			panic(fmt.Sprintf("No valid replacements found for molecule: %s", reversedMolecule))
		}

		reversedMolecule = newReversedMolecule
	}

	return replacementCount

}

// Utility function to reverse a string (for "from" and molecule)
func reverseString(input string) string {
	runes := []rune(input)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

// Utility function to reverse a slice of runes (for "to")
func reverseRuneSlice(input []rune) string {
	n := len(input)
	reversed := make([]rune, n)
	for i := 0; i < n; i++ {
		reversed[i] = input[n-1-i]
	}
	return string(reversed)
}
