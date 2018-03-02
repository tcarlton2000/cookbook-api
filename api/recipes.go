package main

import (
	"database/sql"
	"encoding/json"

	"github.com/martinlindhe/unit"
)

func getRecipes(db *sql.DB, start, count int) (recipes, error) {
	rows, err := db.Query(
		"SELECT id, name , COUNT(*) OVER() FROM recipes LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return recipes{}, err
	}

	defer rows.Close()

	recipeList := []recipe{}

	var total int
	for rows.Next() {
		var r recipe
		if err := rows.Scan(&r.ID, &r.Name, &total); err != nil {
			return recipes{}, err
		}
		recipeList = append(recipeList, r)
	}

	var ret recipes
	ret = recipes{recipeList, links{}}

	p := pagination{"/recipes", count, start, total}
	ret.Links = p.generatePaginationLinks()

	return ret, nil
}

func (r *detailedRecipe) getRecipe(db *sql.DB) error {
	var steps []byte
	err := db.QueryRow("SELECT * FROM recipes WHERE id=$1",
		r.ID).Scan(&r.ID, &r.Name, &steps)
	if err != nil {
		return err
	}

	err = json.Unmarshal(steps, &r.Steps)
	if err != nil {
		return err
	}

	var rows *sql.Rows
	rows, err = db.Query(
		`SELECT ri.ingredient_id, ri.amount, ri.unit, i.name,
		i.serving_size, i.unit, i.calories, i.carbs, i.protein,
		i.fat, i.cholestorol FROM recipe_ingredients AS ri
		INNER JOIN ingredients AS i ON (i.id = ri.ingredient_id)
		WHERE ri.recipe_id = $1`, &r.ID)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var ri recipeIngredient
		var i detailedIngredient
		err = rows.Scan(&ri.ID, &ri.Amount, &ri.Unit, &ri.Name,
			&i.ServingSize.Amount, &i.ServingSize.Unit,
			&i.Nutrition.Calories, &i.Nutrition.Carbs,
			&i.Nutrition.Protein, &i.Nutrition.Fat,
			&i.Nutrition.Cholestorol)

		if err != nil {
			return err
		}

		r.Ingredients = append(r.Ingredients, ri)

		n := calculateServingNutrition(i, ri)
		r.Nutrition.append(n)
	}

	r.Nutrition.truncate()

	return nil
}

func calculateServingNutrition(i detailedIngredient, ri recipeIngredient) nutrition {
	var unitMap = map[string]unit.Unit{
		"tbsp":      unit.Unit(unit.USTableSpoon),
		"tsp":       unit.Unit(unit.USTeaSpoon),
		"cups":      unit.Unit(unit.USCup),
		"fl ounces": unit.Unit(unit.USFluidOunce),
		"ounces":    unit.Unit(unit.AvoirdupoisOunce),
		"pounds":    unit.Unit(unit.AvoirdupoisPound),
		"grams":     unit.Unit(unit.Gram),
	}

	type convFunc func() float64

	servingAmount := unit.Unit(i.ServingSize.Amount) * unitMap[i.ServingSize.Unit]
	var conversionFuctions = map[string]convFunc{
		"tbsp":      unit.Volume(servingAmount).USTableSpoons,
		"tsp":       unit.Volume(servingAmount).USTeaSpoons,
		"cups":      unit.Volume(servingAmount).USCups,
		"fl ounces": unit.Volume(servingAmount).USFluidOunces,
		"ounces":    unit.Mass(servingAmount).AvoirdupoisOunces,
		"pounds":    unit.Mass(servingAmount).AvoirdupoisPounds,
		"grams":     unit.Mass(servingAmount).Grams,
	}

	weight := ri.Amount / conversionFuctions[ri.Unit]()

	var n nutrition

	n.Calories = weight * i.Nutrition.Calories
	n.Carbs = weight * i.Nutrition.Carbs
	n.Protein = weight * i.Nutrition.Protein
	n.Fat = weight * i.Nutrition.Fat
	n.Cholestorol = weight * i.Nutrition.Cholestorol

	return n
}
