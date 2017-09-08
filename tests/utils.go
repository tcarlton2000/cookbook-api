package tests

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
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
