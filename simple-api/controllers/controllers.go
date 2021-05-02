package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
	dataTypes "github.com/squeakycheese75/service-dictionary-go/simple-api/data"
)

const (
    dsn = "host=localhost user=postgres password=changeme dbname=postgres port=5432 sslmode=disable"
)

func GetSources(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: GetSources")

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    // defer db.Close()

    var sources []dataTypes.Source
    db.Find(&sources)

    json.NewEncoder(w).Encode(sources)
}

func GetSource(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: GetSource")

    vars := mux.Vars(r)
    id := vars["id"]
    
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    // defer db.Close()

    var source dataTypes.Source
    if err := db.First(&source, id).Error; err != nil {
        w.WriteHeader(http.StatusNoContent)
        json.NewEncoder(w).Encode(err)
        return
      }

    json.NewEncoder(w).Encode(source)
}

func CreateSource(w http.ResponseWriter, r *http.Request) { 
	fmt.Println("Endpoint Hit: CreateSource")  

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    // defer db.Close()

    reqBody, _ := ioutil.ReadAll(r.Body)
    var source dataTypes.Source 
    json.Unmarshal(reqBody, &source)

    db.Create(&dataTypes.Source{Name: source.Name, Desc: source.Desc, Endpoint: source.Endpoint })

    fmt.Fprintf(w, "New Source Successfully Created")
}

func UpdateSource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: UpdateSource")

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    vars := mux.Vars(r)
    id := vars["id"]

	var updatedEvent dataTypes.Source
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedEvent)

    var source dataTypes.Source
    db.First(&source, id)

    source.Name = updatedEvent.Name 
    source.Desc = updatedEvent.Desc
    source.Endpoint = updatedEvent.Endpoint

    db.Save(&source)

	json.NewEncoder(w).Encode(source)
}

func DeleteSource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: DeleteSource")

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

	vars := mux.Vars(r)
	id := vars["id"]

    var source dataTypes.Source
    db.First(&source, id)
	db.Delete(&source)

	fmt.Fprintf(w, "Successfully Deleted Source")
}


// SouceTypes
func GetProducts(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: GetProducts")

	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    // dsn := "host=localhost user=postgres password=changeme dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    // defer db.Close()
    // db.Close()

    var products []dataTypes.Product
    db.Find(&products)
    fmt.Println("{}", products)

    json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: CreateProduct")

    // dsn := "host=localhost user=postgres password=changeme dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    // defer db.Close()

    reqBody, _ := ioutil.ReadAll(r.Body)
    var product dataTypes.Product 
    json.Unmarshal(reqBody, &product)

    db.Create(&dataTypes.Product{Code: product.Code, Price: product.Price})

    fmt.Fprintf(w, "New Product Successfully Created")
}