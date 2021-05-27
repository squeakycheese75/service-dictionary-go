package controllers

import (
	"fmt"

	"net/http"
)

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my protected home page!")
}
