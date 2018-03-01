package main

import "fmt"

type pagination struct {
	Path  string
	Count int
	Start int
	Total int
}

type links struct {
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
}

// Generates the links for a paginated API
// using the path, count, and start of the current call
func (p *pagination) generatePaginationLinks() links {
	var l links

	previous := pagination{p.Path, p.Count, p.Start, p.Total}
	next := pagination{p.Path, p.Count, p.Start, p.Total}

	previous.Start = p.Start - p.Count
	next.Start = p.Start + p.Count

	if previous.Start < 0 {
		previous.Start = 0
	}

	if p.Start == 0 {
		l.Previous = nil
	} else {
		l.Previous = previous.generateURL()
	}

	if p.Start+p.Count >= p.Total {
		l.Next = nil
	} else {
		l.Next = next.generateURL()
	}

	return l
}

// Generates a pagination link for previous/next fields
func (p *pagination) generateURL() *string {
	link := fmt.Sprintf("%s?start=%d&count=%d", p.Path, p.Start, p.Count)
	return &link
}
