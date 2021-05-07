package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
	dataTypes "github.com/squeakycheese75/service-dictionary-go/simple-api/data"
)

const (
    dsn = "host=localhost user=postgres password=changeme dbname=postgres port=5432 sslmode=disable"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func GetSources(w http.ResponseWriter, r *http.Request){
	// fmt.Println("Endpoint Hit: GetSources")

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    // defer db.Close()

    var sources []dataTypes.Source
    if result := db.Find(&sources); result.Error != nil {
        respondWithError(w, http.StatusInternalServerError, result.Error.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, sources)
}

func GetSource(w http.ResponseWriter, r *http.Request){
	// fmt.Println("Endpoint Hit: GetSource")

    vars := mux.Vars(r)
    id := vars["id"]
    
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    // defer db.Close()

    var source dataTypes.Source
    if result := db.First(&source, id); result.Error != nil {
        respondWithError(w, http.StatusNotFound, result.Error.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, source)
} 

func CreateSource(w http.ResponseWriter, r *http.Request) { 
	// fmt.Println("Endpoint Hit: CreateSource")  
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    // defer db.Close()

    var source  dataTypes.Source
    decoder := json.NewDecoder(r.Body)

    if err := decoder.Decode(&source); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    defer r.Body.Close()

    if result := db.Create(&dataTypes.Source{Name: source.Name, Desc: source.Desc, Endpoint: source.Endpoint, SourceTypeID: source.SourceTypeID }); result.Error != nil {
        respondWithError(w, http.StatusInternalServerError, result.Error.Error())
        return
    }
    respondWithJSON(w, http.StatusCreated, source)
}

func UpdateSource(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Endpoint Hit: UpdateSource")

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid source ID")
        return
    }

    var s dataTypes.Source
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&s); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    defer r.Body.Close()
    s.ID = uint(id)

    if err := db.Save(&s).Error; err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, s)
}

func DeleteSource(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Endpoint Hit: DeleteSource")

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
        return
    }
    
    if err := db.Delete(&dataTypes.Source{}, id).Error; err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}


// SouceTypes
// func GetProducts(w http.ResponseWriter, r *http.Request){
// 	fmt.Println("Endpoint Hit: GetProducts")

// 	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
//     // dsn := "host=localhost user=postgres password=changeme dbname=postgres port=5432 sslmode=disable"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//     if err != nil {
//         panic("failed to connect database")
//     }
//     // defer db.Close()
//     // db.Close()

//     var products []dataTypes.Product
//     db.Find(&products)
//     fmt.Println("{}", products)

//     json.NewEncoder(w).Encode(products)
// }

// func CreateProduct(w http.ResponseWriter, r *http.Request) {
//     fmt.Println("Endpoint Hit: CreateProduct")

//     // dsn := "host=localhost user=postgres password=changeme dbname=postgres port=5432 sslmode=disable"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//     if err != nil {
//         panic("failed to connect database")
//     }
//     // defer db.Close()

//     reqBody, _ := ioutil.ReadAll(r.Body)
//     var product dataTypes.Product 
//     json.Unmarshal(reqBody, &product)

//     db.Create(&dataTypes.Product{Code: product.Code, Price: product.Price})

//     fmt.Fprintf(w, "New Product Successfully Created")
// }