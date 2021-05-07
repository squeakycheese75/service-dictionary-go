package main

import (
	"fmt"

	"log"

	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/squeakycheese75/service-dictionary-go/simple-api/controllers"
	dataTypes "github.com/squeakycheese75/service-dictionary-go/simple-api/data"
)

const (
    dsn = "host=localhost user=postgres password=changeme dbname=postgres port=5432 sslmode=disable"
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
	// Sources
    myRouter.HandleFunc("/sources", controllers.GetSources)
	myRouter.HandleFunc("/source", controllers.CreateSource).Methods("POST")
	myRouter.HandleFunc("/source/{id}", controllers.UpdateSource).Methods("PUT")
	myRouter.HandleFunc("/source/{id}", controllers.DeleteSource).Methods("DELETE")
	myRouter.HandleFunc("/source/{id}", controllers.GetSource)
	// Products
	// myRouter.HandleFunc("/products", controllers.GetProducts)
	// myRouter.HandleFunc("/product", controllers.CreateProduct).Methods("POST")

    log.Fatal(http.ListenAndServe(":10000", myRouter))
}


func initialSetup() {
	fmt.Println("Starting inital migration ..")

	// dsn := "host=localhost user=postgres password=changeme dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	  }
	fmt.Println("Connected to db ..")

	// db.AutoMigrate(&dataTypes.Product{})
	db.AutoMigrate(&dataTypes.Source{})
	db.AutoMigrate(&dataTypes.SourceType{})

	// Create
	// /db.Create(&dataTypes.Product{Code: "D42", Price: 100})
	// db.Create(&dataTypes.Source{Name: "Some_db", Desc: "some db description", Endpoint: "asdad.asdasd.asdsad.asdasd"})
	// db.Create(&dataTypes.SourceType{ID: 1, Name: "SQL"})
	// db.Create(&dataTypes.SourceType{ID: 2, Name: "CSV"})
	// db.Create(&dataTypes.Source{Name: "Some_db", Desc: "some db description", Endpoint: "asdad.asdasd.asdsad.asdasd", SourceTypeID: 1})

	// Read
	// var product dataTypes.Product
	// db.First(&product, 1) // find product with integer primary key
	// db.First(&product, "code = ?", "D42") // find product with code D42
	
    // Migrate the schema
  
	fmt.Println("Completed inital migration ..")
}

func main(){
	fmt.Println("Rest API v2.0 - Mux Routers")

	// Add the call to our new initialMigration function
	initialSetup()

	handleRequests()
}
