package handlers

import (
	"github.com/afifialaa/auth"
	"github.com/afifialaa/validation"

	"net/http"
	"fmt"
	"encoding/json"
)

type user_type struct{
	firstName string
	lastName string
	email string
	password string
}

//signin handle
func LoginHandle(w http.ResponseWriter, r *http.Request){

	mockUser := struct{
		email string
		password string
	}{
		"afifi@gmail.com",
		"afifi123",
	}

	fmt.Println("mock user email " + mockUser.email)

	user := struct{
		email string
		password string
	}{
		r.FormValue("email"),
		r.FormValue("password"),
	}

	fmt.Println("user email " + user.email)

	//aquiring token from request header
	reqToken := auth.GetToken(r)
	auth.ValidateToken(reqToken)
}

func SignupHandle(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "*")

	user := validation.User_type{
		r.FormValue("firstName"),
		r.FormValue("lastName"),
		r.FormValue("email"),
		r.FormValue("password"),
	}

		okay := validation.ValidateUser(&user)
	if okay {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		// Generate token
		token := auth.GenerateToken(user.Email)

		// Generate json
		data := map[string] string{"msg": "saved user", "token":token}

		// Sending response
		json.NewEncoder(w).Encode(data)
	}else{
		w.Header().Set("Content-Type", "application/json")

		data := map[string] string{"msg" : "user was not created"}
		json.NewEncoder(w).Encode(data)
	}
}
