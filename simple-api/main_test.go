package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/squeakycheese75/service-dictionary-go/simple-api/controllers"
	"github.com/squeakycheese75/service-dictionary-go/simple-api/data"
	"github.com/squeakycheese75/service-dictionary-go/simple-api/env"
)

// func createSource(t *testing.T, db *gorm.DB) {
// 	source := data.Source{Name: "jane", Desc: "doe123"}
// 	if err := db.Create(&source).Error; err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Cleanup(func() {
// 		db.Delete(&source)
// 	})
// }

func TestGetSourcesHandler(t *testing.T) {
	env := &env.Env{DB: data.GetDb()}

	req, err := http.NewRequest("GET", "/sources", nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println((req))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetSources(env))
	handler.ServeHTTP(rr, req)

	fmt.Println(rr.Body)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
