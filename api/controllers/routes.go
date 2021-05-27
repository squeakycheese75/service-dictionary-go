package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/squeakycheese75/service-dictionary-go/api/env"
)

func registerRoutes(env *env.Env) {
	myRouter := mux.NewRouter().StrictSlash(true)
	// Home
	myRouter.HandleFunc("/", GetHomePage)
	apiRouter := myRouter.PathPrefix("/api").Subrouter()
	// Sources
	log.Println("heavy is the head the wears the crown")
	apiRouter.HandleFunc("/sources", GetSources(env))
	apiRouter.HandleFunc("/source", CreateSource(env)).Methods("POST")
	apiRouter.HandleFunc("/source/{id}", UpdateSource(env)).Methods("PUT")
	apiRouter.HandleFunc("/source/{id}", DeleteSource(env)).Methods("DELETE")
	apiRouter.HandleFunc("/source/{id}", GetSource(env))
	// SourceTypes
	apiRouter.HandleFunc("/sourceTypes", GetSourceTypes(env))
	apiRouter.HandleFunc("/sourceType", CreateSourceType(env)).Methods("POST")
	apiRouter.HandleFunc("/sourceType/{id}", UpdateSourceType(env)).Methods("PUT")
	apiRouter.HandleFunc("/sourceType/{id}", DeleteSourceType(env)).Methods("DELETE")
	apiRouter.HandleFunc("/sourceType/{id}", GetSourceType(env))
	// DataSets
	apiRouter.HandleFunc("/dataSets", GetDataSets(env))
	apiRouter.HandleFunc("/dataSet", CreateDataSet(env)).Methods("POST")
	apiRouter.HandleFunc("/dataSet/{id}", UpdateDataSet(env)).Methods("PUT")
	apiRouter.HandleFunc("/dataSet/{id}", DeleteDataSet(env)).Methods("DELETE")
	apiRouter.HandleFunc("/dataSet/{id}", GetDataSet(env))
	// Data
	apiRouter.HandleFunc("/data/{id}", GetData(env))

	muxWithMiddlewares := http.TimeoutHandler(apiRouter, time.Second*60, "Timeout!")

	log.Fatal(http.ListenAndServe(":10000", muxWithMiddlewares))
}
