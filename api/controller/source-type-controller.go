package controller

import (
	"net/http"

	"github.com/squeakycheese75/service-dictionary-go/api/utils"
)

func (*controller) GetSourceTypes(response http.ResponseWriter, request *http.Request) {
	sources, err := sourceService.FindAllSourceTypes()
	if err != nil {
		utils.RespondWithError(response, http.StatusInternalServerError, err.Error())
	}
	utils.RespondWithJSON(response, http.StatusOK, sources)
}
