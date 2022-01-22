package main

import (
	"fmt"
	"log"
	"net/http"

	database "github.com/afifialaa/book-directory-GO/database"
	"github.com/afifialaa/book-directory-GO/handlers"

	gorillaHandler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	database.Connect()
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")

	r.HandleFunc("/search/searchByAuthor", handlers.SearchByAuthor).Methods("GET")
	r.HandleFunc("/search/searchByTitle", handlers.SearchByTitle).Methods("GET")
	r.HandleFunc("/search/searchByID", handlers.SearchByID).Methods("GET")

	r.HandleFunc("/delete", handlers.DeleteByID).Methods("DELETE")

	r.HandleFunc("/update", handlers.UpdateBook).Methods("PUT")

	// Listening for requests
	fmt.Println("server is running")
	err := http.ListenAndServe(":8080", gorillaHandler.CORS()(r))
	if err != nil {
		log.Fatal("server crashed")
		log.Fatal(err)
	}
}
