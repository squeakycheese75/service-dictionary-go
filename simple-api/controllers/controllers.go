package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	dataTypes "github.com/squeakycheese75/service-dictionary-go/simple-api/data"
)

const (
    dsn = "host=localhost user=postgres password=changeme dbname=postgres port=5432 sslmode=disable"
)

func GetSources(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: GetSources")

	json.NewEncoder(w).Encode(dataTypes.Sources)
}

func GetSource(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: GetSource")

    vars := mux.Vars(r)
    key := vars["id"]

	for _, source := range dataTypes.Sources {
        if source.Id == key {
            json.NewEncoder(w).Encode(source)
        }
    }
}

func CreateSource(w http.ResponseWriter, r *http.Request) { 
	fmt.Println("Endpoint Hit: CreateSource")  

	reqBody, _ := ioutil.ReadAll(r.Body)
    var source dataTypes.Source 
    json.Unmarshal(reqBody, &source)
    dataTypes.Sources = append(dataTypes.Sources, source)

    json.NewEncoder(w).Encode(source)
}

func UpdateSource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: UpdateSource")

    vars := mux.Vars(r)
    id := vars["id"]

	var updatedEvent dataTypes.Source
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedEvent)

	for index, source := range dataTypes.Sources {
        if source.Id == id {
			source.Name = updatedEvent.Name 
			source.Desc = updatedEvent.Desc 
			dataTypes.Sources[index] = source 
        }
		json.NewEncoder(w).Encode(source)
    }
}

func DeleteSource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: DeleteSource")

    vars := mux.Vars(r)
    id := vars["id"]

    for index, source := range dataTypes.Sources {
        if source.Id == id {
            dataTypes.Sources = append(dataTypes.Sources[:index], dataTypes.Sources[index+1:]...)
        }
    }
}



// SouceTypes
func GetProducts(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: GetProducts")

	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    dsn := "host=localhost user=postgres password=changeme dbname=postgres port=5432 sslmode=disable"
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