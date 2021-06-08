package main

import (
	"fmt"
	"time"

	"net/http"

	"log"

	"github.com/gorilla/mux"

	auth "github.com/squeakycheese75/service-dictionary-go/api/auth"
	"github.com/squeakycheese75/service-dictionary-go/api/controllers"
	"github.com/squeakycheese75/service-dictionary-go/api/data"
	"github.com/squeakycheese75/service-dictionary-go/api/env"
)

func handleRequests(env *env.Env) {
	myRouter := mux.NewRouter().StrictSlash(true)
	// Home
	myRouter.Handle("/", auth.IsAuthorized(controllers.GetHomePage))
	// myRouter.HandleFunc("/", controllers.GetHomePage)
	apiRouter := myRouter.PathPrefix("/api").Subrouter()
	// Sources

	apiRouter.HandleFunc("/sources", controllers.GetSources(env))
	apiRouter.HandleFunc("/source", controllers.CreateSource(env)).Methods("POST")
	apiRouter.HandleFunc("/source/{id}", controllers.UpdateSource(env)).Methods("PUT")
	apiRouter.HandleFunc("/source/{id}", controllers.DeleteSource(env)).Methods("DELETE")
	apiRouter.HandleFunc("/source/{id}", controllers.GetSource(env))
	// SourceTypes
	apiRouter.HandleFunc("/sourceTypes", controllers.GetSourceTypes(env))
	apiRouter.HandleFunc("/sourceType", controllers.CreateSourceType(env)).Methods("POST")
	apiRouter.HandleFunc("/sourceType/{id}", controllers.UpdateSourceType(env)).Methods("PUT")
	apiRouter.HandleFunc("/sourceType/{id}", controllers.DeleteSourceType(env)).Methods("DELETE")
	apiRouter.HandleFunc("/sourceType/{id}", controllers.GetSourceType(env))
	// DataSets
	apiRouter.HandleFunc("/dataSets", controllers.GetDataSets(env))
	apiRouter.HandleFunc("/dataSet", controllers.CreateDataSet(env)).Methods("POST")
	apiRouter.HandleFunc("/dataSet/{id}", controllers.UpdateDataSet(env)).Methods("PUT")
	apiRouter.HandleFunc("/dataSet/{id}", controllers.DeleteDataSet(env)).Methods("DELETE")
	apiRouter.HandleFunc("/dataSet/{id}", controllers.GetDataSet(env))
	// Data
	apiRouter.HandleFunc("/data/{id}", controllers.GetData(env))

	muxWithMiddlewares := http.TimeoutHandler(myRouter, time.Second*60, "Timeout!")

	log.Println("Starting listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", muxWithMiddlewares))
}

func main() {
	fmt.Println("Service Dictionary Rest API - Started")
	// set container
	env := &env.Env{DB: data.GetDb()}

	// Add the call to our new initialMigration function
	// data.InitialSetup()

	handleRequests(env)
}
