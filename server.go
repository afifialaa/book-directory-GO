package main

import (
	"fmt"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"encoding/json"
	"regexp"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

type User struct{
	firstName string
	lastName string
	email string
	password string
}

type Status struct{
	msg string
}

func main(){
	//routes
	http.HandleFunc("/user/signup", signupHandle)
	http.HandleFunc("/user/signin", signinHandle)
	http.HandleFunc("/user/test", signinHandle)

	//listening for requests
	fmt.Println("server is running")
	http.ListenAndServe(":8080", nil)

}

//signin handle
func signinHandle(w http.ResponseWriter, r *http.Request){

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
	reqToken := getToken(r)
	validateToken(reqToken)
}



//validate token
func validateToken(tokenString string){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if err != nil {
			return err
		}else{
			return token
		}
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("from token")
		fmt.Println(claims["email"])
	} else {
		fmt.Println(err)
	}
}

//get token from request
func getToken(r *http.Request)string{

	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	if reqToken != nil {
		return reqToken
	} 

	return "no token"
}


//signup handle
func signupHandle(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "*")

	user := User{
		r.FormValue("firstName"),
		r.FormValue("lastName"),
		r.FormValue("email"),
		r.FormValue("password"),
	}

	okay := validateUser(&user)
	if okay {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		//generate token
		token := generateToken(user.email)

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

func validateUser(user *User) bool{

	//empty fields
	if user.firstName == "" || user.lastName == "" || user.email == "" || user.password == ""{
		fmt.Println("empyty fields")
		return false
	}

	//empty field
	if len(user.firstName) == 0 || len(user.lastName) == 0 || len(user.email) == 0 || len(user.password) == 0{
		fmt.Println("empty fields")

		return false
	}

	//validating email
	if !validateEmail(user.email){
		fmt.Println("not a valid email")
		return false
	}

	fmt.Println("email: " + user.email + " " , validateEmail(user.email))

	user.password = hashPassword(user.password)
	return true
}

//validating email
func validateEmail(email string) bool{
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)
}

//hashing password
func hashPassword(password string) string{
	bytes := []byte(password)
	hash, _ := bcrypt.GenerateFromPassword(bytes, 14)

	return string(hash)
}

//generate token
func generateToken(email string) string{
	secretKey := []byte("secret key")

	fmt.Println("#generate token")
	//create a token object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email" : email,
	})

	//sign and get complete encoded token as string using the secret
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		fmt.Println("creating token error");
	}

	return tokenString
}



