package tests

import "testing"

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

func TestGetIngredientsMatchesSchema(t *testing.T) {
	resp := getIngredients(t, nil, nil)
	validateResponseJSON(t, resp, "../docs/schemas/ingredients/ingredients.json")
}

func TestGetIngredientsPagination(t *testing.T) {
	allResp := getIngredients(t, nil, nil)

	var allI ingredients
	decodeJSON(t, allResp, &allI)

	count := 1
	allISlice := allI.Ingredient
	for start := 0; start < len(allISlice); start++ {
		resp := getIngredients(t, &start, &count)

		var i ingredients
		decodeJSON(t, resp, &i)

		iSlice := i.Ingredient
		if len(iSlice) != 1 {
			t.Errorf("Returned count %d, expected 1", len(iSlice))
		}
		if iSlice[0] != allISlice[start] {
			t.Errorf("Offset %d does not start at correct entry", start)
		}
	}
}

func TestGetIngredient(t *testing.T) {
	resp := getIngredient(t, 1)
	validateResponseJSON(t, resp, "../docs/schemas/ingredients/ingredient.json")
}
