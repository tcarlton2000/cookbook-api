package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/xeipuuv/gojsonschema"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func readRequestBody(req *http.Request) ([]byte, error) {
	bodyBytes, err := ioutil.ReadAll(req.Body)
	return bodyBytes, err
}

func validateResponseJSON(bodyBytes []byte, schemaPath string) (bool, string, error) {
	schemaAbsPath, err := filepath.Abs(schemaPath)
	if err != nil {
		return false, "", err
	}

	schemaBytes, err := ioutil.ReadFile(schemaAbsPath)
	if err != nil {
		return false, "", err
	}

	schemaString := string(schemaBytes)
	schemaLoader := gojsonschema.NewStringLoader(schemaString)

	bodyString := string(bodyBytes)
	documentLoader := gojsonschema.NewStringLoader(bodyString)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return false, "", err
	}

	if !result.Valid() {
		errSlice := make([]string, 0)
		for _, err := range result.Errors() {
			errSlice = append(errSlice, err.String())
		}
		err := strings.Join(errSlice, ", ")
		return false, err, nil
	}

	return true, "", nil
}

// Ingredient Handles
func (a *app) getIngredient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ingredient ID")
		return
	}

	i := ingredient{ID: id}
	if err := i.getIngredient(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Ingredient not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, i)
}

func (a *app) createIngredient(w http.ResponseWriter, r *http.Request) {
	var i ingredient

	bodyBytes, err := readRequestBody(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	result, issues, err := validateResponseJSON(bodyBytes, "/docs/schemas/ingredients/ingredients-post.json")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	} else if !result {
		respondWithError(w, http.StatusBadRequest, issues)
		return
	}

	err = json.Unmarshal(bodyBytes, &i)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%q", err))
		return
	}
	defer r.Body.Close()

	if err := i.createIngredient(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, i)
}

func (a *app) getIngredients(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	ingredients, err := getIngredients(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, ingredients)
}

func (a *app) deleteIngredient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Ingredient ID")
		return
	}

	i := ingredient{ID: id}
	if err := i.deleteIngredient(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Ingredient not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// Recipe Handlers
func (a *app) getRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := getRecipes()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, recipes)
}
