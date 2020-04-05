package main

import (
	"github.com/afifialaa/handlers"
	database "github.com/afifialaa/database"
	"fmt"
	"net/http"

)

type Status struct{
	msg string
}

func main(){
	database.Connect()

	// routes
	http.HandleFunc("/user/signup", handlers.SignupHandle)
	http.HandleFunc("/user/signin", handlers.LoginHandle)
	http.HandleFunc("/api/service/test", handlers.TestHandle)

	// listening for requests
	fmt.Println("server is running")
	http.ListenAndServe(":8080", nil)

}
