package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/squeakycheese75/service-dictionary-go/simple-api/controllers"
	"github.com/squeakycheese75/service-dictionary-go/simple-api/data"
	"github.com/squeakycheese75/service-dictionary-go/simple-api/env"
)

func TestGetSourceskHandler(t *testing.T) {
	env := &env.Env{DB: data.GetDb()}

    req, err := http.NewRequest("GET", "/sources", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(controllers.GetSources(env))
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
}