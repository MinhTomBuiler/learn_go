package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/learn_go_db/controllers"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/s3/{videoId}", controllers.GetS3Path)
	mux.HandleFunc("/mysql/{movieId}", controllers.GetMovieData)
	mux.HandleFunc("/mysql1/{userId}", controllers.GetMovieData1)
	mux.HandleFunc("/mysql2/{usermovieId}", controllers.GetMovieData2)

	log.Println("Server DB dang khoi dong...")
	err := http.ListenAndServe(":4001", mux) //khai bao cong
	log.Fatal(err)                           //xuat chu mau do
}
