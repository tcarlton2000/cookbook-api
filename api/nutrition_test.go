package main

import "testing"

func TestTruncateNutrition(t *testing.T) {
	n := nutrition{1.234, 2.345, 3.456, 4.567, 5.678}
	expected := nutrition{1.23, 2.35, 3.46, 4.57, 5.68}
	n.truncate()

	if n != expected {
		t.Errorf("Expected %v, found %v", expected, n)
	}
}

func TestAppendNutrition(t *testing.T) {
	total := nutrition{1, 2, 3, 4, 5}
	add := nutrition{2, 4, 6, 8, 10}
	expected := nutrition{3, 6, 9, 12, 15}

	total.append(add)

	if total != expected {
		t.Errorf("Expected %v, found %v", expected, total)
	}
}
