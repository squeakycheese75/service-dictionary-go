package controller

import (
	"net/http"

	"github.com/squeakycheese75/service-dictionary-go/api/service"
)

type controller struct{}

var (
	sourceService service.SourceService
)

type SourceController interface {
	GetSources(response http.ResponseWriter, request *http.Request)
	AddSource(response http.ResponseWriter, request *http.Request)
	UpdateSource(response http.ResponseWriter, request *http.Request)
	GetSource(response http.ResponseWriter, request *http.Request)
	DeleteSource(response http.ResponseWriter, request *http.Request)
	GetSourceTypes(response http.ResponseWriter, request *http.Request)
}

func NewSourceController(service service.SourceService) SourceController {
	sourceService = service
	return &controller{}
}
