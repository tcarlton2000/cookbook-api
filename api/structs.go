package main

// Common Structs
type links struct {
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
}

// Ingredient Structs
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

type detailedIngredient struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	ServingSize servingSize `json:"servingSize"`
	Nutrition   nutrition   `json:"nutrition"`
}

type listIngredient struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type recipeIngredient struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Unit   string `json:"unit"`
}

type ingredients struct {
	Ingredient []listIngredient `json:"ingredients"`
}

// Recipe Structs
type recipe struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type detailedRecipe struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Ingredients []recipeIngredient `json:"ingredients"`
	Steps       []string           `json:"steps"`
	Nutrition   nutrition          `json:"nutrition"`
}

type recipes struct {
	Recipes []recipe `json:"recipes"`
	Links   links    `json:"links"`
}
