package repository

import (
	"github.com/squeakycheese75/service-dictionary-go/api/data"
	env "github.com/squeakycheese75/service-dictionary-go/api/env"
)

type repo struct{}

var (
	environment env.Env
)

//NewSqlLiteSourceRepository creates a new repo
func NewSqlLiteSourceRepository(env env.Env) SourceRepository {
	environment = env
	return &repo{}
}


func (*repo) Save(source *data.Source) (*data.Source, error) {
	if result := environment.DB.Create(&data.Source{Name: source.Name, Desc: source.Desc, Endpoint: source.Endpoint, SourceTypeID: source.SourceTypeID}); result.Error != nil {
		return nil, result.Error
	}
	return source, nil
}

func (*repo) FindAll() ([]data.Source, error) {
	var sources []data.Source
	environment.DB.Find(&sources)
	if result := environment.DB.Find(&sources); result.Error != nil {
		return nil, result.Error
	}
	return sources, nil
}
