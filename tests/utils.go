package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

func checkResponseCode(t *testing.T, resp *http.Response, expected int) bool {
	if resp.StatusCode != expected {
		bs, _ := ioutil.ReadAll(resp.Body)
		t.Errorf("Expected response code %d. Got %d\nResponse: %s", expected, resp.StatusCode, bs)
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return true
	}

	return false
}

func validateResponseJSON(t *testing.T, resp *http.Response, schemaPath string) {
	schemaAbsPath, err := filepath.Abs(schemaPath)
	if err != nil {
		t.Errorf("%q", err)
		return
	}

	schemaBytes, err := ioutil.ReadFile(schemaAbsPath)
	if err != nil {
		t.Errorf("%q", err)
		return
	}

	schemaString := string(schemaBytes)
	schemaLoader := gojsonschema.NewStringLoader(schemaString)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("%q", err)
		return
	}
	bodyString := string(bodyBytes)
	documentLoader := gojsonschema.NewStringLoader(bodyString)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		t.Errorf("%q", err)
		return
	}

	if !result.Valid() {
		t.Errorf("%q", result.Errors())
	}
}

func decodeJSON(t *testing.T, resp *http.Response, i interface{}) {
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&i); err != nil {
		t.Errorf("%q", err)
		t.FailNow()
	}
	defer resp.Body.Close()
}

func createDefaultIngredient(t *testing.T) detailedIngredient {
	nut := nutrition{1, 2, 3, 4, 5}
	serving := servingSize{4, "grams"}
	i := detailedIngredient{0, "name", "meat", serving, nut}
	resp := createIngredient(t, i, 201)

	return resp
}
