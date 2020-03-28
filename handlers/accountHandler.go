package handlers

import (
	"github.com/afifialaa/auth"

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

	user := user_type{
		r.FormValue("firstName"),
		r.FormValue("lastName"),
		r.FormValue("email"),
		r.FormValue("password"),
	}

	// okay := validation.ValidateUser(user)
	okay := true
	if okay {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		//generate token
		token := auth.GenerateToken(user.email)

		//creating json
		data := map[string] string{"msg": "saved user", "token":token}

		//sending response
		json.NewEncoder(w).Encode(data)
	}else{
		w.Header().Set("Content-Type", "application/json")

		data := map[string] string{"msg" : "user was not created"}
		json.NewEncoder(w).Encode(data)
	}
}