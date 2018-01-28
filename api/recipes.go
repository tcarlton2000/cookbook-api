package main

func getRecipes() (recipes, error) {
	var recipeList []recipe
	recipeList = make([]recipe, 2)
	recipeList[0] = recipe{0, "Baked Chicken"}
	recipeList[1] = recipe{1, "Stir Fried Buffallo"}

	var ret recipes
	ret = recipes{recipeList}

	return ret, nil
}
