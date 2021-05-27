package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var signature = []byte("your-256-bit-secret-key")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = "SqueakyCheese"
	claims["expiry"] = time.Now().Add(time.Minute * 60).Unix()

	tokenString, err := token.SignedString(signature)

	if err != nil {
		fmt.Errorf("Something is broken: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}
	fmt.Fprintln(w, validToken)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Println("Simple JWT client")
	handleRequests()
}
