package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type app struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *app) initialize(user, password, host, dbname string) {
	connectionString :=
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter().StrictSlash(true)
	a.initializeRoutes()
}

func (a *app) run(addr string) {
	// open a file
	logfile, err := os.OpenFile("/var/log/cookbook_api.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	log.SetOutput(logfile)
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}

func (a *app) initializeRoutes() {
	a.Router.HandleFunc("/recipes", a.getRecipes).Methods("GET")
	a.Router.HandleFunc("/ingredients", a.getIngredients).Methods("GET")
	a.Router.HandleFunc("/ingredients", a.createIngredient).Methods("POST")
	a.Router.HandleFunc("/ingredients/{id:[0-9]+}", a.getIngredient).Methods("GET")
	a.Router.HandleFunc("/ingredients/{id:[0-9]+}", a.deleteIngredient).Methods("DELETE")
}
