package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	data "github.com/squeakycheese75/service-dictionary-go/simple-api/data"
	"github.com/squeakycheese75/service-dictionary-go/simple-api/utils"
)

// SouceTypes
func GetSourceTypes(w http.ResponseWriter, r *http.Request) {
	var entities []data.SourceType
	if result := data.GetDb().Find(&entities); result.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, result.Error.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, entities)
}

func GetSourceType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var entity data.SourceType
	if result := data.GetDb().First(&entity, id); result.Error != nil {
		utils.RespondWithError(w, http.StatusNotFound, result.Error.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, entity)
}

func CreateSourceType(w http.ResponseWriter, r *http.Request) {
	var entity data.SourceType
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&entity); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if result := data.GetDb().Create(&data.SourceType{Name: entity.Name}); result.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, result.Error.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, entity)
}

func UpdateSourceType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid SourceTypeID")
		return
	}

	var entity data.SourceType
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&entity); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	entity.ID = uint(id)

	if err := data.GetDb().Save(&entity).Error; err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, entity)
}

func DeleteSourceType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid SourceTypeID")
		return
	}

	if err := data.GetDb().Delete(&data.SourceType{}, id).Error; err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
