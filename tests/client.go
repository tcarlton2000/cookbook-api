package tests

import "net/http"
import "bytes"

func getURL(relativePath string) string {
	baseURL := "http://192.168.99.100:8080"
	var url bytes.Buffer

	url.WriteString(baseURL)
	url.WriteString(relativePath)

	return url.String()
}

func getRecipes() (*http.Response, error) {
	url := getURL("/recipes")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
