package controllers

import (
	"fmt"

	"net/http"
)

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my home page!")
	fmt.Println("Endpoint Hit: home")
}