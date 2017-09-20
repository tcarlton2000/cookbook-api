package main

import "database/sql"

type servingSize struct {
	Amount float32 `json:"amount"`
	Unit   string  `json:"unit"`
}

type nutrition struct {
	Calories    float32 `json:"calories"`
	Carbs       float32 `json:"carbs"`
	Protein     float32 `json:"protein"`
	Fat         float32 `json:"fat"`
	Cholestorol float32 `json:"cholestorol"`
}

type ingredient struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	ServingSize servingSize `json:"servingSize"`
	Nutrition   nutrition   `json:"nutrition"`
}

type ingredientsIngredient struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type ingredients struct {
	Ingredient []ingredientsIngredient `json:"ingredients"`
}

func (i *ingredient) getIngredient(db *sql.DB) error {
	return db.QueryRow("SELECT * FROM ingredients WHERE id=$1",
		i.ID).Scan(&i.ID, &i.Name, &i.Type, &i.ServingSize.Amount,
		&i.ServingSize.Unit, &i.Nutrition.Calories, &i.Nutrition.Carbs,
		&i.Nutrition.Protein, &i.Nutrition.Fat, &i.Nutrition.Cholestorol)
}

func getIngredients(db *sql.DB, start, count int) (ingredients, error) {
	rows, err := db.Query(
		"SELECT id, name, type FROM ingredients LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return ingredients{}, err
	}

	defer rows.Close()

	ingredientList := []ingredientsIngredient{}

	for rows.Next() {
		var i ingredientsIngredient
		if err := rows.Scan(&i.ID, &i.Name, &i.Type); err != nil {
			return ingredients{}, err
		}
		ingredientList = append(ingredientList, i)
	}

	var ret ingredients
	ret = ingredients{ingredientList}

	return ret, nil
}
