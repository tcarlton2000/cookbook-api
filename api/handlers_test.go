package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testStruct struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Shape  string `json:"shape"`
}

func TestRespondWithJSON(t *testing.T) {
	w := httptest.NewRecorder()

	payload := testStruct{4, 4, "square"}
	respondWithJSON(w, 204, payload)

	if w.Code != 204 {
		t.Errorf("Expected code 204, found %d", w.Code)
	}

	payloadStr := `{"height":4,"width":4,"shape":"square"}`
	responseStr := w.Body.String()
	if responseStr != payloadStr {
		t.Errorf("Expected payload %s, found %s", payloadStr, responseStr)
	}

	contentType := w.HeaderMap["Content-Type"][0]
	if contentType != "application/json" {
		t.Errorf("Expected content type 'application/json', found '%s'", contentType)
	}
}

func TestResponseWithError(t *testing.T) {
	w := httptest.NewRecorder()

	message := "Could not form a valid shape"
	respondWithError(w, 404, message)

	if w.Code != 404 {
		t.Errorf("Expected code 404, found %d", w.Code)
	}

	payloadStr := `{"error":"Could not form a valid shape"}`
	responseStr := w.Body.String()
	if responseStr != payloadStr {
		t.Errorf("Expected payload %s, found %s", payloadStr, responseStr)
	}

	contentType := w.HeaderMap["Content-Type"][0]
	if contentType != "application/json" {
		t.Errorf("Expected content type 'application/json', found '%s'", contentType)
	}
}

func TestValidateResponseJSON(t *testing.T) {
	schemaLocation := "../docs/schemas/recipes/recipes.json"
	payload := `{
		"recipes": [
			{
				"id": 1,
				"name": "One"
			},
			{
				"id": 2,
				"name": "Two"
			}
		]
	}`

	payloadBytes := []byte(payload)
	result, issues, err := validateResponseJSON(payloadBytes, schemaLocation)

	if result != true {
		t.Errorf("Expected true result, found %t", result)
	}

	if issues != "" {
		t.Errorf("Expected no issues, found '%s'", issues)
	}

	if err != nil {
		t.Errorf("Expected no errors, found %q", err)
	}
}

func TestValidateResponseJSONInvalidJSON(t *testing.T) {
	schemaLocation := "../docs/schemas/recipes/recipes.json"
	payload := `{
		"recipes": [
			{
				"id": 1,
				"name": "One"
			},
			{
				"id": 2,
				"name": "Two"
			}
		]`

	payloadBytes := []byte(payload)
	result, issues, err := validateResponseJSON(payloadBytes, schemaLocation)

	if result != false {
		t.Errorf("Expected false result, found %t", result)
	}

	if issues != "" {
		t.Errorf("Expected no issues, found '%s'", issues)
	}

	if err.Error() != "unexpected EOF" {
		t.Errorf("Expected unexpected EOF, found %q", err)
	}
}

func TestValidateResponseJSONFailValidation(t *testing.T) {
	schemaLocation := "../docs/schemas/recipes/recipes.json"
	payload := `{
		"recipes": [
			{
				"id": "1",
				"name": "One"
			},
			{
				"id": 2,
				"name": "Two"
			}
		]
	}`

	payloadBytes := []byte(payload)
	result, issues, err := validateResponseJSON(payloadBytes, schemaLocation)

	if result != false {
		t.Errorf("Expected false result, found %t", result)
	}

	expectedIssue := "recipes.0.id: Invalid type. Expected: integer, given: string"
	if issues != expectedIssue {
		t.Errorf("Expected '%s', found '%s'", expectedIssue, issues)
	}

	if err != nil {
		t.Errorf("Expected no errors, found %q", err)
	}
}

func TestReadRequestBody(t *testing.T) {
	payload := "This is a test"
	payloadReader := strings.NewReader(payload)

	req, err := http.NewRequest("GET", "http://test.com", payloadReader)
	bodyBytes, err := readRequestBody(req)

	bodyString := string(bodyBytes[:])
	if bodyString != payload {
		t.Errorf("Expected body '%s', found '%s'", payload, bodyString)
	}

	if err != nil {
		t.Errorf("Expected no error, found '%q'", err)
	}
}
