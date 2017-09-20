package tests

import "testing"

func TestGetRecipes(t *testing.T) {
	resp := getRecipes(t)
	validateResponseJSON(t, resp, "../docs/schemas/recipes/recipes.json")
}
