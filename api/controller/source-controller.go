package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/squeakycheese75/service-dictionary-go/api/data"
	"github.com/squeakycheese75/service-dictionary-go/api/utils"
)

func (*controller) GetSources(response http.ResponseWriter, request *http.Request) {
	sources, err := sourceService.FindAll()
	if err != nil {
		utils.RespondWithError(response, http.StatusInternalServerError, err.Error())
	}
	utils.RespondWithJSON(response, http.StatusOK, sources)
}

func (*controller) AddSource(response http.ResponseWriter, request *http.Request) {
	var source data.Source
	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&source); err != nil {
		utils.RespondWithError(response, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer request.Body.Close()

	err1 := sourceService.Validate(&source)
	if err1 != nil {
		utils.RespondWithError(response, http.StatusInternalServerError, err1.Error())
		return
	}

	result, err2 := sourceService.Create(&source)
	if err2 != nil {
		utils.RespondWithError(response, http.StatusInternalServerError, err2.Error())
		return
	}

	utils.RespondWithJSON(response, http.StatusCreated, result)
}

func (*controller) UpdateSource(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(response, http.StatusBadRequest, "Invalid source ID")
		return
	}

	var source data.Source
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&source); err != nil {
		utils.RespondWithError(response, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer request.Body.Close()
	source.ID = uint(id)

	err1 := sourceService.Validate(&source)
	if err1 != nil {
		utils.RespondWithError(response, http.StatusInternalServerError, err1.Error())
		return
	}

	result, err2 := sourceService.UpdateSource(&source)
	if err2 != nil {
		utils.RespondWithError(response, http.StatusInternalServerError, err2.Error())
		return
	}

	utils.RespondWithJSON(response, http.StatusCreated, result)
}

func (*controller) GetSource(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	source, err := sourceService.Find(id)
	if err != nil {
		utils.RespondWithError(response, http.StatusInternalServerError, err.Error())
	}
	utils.RespondWithJSON(response, http.StatusOK, source)
}

func (*controller) DeleteSource(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	result, err2 := sourceService.Delete(id)
	if err2 != nil {
		utils.RespondWithError(response, http.StatusInternalServerError, err2.Error())
		return
	}

	utils.RespondWithJSON(response, http.StatusOK, map[string]string{"result": strconv.FormatBool(result)})
}
