package main

// Ingredient Structs
type servingSize struct {
	Amount float64 `json:"amount"`
	Unit   string  `json:"unit"`
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
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Unit   string  `json:"unit"`
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
