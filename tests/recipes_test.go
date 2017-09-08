package tests

import "testing"

func TestGetRecipes(t *testing.T) {
	resp, err := getRecipes()
	if err != nil {
		t.Errorf("%q", err)
		t.FailNow()
	}
	checkResponseCode(t, 200, resp.StatusCode)
	validateResponseJSON(t, resp, "../docs/schemas/recipes/recipes.json")
}
