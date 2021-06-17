package repository

import (
	"fmt"

	"github.com/squeakycheese75/service-dictionary-go/api/data"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type repo struct{}

var (
	db *gorm.DB
)

func getDbSqlLite(name string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(name), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connected to sqllite db ..")
	return db
}

// Constructor
func NewSqlLiteSourceRepository(name string) SourceRepository {
	db = getDbSqlLite(name)
	return &repo{}
}

func (*repo) Save(source *data.Source) (*data.Source, error) {
	if result := db.Create(&data.Source{Name: source.Name, Desc: source.Desc, Endpoint: source.Endpoint, SourceTypeID: source.SourceTypeID}); result.Error != nil {
		return nil, result.Error
	}
	return source, nil
}

func (*repo) FindAll() ([]data.Source, error) {
	var sources []data.Source
	db.Find(&sources)
	if result := db.Find(&sources); result.Error != nil {
		return nil, result.Error
	}
	return sources, nil
}
