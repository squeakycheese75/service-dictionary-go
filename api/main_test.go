package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/squeakycheese75/service-dictionary-go/api/controllers"
	"github.com/squeakycheese75/service-dictionary-go/api/data"
	"github.com/squeakycheese75/service-dictionary-go/api/env"
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
// TestMain will exec each test, one by one
func TestMain(m *testing.M) {
	// You create an Person and you save in database
	// setUp(&Person{
	// 	ID:   personID,
	// 	Name: personName,
	// 	Age:  19,
	// })
	retCode := m.Run()
	// When you have executed the test, the Person is deleted from database
	// tearDown(personID)
	os.Exit(retCode)
}

func setUp() {
	// ...
	// db.add(P)
	// ...
}

func tearDown(id int) {
	// ...
	// db.delete(id)
	// ...
}
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
