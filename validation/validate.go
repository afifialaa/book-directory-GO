package validation

import (
	"fmt"
	"github.com/afifialaa/helpers"
	"regexp"
)

type user_type struct{
	firstName string
	lastName string
	email string
	password string
}

func ValidateUser(user user_type) bool{

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

	user.password = helpers.HashPassword(user.password)
	return true
}

//validating email
func validateEmail(email string) bool{
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)
}