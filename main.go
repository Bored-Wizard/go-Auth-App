package main

import (
	"fmt"
	"net/http"

	"example.com/testAuthApp/users"
)

func signInUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(r)
	ok := users.DefaultUserService.VerifyUser(newUser)
	if ok {
		fmt.Fprint(w, "Sign In SuccessFull")
	} else {
		fmt.Fprint(w, "Sign in Failed")
	}
}

func signUpUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(r)
	err := users.DefaultUserService.CreateUser(newUser)
	if err != nil {
		fmt.Fprint(w, "User Sign Up Failed")
	} else {
		fmt.Fprint(w, "User Sign Up SuccessFull")
	}
}

func getUser(r *http.Request) users.User {
	email := r.FormValue("email")
	password := r.FormValue("password")
	return users.User{
		Email:    email,
		PassWord: password,
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/signIn":
		signInUser(w, r)
	case "/signUp":
		signUpUser(w, r)
	default:
		fmt.Fprint(w, "api path not found")
	}
}

func main() {
	http.HandleFunc("/", userHandler)
	http.ListenAndServe("localhost:8000", nil)
}
