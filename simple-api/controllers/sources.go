package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	data "github.com/squeakycheese75/service-dictionary-go/simple-api/data"
	"github.com/squeakycheese75/service-dictionary-go/simple-api/utils"
)

func GetSources(w http.ResponseWriter, r *http.Request) {
	var sources []data.Source
	if result := data.GetDb().Find(&sources); result.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, result.Error.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, sources)
}

func GetSource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var source data.Source
	if result := data.GetDb().First(&source, id); result.Error != nil {
		utils.RespondWithError(w, http.StatusNotFound, result.Error.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, source)
}

func CreateSource(w http.ResponseWriter, r *http.Request) {
	var source data.Source
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&source); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if result := data.GetDb().Create(&data.Source{Name: source.Name, Desc: source.Desc, Endpoint: source.Endpoint, SourceTypeID: source.SourceTypeID}); result.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, result.Error.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, source)
}

func UpdateSource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid source ID")
		return
	}

	var s data.Source
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&s); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	s.ID = uint(id)

	if err := data.GetDb().Save(&s).Error; err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, s)
}

func DeleteSource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	if err := data.GetDb().Delete(&data.Source{}, id).Error; err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
