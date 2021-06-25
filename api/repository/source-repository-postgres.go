package repository

import (
	"fmt"

	"github.com/squeakycheese75/service-dictionary-go/api/data"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repo_postgres struct{}

var (
	db_postgres *gorm.DB
)

// Constructor
func NewPostgresSourceRepository(dsn string) SourceRepository {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db_postgres = db
	return &repo_postgres{}
}

func (*repo_postgres) Save(source *data.Source) (*data.Source, error) {
	if result := db_postgres.Create(&source); result.Error != nil {
		return nil, result.Error
	}
	fmt.Println(source.ID)
	return source, nil
}

func (*repo_postgres) FindAll() ([]data.Source, error) {
	var sources []data.Source
	db_sqlite.Find(&sources)
	if result := db_postgres.Find(&sources); result.Error != nil {
		return nil, result.Error
	}
	return sources, nil
}

func (*repo_postgres) Update(source *data.Source) (*data.Source, error) {
	if result := db_postgres.Save(&source); result.Error != nil {
		return nil, result.Error
	}
	return source, nil
}

func (*repo_postgres) Find(id string) (*data.Source, error) {
	var source data.Source
	if result := db_postgres.First(&source, id); result.Error != nil {
		return nil, result.Error
	}
	return &source, nil
}

func (*repo_postgres) Delete(id string) (bool, error) {
	if err := db_postgres.Delete(&data.Source{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (*repo_postgres) FindAllSourceTypes() ([]data.SourceType, error) {
	var results []data.SourceType
	db_sqlite.Find(&results)
	if result := db_postgres.Find(&results); result.Error != nil {
		return nil, result.Error
	}
	return results, nil
}
