package main

func getRecipes(start, count int) (recipes, error) {
	var recipeList []recipe
	var recipeLinks links
	recipeList = make([]recipe, 1)
	if start > 1 {
		recipeList[0] = recipe{1, "Stir Fried Buffallo"}
		prev := "/recipes?start=1"
		recipeLinks = links{&prev, nil}
	} else if start == 1 {
		recipeList[0] = recipe{0, "Oriental Pork Cabbage"}
		next := "/recipes?start=2"
		prev := "/recipes?start=0"
		recipeLinks = links{&prev, &next}
	} else {
		recipeList[0] = recipe{0, "Baked Chicken"}
		next := "/recipes?start=1"
		recipeLinks = links{nil, &next}
	}

	var ret recipes
	ret = recipes{recipeList, recipeLinks}

	return ret, nil
}
