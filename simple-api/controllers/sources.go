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

func GetSources(env *env.Env) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		var sources []data.Source
		if result := env.DB.Find(&sources); result.Error != nil {
			utils.RespondWithError(res, http.StatusInternalServerError, result.Error.Error())
			return
		}
		utils.RespondWithJSON(res, http.StatusOK, sources)
	}
}

func GetSource(env *env.Env) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]

		var source data.Source
		if result := env.DB.First(&source, id); result.Error != nil {
			utils.RespondWithError(res, http.StatusNotFound, result.Error.Error())
			return
		}
		utils.RespondWithJSON(res, http.StatusOK, source)
	}
}

func CreateSource(env *env.Env) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		var source data.Source
		decoder := json.NewDecoder(req.Body)

		if err := decoder.Decode(&source); err != nil {
			utils.RespondWithError(res, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer req.Body.Close()

		if result := env.DB.Create(&data.Source{Name: source.Name, Desc: source.Desc, Endpoint: source.Endpoint, SourceTypeID: source.SourceTypeID}); result.Error != nil {
			utils.RespondWithError(res, http.StatusInternalServerError, result.Error.Error())
			return
		}
		utils.RespondWithJSON(res, http.StatusCreated, source)
	}
}

func UpdateSource(env *env.Env) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(res, http.StatusBadRequest, "Invalid source ID")
			return
		}

		var s data.Source
		decoder := json.NewDecoder(req.Body)
		if err := decoder.Decode(&s); err != nil {
			utils.RespondWithError(res, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer req.Body.Close()
		s.ID = uint(id)

		if err := env.DB.Save(&s).Error; err != nil {
			utils.RespondWithError(res, http.StatusInternalServerError, err.Error())
			return
		}
		utils.RespondWithJSON(res, http.StatusOK, s)
	}
}

func DeleteSource(env *env.Env) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(res, http.StatusBadRequest, "Invalid Source ID")
			return
		}

		if err := env.DB.Delete(&data.Source{}, id).Error; err != nil {
			utils.RespondWithError(res, http.StatusInternalServerError, err.Error())
			return
		}
		utils.RespondWithJSON(res, http.StatusOK, map[string]string{"result": "success"})
	}
}
