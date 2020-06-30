package main

import (
	"fmt"
	"net/http"

	database "github.com/afifialaa/REST-GO/database"
	"github.com/afifialaa/REST-GO/handlers"

	gorillaHandler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Status struct {
	msg string
}

func main() {
	database.Connect()
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")

	r.HandleFunc("/search/searchByAuthor", handlers.SearchByAuthor).Methods("GET")
	r.HandleFunc("/search/searchByTitle", handlers.SearchByTitle).Methods("GET")
	r.HandleFunc("/search/searchByID", handlers.SearchByID).Methods("GET")
	r.HandleFunc("/search/test", handlers.TestHandler).Methods("GET")

	// Listening for requests
	fmt.Println("server is running")
	http.ListenAndServe(":8080", gorillaHandler.CORS()(r))
}
