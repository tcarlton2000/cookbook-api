package main

import "testing"

func TestCalculateServingNutrition(t *testing.T) {
	var i detailedIngredient
	var ri recipeIngredient

	i.Nutrition = nutrition{1, 2, 3, 4, 5}
	i.ServingSize.Amount = 8
	i.ServingSize.Unit = "ounces"

	ri.Amount = 1.5
	ri.Unit = "pounds"

	expected := nutrition{3, 6, 9, 12, 15}
	actual := calculateServingNutrition(i, ri)
	if expected != actual {
		t.Errorf("Expected %v, found %v", expected, actual)
	}
}

func TestTruncatePrecision(t *testing.T) {
	untruncated := 1.245
	expected := 1.25
	truncated := truncatePrecision(untruncated)

	if expected != truncated {
		t.Errorf("Expected %f, found %f", expected, truncated)
	}
}
