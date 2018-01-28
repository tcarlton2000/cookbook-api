package main

import "database/sql"

func (i *detailedIngredient) getIngredient(db *sql.DB) error {
	return db.QueryRow("SELECT * FROM ingredients WHERE id=$1",
		i.ID).Scan(&i.ID, &i.Name, &i.Type, &i.ServingSize.Amount,
		&i.ServingSize.Unit, &i.Nutrition.Calories, &i.Nutrition.Carbs,
		&i.Nutrition.Protein, &i.Nutrition.Fat, &i.Nutrition.Cholestorol)
}

func (i *detailedIngredient) createIngredient(db *sql.DB) error {
	err := db.QueryRow(
		`INSERT INTO ingredients(name, type, serving_size, unit,
		calories, carbs, protein, fat, cholestorol) VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`,
		i.Name, i.Type, i.ServingSize.Amount, i.ServingSize.Unit,
		i.Nutrition.Calories, i.Nutrition.Carbs, i.Nutrition.Protein,
		i.Nutrition.Fat, i.Nutrition.Cholestorol).Scan(&i.ID)

	if err != nil {
		return err
	}

	return nil
}

func (i *detailedIngredient) deleteIngredient(db *sql.DB) error {
	getErr := db.QueryRow("SELECT * FROM ingredients WHERE id=$1",
		i.ID).Scan(&i.ID, &i.Name, &i.Type, &i.ServingSize.Amount,
		&i.ServingSize.Unit, &i.Nutrition.Calories, &i.Nutrition.Carbs,
		&i.Nutrition.Protein, &i.Nutrition.Fat, &i.Nutrition.Cholestorol)

	if getErr != nil {
		return getErr
	}

	_, err := db.Exec("DELETE FROM ingredients WHERE id=$1", i.ID)

	return err
}

func getIngredients(db *sql.DB, start, count int) (ingredients, error) {
	rows, err := db.Query(
		"SELECT id, name, type FROM ingredients LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return ingredients{}, err
	}

	defer rows.Close()

	ingredientList := []listIngredient{}

	for rows.Next() {
		var i listIngredient
		if err := rows.Scan(&i.ID, &i.Name, &i.Type); err != nil {
			return ingredients{}, err
		}
		ingredientList = append(ingredientList, i)
	}

	var ret ingredients
	ret = ingredients{ingredientList}

	return ret, nil
}
