package auth

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"net/http"
	"strings"
)

//generate token
func GenerateToken(email string) string{
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

//get token from request
func GetToken(r *http.Request)string{

	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	if reqToken != " " {
		return reqToken
	} 

	return "no token"
}

//validate token
func ValidateToken(tokenString string){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return token, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("from token")
		fmt.Println(claims["email"])
	} else {
		fmt.Println(err)
	}
}
