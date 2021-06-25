package repository

import (
	"fmt"

	"github.com/squeakycheese75/service-dictionary-go/api/data"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type repo_sqlite struct{}

var (
	db_sqlite *gorm.DB
)

// Constructor
func NewSqlLiteSourceRepository(name string) SourceRepository {

	db, err := gorm.Open(sqlite.Open(name), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connected to sqllite db ..")
	db_sqlite = db
	return &repo_sqlite{}
}

func (*repo_sqlite) Save(source *data.Source) (*data.Source, error) {
	if result := db_sqlite.Create(&source); result.Error != nil {
		return nil, result.Error
	}
	fmt.Println(source.ID)
	return source, nil
}

func (*repo_sqlite) FindAll() ([]data.Source, error) {
	var sources []data.Source
	db_sqlite.Find(&sources)
	if result := db_sqlite.Find(&sources); result.Error != nil {
		return nil, result.Error
	}
	return sources, nil
}

func (*repo_sqlite) Update(source *data.Source) (*data.Source, error) {
	if result := db_sqlite.Save(&source); result.Error != nil {
		return nil, result.Error
	}
	return source, nil
}

func (*repo_sqlite) Find(id string) (*data.Source, error) {
	var source data.Source
	if result := db_sqlite.First(&source, id); result.Error != nil {
		return nil, result.Error
	}
	return &source, nil
}

func (*repo_sqlite) Delete(id string) (bool, error) {
	if err := db_sqlite.Delete(&data.Source{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (*repo_sqlite) FindAllSourceTypes() ([]data.SourceType, error) {
	var results []data.SourceType
	db_sqlite.Find(&results)
	if result := db_sqlite.Find(&results); result.Error != nil {
		return nil, result.Error
	}
	return results, nil
}
