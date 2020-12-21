package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/neilo40/adventofcode2020/internal/common"
)

func main() {
	re := regexp.MustCompile(`(.*) \(contains (.*)\)`)
	lines := common.ReadFileString("day21.input")
	allergenFoodCount := make(map[string]int)                  // how many foods contain each allergen
	ingredientAllergenCount := make(map[string]map[string]int) // how many of each allergen could be in each ingredient
	ingredientSeenCount := make(map[string]int)                // how many times was each ingredient seen in a food
	for _, l := range lines {
		parts := re.FindStringSubmatch(l)
		allergens := strings.Fields(strings.ReplaceAll(parts[2], ",", ""))
		ingredients := strings.Fields(parts[1])
		for _, i := range ingredients {
			_, ok := ingredientSeenCount[i]
			if !ok {
				ingredientSeenCount[i] = 1
			} else {
				ingredientSeenCount[i]++
			}
		}
		for _, a := range allergens {
			_, ok := allergenFoodCount[a]
			if !ok {
				allergenFoodCount[a] = 1
			} else {
				allergenFoodCount[a]++
			}
			for _, i := range ingredients {
				_, ok := ingredientAllergenCount[i]
				if !ok {
					ingredientAllergenCount[i] = map[string]int{a: 1}
				} else {
					_, ok = ingredientAllergenCount[i][a]
					if !ok {
						ingredientAllergenCount[i][a] = 1
					} else {
						ingredientAllergenCount[i][a]++
					}
				}
			}
		}
	}

	// get ingredients that are potentially allergic (seen as many times as an allergen)
	potentiallyAllergic := make(map[string]bool)
	for a, count := range allergenFoodCount {
		for i, aCount := range ingredientAllergenCount {
			if count == aCount[a] {
				potentiallyAllergic[i] = true
			}
		}
	}

	// get ingredients that can't possibly be allergic (all ingredients - potentially allergenic ones)
	safeIngredients := make([]string, 0, len(ingredientAllergenCount))
	for i := range ingredientAllergenCount {
		_, allergic := potentiallyAllergic[i]
		if !allergic {
			safeIngredients = append(safeIngredients, i)
		}
	}

	// how many times was each of these seen?
	seenCount := 0
	for _, i := range safeIngredients {
		seenCount += ingredientSeenCount[i]
	}

	fmt.Printf("Part 1 result: %d\n", seenCount)

	// remove the safe ingredients from the list
	for _, i := range safeIngredients {
		delete(ingredientAllergenCount, i)
	}

	definitelyAllergic := make(map[string]string)
	for a, count := range allergenFoodCount {
		for i, aCount := range ingredientAllergenCount {
			if count == aCount[a] {
				fmt.Printf("%s contains %s\n", i, a)
				definitelyAllergic[i] = a
			}
		}
	}

}
