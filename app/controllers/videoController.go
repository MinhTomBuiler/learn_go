package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO WORLD"))

}

func ServeVideo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	phanloai := vars["phanloai"]
	id := vars["id"]
	w.Header().Add("Content-Type", "video/mp4") //dinh dang mp4 truoc
	path := fmt.Sprintf("video/%s/%s.mp4", phanloai, id)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		w.Header().Add("Content-Type", "text/html; charset=UTF-8")

		http.ServeFile(w, r, "static/a.html")

		return

	}

	http.ServeFile(w, r, path)
}
