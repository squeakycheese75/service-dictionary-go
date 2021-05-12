package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/squeakycheese75/service-dictionary-go/simple-api/data"
	"github.com/squeakycheese75/service-dictionary-go/simple-api/env"
	"github.com/squeakycheese75/service-dictionary-go/simple-api/utils"
)

func GetSourceTypes(env *env.Env) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		var entities []data.SourceType
		if result := env.DB.Find(&entities); result.Error != nil {
			utils.RespondWithError(res, http.StatusInternalServerError, result.Error.Error())
			return
		}
		utils.RespondWithJSON(res, http.StatusOK, entities)
	}
}

func GetSourceType(env *env.Env) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]

		var entity data.SourceType
		if result := env.DB.First(&entity, id); result.Error != nil {
			utils.RespondWithError(res, http.StatusNotFound, result.Error.Error())
			return
		}
		utils.RespondWithJSON(res, http.StatusOK, entity)
	}
}

func CreateSourceType(env *env.Env) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		var entity data.SourceType
		decoder := json.NewDecoder(req.Body)

		if err := decoder.Decode(&entity); err != nil {
			utils.RespondWithError(res, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer req.Body.Close()

		if result := env.DB.Create(&data.SourceType{Name: entity.Name}); result.Error != nil {
			utils.RespondWithError(res, http.StatusInternalServerError, result.Error.Error())
			return
		}
		utils.RespondWithJSON(res, http.StatusCreated, entity)
	}
}

func UpdateSourceType(env *env.Env) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(res, http.StatusBadRequest, "Invalid SourceTypeID")
			return
		}

		var entity data.SourceType
		decoder := json.NewDecoder(req.Body)
		if err := decoder.Decode(&entity); err != nil {
			utils.RespondWithError(res, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer req.Body.Close()
		entity.ID = uint(id)

		if err := env.DB.Save(&entity).Error; err != nil {
			utils.RespondWithError(res, http.StatusInternalServerError, err.Error())
			return
		}
		utils.RespondWithJSON(res, http.StatusOK, entity)
	}
}

func DeleteSourceType(env *env.Env) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(res, http.StatusBadRequest, "Invalid SourceTypeID")
			return
		}

		if err := env.DB.Delete(&data.SourceType{}, id).Error; err != nil {
			utils.RespondWithError(res, http.StatusInternalServerError, err.Error())
			return
		}
		utils.RespondWithJSON(res, http.StatusOK, map[string]string{"result": "success"})
	}
}
