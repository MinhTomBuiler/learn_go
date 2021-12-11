package main

import (
    "net/http"
    "log"
    "fmt"
)

func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("HELLO WORLD"))


}
func serveVideo(w http.ResponseWriter, r *http.Request){

    phanloai := r.URL.Query().Get("phanloai")
    id := r.URL.Query().Get("id")
    w.Header().Add("Content-Type","video/mp4")//dinh dang mp4 truoc
    path := fmt.Sprintf("video/%s/%s.mp4", phanloai, id)
    http.ServeFile(w,r, path)
}

func main () {
    mux := http.NewServeMux()
    mux.HandleFunc("/abc",home)
    mux.HandleFunc("/video",serveVideo)

    log.Println("Server dang khoi dong...")
    err := http.ListenAndServe(":4000", mux)//khai bao cong
    log.Fatal(err)//xuat chu mau do
}

    