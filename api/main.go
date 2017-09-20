package main

import "os"

func main() {
	a := app{}
	a.initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_NAME"))

	a.run(":8080")
}
