package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func getURL(relativePath string) string {
	baseURL := "http://192.168.99.100:8080"
	var url bytes.Buffer

	url.WriteString(baseURL)
	url.WriteString(relativePath)

	return url.String()
}

func getRecipes(t *testing.T) *http.Response {
	url := getURL("/recipes")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Errorf("%q", err)
		t.FailNow()
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("%q", err)
		t.FailNow()
	}

	checkResponseCode(t, 200, resp.StatusCode)

	return resp
}

func getIngredients(t *testing.T, start *int, count *int) *http.Response {
	var urlString string
	if start != nil && count != nil {
		urlString = fmt.Sprintf("/ingredients?start=%d&count=%d", *start, *count)
	} else if start != nil {
		urlString = fmt.Sprintf("/ingredients?start=%d", *start)
	} else if count != nil {
		urlString = fmt.Sprintf("/ingredients?count=%d", *count)
	} else {
		urlString = "/ingredients"
	}
	url := getURL(urlString)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Errorf("%q", err)
		t.FailNow()
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("%q", err)
		t.FailNow()
	}

	checkResponseCode(t, 200, resp.StatusCode)

	return resp
}

func getIngredient(t *testing.T, id int) *http.Response {
	urlString := fmt.Sprintf("/ingredients/%d", id)
	url := getURL(urlString)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Errorf("%q", err)
		t.FailNow()
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("%q", err)
		t.FailNow()
	}

	checkResponseCode(t, 200, resp.StatusCode)

	return resp
}
