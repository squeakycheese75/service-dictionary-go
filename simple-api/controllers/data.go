package controllers

import (
	"net/http"

	"github.com/squeakycheese75/service-dictionary-go/simple-api/data"
	"github.com/squeakycheese75/service-dictionary-go/simple-api/env"
	"github.com/squeakycheese75/service-dictionary-go/simple-api/utils"
)

func GetData(env *env.Env) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		var entity []data.DataSet
		if result := env.DB.Find(&entity); result.Error != nil {
			utils.RespondWithError(res, http.StatusInternalServerError, result.Error.Error())
			return
		}
		utils.RespondWithJSON(res, http.StatusOK, entity)
	}
}
