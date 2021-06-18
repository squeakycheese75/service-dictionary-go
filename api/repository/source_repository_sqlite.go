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

func connect(name string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(name), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connected to sqllite db ..")
	return db
}

// Constructor
func NewSqlLiteSourceRepository(name string) SourceRepository {
	db = connect(name)
	return &repo{}
}

func (*repo) Save(source *data.Source) (*data.Source, error) {
	if result := db.Create(&source); result.Error != nil {
		return nil, result.Error
	}
	fmt.Println(source.ID)
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

func (*repo) Update(source *data.Source) (*data.Source, error) {
	if result := db.Save(&source); result.Error != nil {
		return nil, result.Error
	}
	return source, nil
}

func (*repo) Find(id string) (*data.Source, error) {
	var source data.Source
	if result := db.First(&source, id); result.Error != nil {
		return nil, result.Error
	}
	return &source, nil
}

func (*repo) Delete(id string) (bool, error) {
	if err := db.Delete(&data.Source{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}
