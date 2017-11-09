package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

func checkResponseCode(t *testing.T, resp *http.Response, actual int) bool {
	if resp.StatusCode != actual {
		bs, _ := ioutil.ReadAll(resp.Body)
		t.Errorf("Expected response code %d. Got %d\nResponse: %s", resp.StatusCode, actual, bs)
	}

	if actual >= 200 && actual < 400 {
		return true
	} else {
		return false
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

func decodeJSON(t *testing.T, resp *http.Response, i interface{}) {
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&i); err != nil {
		t.Errorf("%q", err)
		t.FailNow()
	}
	defer resp.Body.Close()
}
