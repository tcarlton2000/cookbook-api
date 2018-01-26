package tests

import (
	"bytes"
	"net/http"
	"testing"
)

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
	ing := createDefaultIngredient(t)
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

	deleteIngredient(t, ing.ID, 200)
}

func TestGetIngredient(t *testing.T) {
	ing := createDefaultIngredient(t)
	resp := getIngredient(t, ing.ID, 200)
	validateResponseJSON(t, resp, "../docs/schemas/ingredients/ingredient.json")

	deleteIngredient(t, ing.ID, 200)
}

func TestCreateIngredient(t *testing.T) {
	nut := nutrition{1, 2, 3, 4, 5}
	serving := servingSize{4, "grams"}
	i := ingredient{0, "name", "meat", serving, nut}
	resp := createIngredient(t, i, 201)
	i.ID = resp.ID

	if i != resp {
		t.Errorf("Expected ingredient '%v', found '%v'", i, resp)
	}

	deleteIngredient(t, resp.ID, 200)
}

func TestCreateIngredientInvalidType(t *testing.T) {
	nut := nutrition{1, 2, 3, 4, 5}
	serving := servingSize{4, "grams"}
	i := ingredient{0, "name", "invalid", serving, nut}
	resp := createIngredient(t, i, 400)

	deleteIngredient(t, resp.ID, 404)
}

func TestCreateIngredientMissingRequiredField(t *testing.T) {
	payload := `{
	    "name": "Name",
	    "servingSize": {
		"amount": 1,
		"unit": "tsp"
	    },
	    "nutrition": {
		"calories": 0,
		"carbs": 0,
		"protein": 0,
		"fat": 0,
		"cholestorol": 0
	    }
	}`

	url := getURL("/ingredients")

	var payloadBytes = []byte(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("%q", err)
		t.FailNow()
	}

	checkResponseCode(t, resp, 400)
}

func TestCreateDuplicateIngredient(t *testing.T) {
	nut := nutrition{1, 2, 3, 4, 5}
	serving := servingSize{4, "grams"}
	i := ingredient{0, "name", "meat", serving, nut}
	resp := createIngredient(t, i, 201)
	createIngredient(t, i, 400)

	deleteIngredient(t, resp.ID, 200)
}

func TestDeleteIngredient(t *testing.T) {
	nut := nutrition{1, 2, 3, 4, 5}
	serving := servingSize{4, "grams"}
	i := ingredient{0, "name", "meat", serving, nut}
	resp := createIngredient(t, i, 201)

	deleteIngredient(t, resp.ID, 200)
	getIngredient(t, resp.ID, 404)
}

func TestDeleteNonExistantIngredient(t *testing.T) {
	deleteIngredient(t, 9084059, 404)
}
