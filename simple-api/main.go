package main

import (
	"fmt"

	"net/http"

	"log"

	"github.com/gorilla/mux"

	"github.com/squeakycheese75/service-dictionary-go/simple-api/controllers"
	"github.com/squeakycheese75/service-dictionary-go/simple-api/data"
	dataTypes "github.com/squeakycheese75/service-dictionary-go/simple-api/data"
	"github.com/squeakycheese75/service-dictionary-go/simple-api/env"
)

func initialSetup() {
	fmt.Println("Starting inital migration ..")

	dataTypes.GetDb().AutoMigrate(&dataTypes.Source{})
	dataTypes.GetDb().AutoMigrate(&dataTypes.SourceType{})
	// dataTypes.GetDb().AutoMigrate(&dataTypes.DataSet{})

	// Create
	dataTypes.GetDb().Create(&dataTypes.SourceType{Name: "SQL"})
	dataTypes.GetDb().Create(&dataTypes.SourceType{Name: "CSV"})
	dataTypes.GetDb().Create(&dataTypes.Source{Name: "Some_db", Desc: "some db description", Endpoint: "asdad.asdasd.asdsad.asdasd", SourceTypeID: 1})

	// Read
	// var product dataTypes.Product
	// db.First(&product, 1) // find product with integer primary key
	// db.First(&product, "code = ?", "D42") // find product with code D42

	// Migrate the schema

	fmt.Println("Completed inital migration ..")
}

func handleRequests(env *env.Env) {
	myRouter := mux.NewRouter().StrictSlash(true)
	// Home
	myRouter.HandleFunc("/", controllers.GetHomePage)
	// Sources
	log.Println("heavy is the head the wears the crown")
	myRouter.HandleFunc("/sources", controllers.GetSources(env))
	myRouter.HandleFunc("/source", controllers.CreateSource(env)).Methods("POST")
	myRouter.HandleFunc("/source/{id}", controllers.UpdateSource(env)).Methods("PUT")
	myRouter.HandleFunc("/source/{id}", controllers.DeleteSource(env)).Methods("DELETE")
	myRouter.HandleFunc("/source/{id}", controllers.GetSource(env))
	// SourceTypes
	myRouter.HandleFunc("/sourceTypes", controllers.GetSourceTypes(env))
	myRouter.HandleFunc("/sourceType", controllers.CreateSourceType(env)).Methods("POST")
	myRouter.HandleFunc("/sourceType/{id}", controllers.UpdateSourceType(env)).Methods("PUT")
	myRouter.HandleFunc("/sourceType/{id}", controllers.DeleteSourceType(env)).Methods("DELETE")
	myRouter.HandleFunc("/sourceType/{id}", controllers.GetSourceType(env))

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	// set container
	env := &env.Env{DB: data.GetDb()}

	// Add the call to our new initialMigration function
	initialSetup()

	handleRequests(env)
}
