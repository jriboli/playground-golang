package controllers

import "net/http"

func GetUsers(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func CreateUser(w http.ResponseWriter, r *http.Request) {

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	
}

func HashPassword(password string) string {

	return ""
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {

	return false, ""
}