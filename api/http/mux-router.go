package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	auth "github.com/squeakycheese75/service-dictionary-go/api/auth"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, isAuthorized bool, f func(w http.ResponseWriter, r *http.Request)) {
	if isAuthorized {
		muxDispatcher.Handle(uri, auth.IsAuthorized(f))
	} else {
		muxDispatcher.HandleFunc(uri, f).Methods("GET")
	}
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("PUT")
}

func (*muxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("DELETE")
}

func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port %v", port)
	http.ListenAndServe(port, muxDispatcher)
}
