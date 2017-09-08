package main

type recipe struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type recipes struct {
	Recipes []recipe `json:"recipes"`
}

func getRecipes() (recipes, error) {
	var recipeList []recipe
	recipeList = make([]recipe, 2)
	recipeList[0] = recipe{0, "Baked Chicken"}
	recipeList[1] = recipe{1, "Stir Fried Buffallo"}

	var ret recipes
	ret = recipes{recipeList}

	return ret, nil
}
