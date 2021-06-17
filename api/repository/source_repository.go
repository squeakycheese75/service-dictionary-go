package repository

import "github.com/squeakycheese75/service-dictionary-go/api/data"


type SourceRepository interface {
	Save(source *data.Source) (*data.Source, error)
	FindAll() ([]data.Source, error)
}
