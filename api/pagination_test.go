package main

import "testing"

func stringToPointer(s string) *string {
	return &s
}
func TestGeneratePaginationLinks(t *testing.T) {
	type paginationTest struct {
		P             pagination
		ExpectedLinks links
	}

	var tests = []paginationTest{
		{
			pagination{"/recipes", 1, 0, 2},
			links{nil, stringToPointer("/recipes?start=1&count=1")},
		},
		{
			pagination{"/recipes", 1, 1, 2},
			links{stringToPointer("/recipes?start=0&count=1"), nil},
		},
		{
			pagination{"/recipes", 1, 1, 3},
			links{stringToPointer("/recipes?start=0&count=1"), stringToPointer("/recipes?start=2&count=1")},
		},
		{
			pagination{"/recipes", 10, 1, 3},
			links{stringToPointer("/recipes?start=0&count=10"), nil},
		},
		{
			pagination{"/recipes", 10, 1, 30},
			links{stringToPointer("/recipes?start=0&count=10"), stringToPointer("/recipes?start=11&count=10")},
		},
	}

	for i, tt := range tests {
		l := tt.P.generatePaginationLinks()

		if l.Previous != tt.ExpectedLinks.Previous {
			if l.Previous == nil && tt.ExpectedLinks.Previous != nil {
				t.Errorf("For case %d, Expected Previous %v, found nil", i, *tt.ExpectedLinks.Previous)
			} else if l.Previous != nil && tt.ExpectedLinks.Previous == nil {
				t.Errorf("For case %d, Expected Previous nil, found %v", i, *l.Previous)
			} else if *l.Previous != *tt.ExpectedLinks.Previous {
				t.Errorf("For case %d, Expected Previous %v, found %v", i, *tt.ExpectedLinks.Previous, *l.Previous)
			}
		}

		if l.Next != tt.ExpectedLinks.Next {
			if l.Next == nil && tt.ExpectedLinks.Next != nil {
				t.Errorf("For case %d, Expected Next %v, found nil", i, *tt.ExpectedLinks.Next)
			} else if l.Next != nil && tt.ExpectedLinks.Next == nil {
				t.Errorf("For case %d, Expected Next nil, found %v", i, *l.Next)
			} else if *l.Next != *tt.ExpectedLinks.Next {
				t.Errorf("For case %d, Expected Next %v, found %v", i, *tt.ExpectedLinks.Next, *l.Next)
			}
		}
	}
}
