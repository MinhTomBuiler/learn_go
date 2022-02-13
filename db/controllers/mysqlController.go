package controllers

import (
	"database/sql"
	"net/http"
)

func createMySqlConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:example@127.0.0.1/flixtube")
	if err != nil {
		panic(err)
	}
	return db
}

func GetUserMovieData(w http.ResponseWriter, r *http.Request) {
	createMySqlConnection()
}
