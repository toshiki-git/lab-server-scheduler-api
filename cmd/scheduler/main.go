package main

import (
	"fmt"
	"net/http"
)

func hoge(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hoge")
}

func fuga(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "fuga")
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	http.HandleFunc("/hoge", hoge)
	http.HandleFunc("/fuga", fuga)

	server.ListenAndServe()
}