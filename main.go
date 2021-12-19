package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO WORLD"))

}
func serveVideo(w http.ResponseWriter, r *http.Request) {

	phanloai := r.URL.Query().Get("phanloai")
	id := r.URL.Query().Get("id")
	w.Header().Add("Content-Type", "video/mp4") //dinh dang mp4 truoc
	path := fmt.Sprintf("video/%s/%s.mp4", phanloai, id)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		w.Header().Add("Content-Type", "text/html; charset=UTF-8")

		http.ServeFile(w, r, "static/a.html")

		return

	}

	http.ServeFile(w, r, path)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/abc", home)
	mux.HandleFunc("/video", serveVideo)

	log.Println("Server dang khoi dong...")
	err := http.ListenAndServe(":4000", mux) //khai bao cong
	log.Fatal(err)                           //xuat chu mau do
}
