package tests

import "testing"

func TestGetRecipes(t *testing.T) {
	resp := getRecipes(t)
	validateResponseJSON(t, resp, "../docs/schemas/recipes/recipes.json")
}

// Uncomment these when POST/DELETE recipe is implemented
// func TestGetRecipe(t *testing.T) {
// 	resp := getRecipe(t, 1)
// 	checkResponseCode(t, resp, 200)
// 	validateResponseJSON(t, resp, "../docs/schemas/recipes/recipe.json")
// }

// func TestGetRecipeNotFound(t *testing.T) {
// 	resp := getRecipe(t, 999)
// 	checkResponseCode(t, resp, 404)
// }
