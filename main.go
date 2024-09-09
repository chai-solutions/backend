package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, World!")
}

func main() {
	r := chi.NewRouter()
	fmt.Println("Starting")
	r.Get("/hello", helloHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
