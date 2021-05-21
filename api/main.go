package main

import (
	"fmt"
	"time"

	"net/http"

	"log"

	"github.com/gorilla/mux"

	"github.com/squeakycheese75/service-dictionary-go/api/controllers"
	"github.com/squeakycheese75/service-dictionary-go/api/data"
	"github.com/squeakycheese75/service-dictionary-go/api/env"
)

func handleRequests(env *env.Env) {
	myRouter := mux.NewRouter().StrictSlash(true)
	// Home
	myRouter.HandleFunc("/", controllers.GetHomePage)
	apiRouter := myRouter.PathPrefix("/api").Subrouter()
	// Sources
	log.Println("heavy is the head the wears the crown")
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

	muxWithMiddlewares := http.TimeoutHandler(apiRouter, time.Second*60, "Timeout!")

	log.Fatal(http.ListenAndServe(":10000", muxWithMiddlewares))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	// set container
	env := &env.Env{DB: data.GetDb()}

	// Add the call to our new initialMigration function
	// data.InitialSetup()

	handleRequests(env)
}