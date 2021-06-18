package auth

import (
	"net/http"
)

type Authentication interface {
	IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler
}
