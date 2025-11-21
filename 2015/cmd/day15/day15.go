package main

import (
	"fmt"
	"github.com/akolybelnikov/aoc-2015/internal/utils"
	"os"
	"strconv"
	"strings"
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
	amount     int
}

type recipe struct {
	ingredients []*ingredient
	capacity    int
	durability  int
	flavor      int
	texture     int
	calories    int
}

func (r *recipe) addIngredient(ingredient *ingredient) {
	r.ingredients = append(r.ingredients, ingredient)
}

func (r *recipe) calculate() {
	for _, i := range r.ingredients {
		r.capacity += i.capacity * i.amount
		r.durability += i.durability * i.amount
		r.flavor += i.flavor * i.amount
		r.texture += i.texture * i.amount
		r.calories += i.calories * i.amount
	}
	if r.capacity < 0 {
		r.capacity = 0
	}
	if r.durability < 0 {
		r.durability = 0
	}
	if r.flavor < 0 {
		r.flavor = 0
	}
	if r.texture < 0 {
		r.texture = 0
	}
}

func (r *recipe) score() int {
	return r.capacity * r.durability * r.flavor * r.texture
}

func main() {
	data, err := os.ReadFile("inputs/day15.txt")
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
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)
	ingredients := make([]*ingredient, 0)
	for _, line := range lines {
		ingredients = append(ingredients, makeIngredient(line))
	}
	recipes := findBestRecipe(ingredients, 100, false)

	return recipes[0].score()
}

// part two
func part2(input string) int {
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)
	ingredients := make([]*ingredient, 0)
	for _, line := range lines {
		ingredients = append(ingredients, makeIngredient(line))
	}
	recipes := findBestRecipe(ingredients, 100, true)

	return recipes[0].score()
}

func makeIngredient(input string) *ingredient {
	parts := strings.Split(input, ": ")
	name := parts[0]
	properties := parts[1]
	props := strings.Split(properties, ", ")
	ingredientFields := make(map[string]int)
	for _, prop := range props {
		keyValue := strings.Split(prop, " ")
		key := keyValue[0]
		value, _ := strconv.Atoi(keyValue[1])
		ingredientFields[key] = value
	}

	// Create the ingredient object
	return &ingredient{
		name:       name,
		capacity:   ingredientFields["capacity"],
		durability: ingredientFields["durability"],
		flavor:     ingredientFields["flavor"],
		texture:    ingredientFields["texture"],
		calories:   ingredientFields["calories"],
	}
}

func findBestRecipe(items []*ingredient, n int, cal bool) []*recipe {
	var optimalRecipes []*recipe
	maxScore := 0
	count := 0

	var helper func(target, itemsLeft int, cur *recipe)
	helper = func(target, itemsLeft int, cur *recipe) {
		if itemsLeft == 1 {
			lastIngredient := items[len(cur.ingredients)]
			lastIngredient.amount = target
			cur.addIngredient(lastIngredient)
			cur.calculate()
			score := cur.score()
			if score > maxScore {
				if !cal {
					maxScore = score
					optimalRecipes = []*recipe{cur}
				} else {
					if cur.calories == 500 {
						maxScore = score
						optimalRecipes = []*recipe{cur}
					}
				}
			} else if score == maxScore {
				if !cal {
					optimalRecipes = append(optimalRecipes, cur)
				} else {
					if cur.calories == 500 {
						optimalRecipes = append(optimalRecipes, cur)
					}
				}
			}
			count++
			return
		}

		for i := 0; i <= target; i++ {
			newRecipe := &recipe{
				ingredients: append([]*ingredient{}, cur.ingredients...),
			}
			curIngredient := items[len(newRecipe.ingredients)]
			curIngredient.amount = i
			newRecipe.addIngredient(curIngredient)
			helper(target-i, itemsLeft-1, newRecipe)
		}
	}

	helper(n, len(items), &recipe{})

	fmt.Printf("Found %d total recipes.\n", count)

	return optimalRecipes
}
