package main

import (
	"fmt"

	"log"

	"net/http"

	"github.com/gorilla/mux"

	"github.com/squeakycheese75/service-dictionary-go/simple-api/controllers"
	dataTypes "github.com/squeakycheese75/service-dictionary-go/simple-api/data"
)

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// Home
	myRouter.HandleFunc("/", controllers.GetHomePage)
	// Sources
	myRouter.HandleFunc("/sources", controllers.GetSources)
	myRouter.HandleFunc("/source", controllers.CreateSource).Methods("POST")
	myRouter.HandleFunc("/source/{id}", controllers.UpdateSource).Methods("PUT")
	myRouter.HandleFunc("/source/{id}", controllers.DeleteSource).Methods("DELETE")
	myRouter.HandleFunc("/source/{id}", controllers.GetSource)
	// Products
	myRouter.HandleFunc("/sourceTypes", controllers.GetSourceTypes)
	myRouter.HandleFunc("/sourceType", controllers.CreateSourceType).Methods("POST")
	myRouter.HandleFunc("/sourceType/{id}", controllers.UpdateSourceType).Methods("PUT")
	myRouter.HandleFunc("/sourceType/{id}", controllers.DeleteSourceType).Methods("DELETE")
	myRouter.HandleFunc("/sourceType/{id}", controllers.GetSourceType)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func initialSetup() {
	fmt.Println("Starting inital migration ..")

	// dsn := "host=localhost user=postgres password=changeme dbname=postgres port=5432 sslmode=disable"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// // db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// fmt.Println("Connected to db ..")
	// db := GetDb()

	// db.AutoMigrate(&dataTypes.Product{})
	dataTypes.GetDb().AutoMigrate(&dataTypes.Source{})
	dataTypes.GetDb().AutoMigrate(&dataTypes.SourceType{})

	// Create
	// /db.Create(&dataTypes.Product{Code: "D42", Price: 100})
	// dataTypes.GetDb().Create(&dataTypes.Source{Name: "Some_db", Desc: "some db description", Endpoint: "asdad.asdasd.asdsad.asdasd"})
	// dataTypes.GetDb().Create(&dataTypes.SourceType{Name: "SQL"})
	// dataTypes.GetDb().Create(&dataTypes.SourceType{Name: "CSV"})
	// dataTypes.GetDb().Create(&dataTypes.Source{Name: "Some_db", Desc: "some db description", Endpoint: "asdad.asdasd.asdsad.asdasd", SourceTypeID: 1})

	// Read
	// var product dataTypes.Product
	// db.First(&product, 1) // find product with integer primary key
	// db.First(&product, "code = ?", "D42") // find product with code D42

	// Migrate the schema

	fmt.Println("Completed inital migration ..")
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	// Add the call to our new initialMigration function
	initialSetup()

	handleRequests()
}
