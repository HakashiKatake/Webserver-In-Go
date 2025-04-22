package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Handler code for listing books
		json.NewEncoder(w).Encode("Hello world")
	})

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", router)
}
