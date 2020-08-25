package main

import (
	"fmt"
	"log"
	"net/http"
)

func putInfoLog(r *http.Request) {
	log.Printf("Info: Request from %v", r.RequestURI)
}

func root(w http.ResponseWriter, r *http.Request) {
	putInfoLog(r)
	fmt.Fprintf(w, "Hi, ecs test.")
}

func hoge(w http.ResponseWriter, r *http.Request) {
	putInfoLog(r)
	fmt.Fprintf(w, "Hoge.")
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/hoge", hoge)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
