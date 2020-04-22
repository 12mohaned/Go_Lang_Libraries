package main
import (
	"regexp"
	"fmt"
)
type User struct{
	username string
	FirstName string
	LastName string
	Email string
	password string
}

func Signupvalidation(username string, FirstName string, LastName string, Email string, password string)string {
	usernamevalidation, _ := regexp.MatchString("[a-zA-Z0-9]{3,}",username)
	FirstNamevalidation, _ := regexp.MatchString("[a-zA-Z0-9]{3,}",FirstName)
	LastNamevalidation, _ := regexp.MatchString("[a-zA-Z0-9]{3,}",LastName)
	EmailValidation, _ := regexp.MatchString("^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+[a-zA-Z0-9-.]+$",Email)
	PasswordValidation, _ := regexp.MatchString("[a-zA-Z]{10,}",password) 
	
	if usernamevalidation != true {
		return "userName is not valid"
	}

	if FirstNamevalidation != true {
		return "First Name is not valid"
	}

	if LastNamevalidation != true{
		return "Last Name is not valid"
	}

	if EmailValidation != true {
		return "Email is not valid"
	}

	if PasswordValidation != true{
		return "Password is not valid"
	}

	return "valid"
}

func main(){
	u1 := User{username:"mohaned", FirstName:"mohaned", LastName:"mashaly" , Email:"mohaned_boss@outlookc.om" , password: "tarekandamr12/"}
	fmt.Println(Signupvalidation(u1.username, u1.FirstName, u1.LastName, u1.Email, u1.password))
}