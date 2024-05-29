package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
