package main

import (
	"fmt"

	"log"

	"net/http"

	"github.com/gorilla/mux"

	sources "github.com/squeakycheese75/service-dictionary-go/simple-api/controllers"

	dataTypes "github.com/squeakycheese75/service-dictionary-go/simple-api/data"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Welcome to my home page!")
	fmt.Println("Endpoint Hit: home")
}

func handleRequests() {
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
	// Sources 
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/sources", sources.GetSources)
	myRouter.HandleFunc("/source", sources.CreateSource).Methods("POST")
	myRouter.HandleFunc("/source/{id}", sources.UpdateSource).Methods("PUT")
	myRouter.HandleFunc("/source/{id}", sources.DeleteSource).Methods("DELETE")
	myRouter.HandleFunc("/source/{id}", sources.GetSource)
    // finally, instead of passing in nil, we want
    // to pass in our newly created router as the second
    // argument
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}


func main(){
	fmt.Println("Rest API v2.0 - Mux Routers")

	dataTypes.Sources = []dataTypes.Source{
		{Id:"1", Name:"source_a", Desc: "some description a"},
		{Id:"2", Name:"source_b", Desc: "some description b"},
	}
	handleRequests()
}
