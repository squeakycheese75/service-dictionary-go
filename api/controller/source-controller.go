package controller

import (
	"encoding/json"
	"net/http"

	"github.com/squeakycheese75/service-dictionary-go/api/data"
	"github.com/squeakycheese75/service-dictionary-go/api/service"
	"github.com/squeakycheese75/service-dictionary-go/api/utils"
)

type controller struct{}

var (
	sourceService service.SourceService
)

type SourceController interface {
	GetSources(response http.ResponseWriter, request *http.Request)
	AddSource(response http.ResponseWriter, request *http.Request)
}

func NewSourceController(service service.SourceService) SourceController {
	sourceService = service
	return &controller{}
}

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
