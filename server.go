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
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	apiv1 := api.PathPrefix("/v1").Subrouter()

	// Routes
	apiv1.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	apiv1.HandleFunc("/book/{author}", handlers.SearchByAuthor).Methods("GET")
	apiv1.HandleFunc("/book/{title}", handlers.SearchByTitle).Methods("GET")

	apiv1.HandleFunc("/book/{bookId}", handlers.DeleteByID).Methods("DELETE")
	apiv1.HandleFunc("/book/{bookId}", handlers.SearchByID).Methods("GET")
	apiv1.HandleFunc("/book/{bookId}", handlers.UpdateBook).Methods("PUT")

	// Listening for requests
	fmt.Println("server is running")
	err := http.ListenAndServe(":8080", gorillaHandler.CORS()(r))
	if err != nil {
		log.Fatal("server crashed")
		log.Fatal(err)
	}
}
