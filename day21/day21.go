package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	input, _ := utils.ReadLines("./input.txt")
	db := initDB(input)
	defer db.Close()

	var ingredients []string
	var allergens []string
	db.Select(&ingredients, `SELECT DISTINCT ingredient FROM dish;`)
	db.Select(&allergens, `SELECT DISTINCT allergen FROM dish;`)

	part1Solution := part1(db, ingredients, allergens)
	part2Solution := part2(db)

	fmt.Printf("Day 21: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 21: Part 2: = %+v\n", part2Solution)
}

func part1(db *sqlx.DB, ingredients, allergens []string) int {
	for _, allergen := range allergens {
		var timesAllergenAppearsInAllLists int
		db.Get(&timesAllergenAppearsInAllLists, `SELECT COUNT(DISTINCT dish_id) FROM dish WHERE allergen = $1`, allergen)

		for _, ingredient := range ingredients {
			var timesIngredientIsAssociatedWithAllergen int
			db.Get(&timesIngredientIsAssociatedWithAllergen, `SELECT COUNT(*) FROM dish WHERE ingredient = $1 AND allergen = $2`, ingredient, allergen)
			if timesAllergenAppearsInAllLists == timesIngredientIsAssociatedWithAllergen {
				db.Exec(`INSERT INTO ingredient_allergen (ingredient, allergen) VALUES ($1, $2)`, ingredient, allergen)
			}
		}
	}

	var timesNonAllergensAppear int
	db.Get(&timesNonAllergensAppear, `
		SELECT COUNT(*) FROM (
			SELECT DISTINCT ingredient, dish_id
			FROM dish
			WHERE ingredient NOT IN (SELECT DISTINCT ingredient FROM ingredient_allergen)
		)
	`)
	return timesNonAllergensAppear
}

func part2(db *sqlx.DB) string {
	var allergens []string
	db.Select(&allergens, `SELECT DISTINCT allergen FROM ingredient_allergen`)

	// Start trimming down the possibilities in the ingredient_allergen table
	whittled := false
	for whittled != true {
		whittled = true
		for _, allergen := range allergens {
			res, _ := db.Exec(`	
				WITH valid_pair AS (
					SELECT ingredient, allergen FROM ingredient_allergen outer_ia
					WHERE outer_ia.ingredient NOT IN (
						SELECT ingredient
						FROM ingredient_allergen inner_ia
						WHERE inner_ia.ingredient = outer_ia.ingredient
						AND inner_ia.allergen != outer_ia.allergen
					)
					AND allergen = $1
				)
				DELETE
				FROM ingredient_allergen
				WHERE allergen = (SELECT allergen FROM valid_pair)
				AND ingredient != (SELECT ingredient FROM valid_pair);
		`, allergen)
			rowsAffected, _ := res.RowsAffected()
			whittled = whittled && rowsAffected == 0
		}
	}

	var ingredients []string
	db.Select(&ingredients, `SELECT ingredient FROM ingredient_allergen ORDER BY allergen ASC`)
	return strings.Join(ingredients, ",")
}

func initDB(input []string) *sqlx.DB {
	db, _ := sqlx.Open("sqlite3", "./day21.sqlite")
	db.MustExec(`
		DROP TABLE IF EXISTS dish; 
		DROP TABLE IF EXISTS ingredient_allergen; 
		CREATE TABLE dish (dish_id INTEGER, ingredient TEXT, allergen TEXT);
		CREATE TABLE ingredient_allergen (ingredient TEXT, allergen TEXT);

		CREATE INDEX ingredient_idx ON dish(ingredient);
		CREATE INDEX allergen_idx ON dish(allergen);
	`)

	tx, _ := db.Beginx()
	for listID, line := range input {
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")

		parts := strings.Split(line, "contains")
		ingredients := strings.Split(strings.TrimSpace(parts[0]), " ")
		allergens := strings.Split(strings.TrimSpace(parts[1]), ",")

		for _, ingredient := range ingredients {
			for _, allergen := range allergens {
				tx.Exec(`INSERT INTO dish (dish_id, ingredient, allergen) VALUES ($1, $2, $3)`, listID, strings.TrimSpace(ingredient), strings.TrimSpace(allergen))
			}
		}
	}
	tx.Commit()

	return db
}
