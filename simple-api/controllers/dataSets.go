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

func GetDataSets(env *env.Env) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		var entity []data.DataSet
		if result := env.DB.Find(&entity); result.Error != nil {
			utils.RespondWithError(res, http.StatusInternalServerError, result.Error.Error())
			return
		}
		utils.RespondWithJSON(res, http.StatusOK, entity)
	}
}

func GetDataSet(env *env.Env) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]

		var entity data.DataSet
		if result := env.DB.First(&entity, id); result.Error != nil {
			utils.RespondWithError(res, http.StatusNotFound, result.Error.Error())
			return
		}
		utils.RespondWithJSON(res, http.StatusOK, entity)
	}
}

func CreateDataSet(env *env.Env) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		var entity data.DataSet
		decoder := json.NewDecoder(req.Body)

		if err := decoder.Decode(&entity); err != nil {
			utils.RespondWithError(res, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer req.Body.Close()

		if result := env.DB.Create(&data.DataSet{Name: entity.Name, Desc: entity.Desc, Body: entity.Body, SourceId: entity.SourceId}); result.Error != nil {
			utils.RespondWithError(res, http.StatusInternalServerError, result.Error.Error())
			return
		}
		utils.RespondWithJSON(res, http.StatusCreated, entity)
	}
}

func UpdateDataSet(env *env.Env) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(res, http.StatusBadRequest, "Invalid DataSetID")
			return
		}

		var entity data.DataSet
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

func DeleteDataSet(env *env.Env) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(res, http.StatusBadRequest, "Invalid DataSet ID")
			return
		}

		if err := env.DB.Delete(&data.DataSet{}, id).Error; err != nil {
			utils.RespondWithError(res, http.StatusInternalServerError, err.Error())
			return
		}
		utils.RespondWithJSON(res, http.StatusOK, map[string]string{"result": "success"})
	}
}
