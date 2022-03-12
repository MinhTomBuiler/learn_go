package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/learn_go_db/models"
)

func createMySqlConnection() *sql.DB {
	connectionStr := "root:example@/flixtube"
	db, err := sql.Open("mysql", connectionStr)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		db.Close()
		log.Fatalf("Can not connect to database %s: %v", connectionStr, err)
		panic(err)
	}
	return db
}

func GetMovieData(w http.ResponseWriter, r *http.Request) {
	db := createMySqlConnection()
	movieModel := models.Moviemodels{db}
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	movieIdInt, err := strconv.Atoi(movieId)
	if err != nil {
		log.Fatal("Error convert param to number")
		panic(err)

	}
	movieObj, err := movieModel.SelectById(movieIdInt)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Fprintf(w, "%v", movieObj)
}
func GetMovieData1(w http.ResponseWriter, r *http.Request) {
	db := createMySqlConnection()
	userModel := models.Usermodels{db}
	vars := mux.Vars(r)
	userId := vars["userId"]
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		log.Fatal("Error convert param to number")
		panic(err)

	}
	userObj, err := userModel.SelectById(userIdInt)
	log.Println(userObj)
	if err != nil {
		log.Fatal("Error select movie model")
		panic(err)
	}
	fmt.Fprintf(w, "%v", userObj)
}
func GetMovieData2(w http.ResponseWriter, r *http.Request) {
	db := createMySqlConnection()
	usermovieModel := models.Usermoviemodels{db}
	vars := mux.Vars(r)
	usermovieId := vars["usermovieId"]
	usermovieIdInt, err := strconv.Atoi(usermovieId)
	if err != nil {
		log.Fatal("Error convert param to number")
		panic(err)

	}
	usermovieObj, err := usermovieModel.SelectById(usermovieIdInt)
	log.Println(usermovieObj)
	if err != nil {
		log.Fatal("Error select movie model")
		panic(err)
	}
	fmt.Fprintf(w, "%v", usermovieObj)
}
