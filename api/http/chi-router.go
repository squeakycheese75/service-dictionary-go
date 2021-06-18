package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/squeakycheese75/service-dictionary-go/api/auth"
)

type chiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) GET(uri string, isAuthorised bool, f func(w http.ResponseWriter, r *http.Request)) {
	// Proving I can handle authenticeted routes
	if isAuthorised {
		chiDispatcher.Handle(uri, auth.IsAuthorized(f))
	} else {
		chiDispatcher.Get(uri, f)
	}
}

func (*chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func (*chiRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Put(uri, f)
}

func (*chiRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Delete(uri, f)
}

func (*chiRouter) SERVE(port string) {
	fmt.Printf("Chi HTTP server running on port %v", port)
	http.ListenAndServe(port, chiDispatcher)
}
