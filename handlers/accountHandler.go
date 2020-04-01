package handlers

import (
	"github.com/afifialaa/auth"
	"github.com/afifialaa/validation"
	database "github.com/afifialaa/database"

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

// Login handle
func LoginHandle(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "*")

	user := validation.User_login_type{
		r.FormValue("email"),
		r.FormValue("password"),
	}

	// Aquiring token from request header
	reqToken := auth.GetToken(r)

	// No token provided
	if reqToken == "" {

		fmt.Println("token not found")

		// Send failed response 
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		// Generate json
		data := map[string] string{"msg": "token missing"}

		// Sending response
		json.NewEncoder(w).Encode(data)
	}else{
		userFound := database.FindUser(&user)
		if userFound {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			// Generate json
			data := map[string] string{"msg": "user found"}

			// Sending response
			json.NewEncoder(w).Encode(data)
		}else{
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			// Generate json
			data := map[string] string{"msg": "user was not found"}

			// Sending response
			json.NewEncoder(w).Encode(data)
		}

	}
}

func SignupHandle(w http.ResponseWriter, r *http.Request){
	fmt.Println("#Signup_handle")
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

		saved := database.SaveUser(&user)
		if !saved {
			fmt.Println("#signupHandle user was not saved")
			w.Header().Set("Content-Type", "application/json")

			data := map[string] string{"msg" : "user was not created"}
			json.NewEncoder(w).Encode(data)
		}

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
