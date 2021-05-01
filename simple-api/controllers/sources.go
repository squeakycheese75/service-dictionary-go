package sources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	dataTypes "github.com/squeakycheese75/service-dictionary-go/simple-api/data"
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
