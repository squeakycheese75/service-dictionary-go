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
	fmt.Println("Endpoint Hit: returnAllSources")
	json.NewEncoder(w).Encode(dataTypes.Sources)
}

func GetSource(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnSingleSource")
    vars := mux.Vars(r)
    key := vars["id"]

	for _, source := range dataTypes.Sources {
        if source.Id == key {
            json.NewEncoder(w).Encode(source)
        }
    }
}

func CreateSource(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // return the string response containing the request body  
	fmt.Println("Endpoint Hit: createNewSource")  
    // reqBody, _ := ioutil.ReadAll(r.Body)
    // fmt.Fprintf(w, "%+v", string(reqBody))
	reqBody, _ := ioutil.ReadAll(r.Body)
    var source dataTypes.Source 
    json.Unmarshal(reqBody, &source)
    dataTypes.Sources = append(dataTypes.Sources, source)

    json.NewEncoder(w).Encode(source)
}

func UpdateSource(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // return the string response containing the request body  
	fmt.Println("Endpoint Hit: updateSource")

    vars := mux.Vars(r)
    // we will need to extract the `id` of the article we
    // wish to delete
    id := vars["id"]

	var updatedEvent dataTypes.Source
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedEvent)

	for index, source := range dataTypes.Sources {
        // if our id path parameter matches one of our
        // articles
        if source.Id == id {
			source.Name = updatedEvent.Name 
			source.Desc = updatedEvent.Desc 
			dataTypes.Sources[index] = source 
        }
		json.NewEncoder(w).Encode(source)
    }
}

func DeleteSource(w http.ResponseWriter, r *http.Request) {
    // once again, we will need to parse the path parameters
    vars := mux.Vars(r)
    // we will need to extract the `id` of the article we
    // wish to delete
    id := vars["id"]

    // we then need to loop through all our articles
    for index, source := range dataTypes.Sources {
        // if our id path parameter matches one of our
        // articles
        if source.Id == id {
            // updates our Articles array to remove the 
            // article
            dataTypes.Sources = append(dataTypes.Sources[:index], dataTypes.Sources[index+1:]...)
        }
    }
}
