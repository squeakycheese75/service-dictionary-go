package repository

import "github.com/squeakycheese75/service-dictionary-go/api/data"

type SourceRepository interface {
	Save(source *data.Source) (*data.Source, error)
	FindAll() ([]data.Source, error)
	Update(source *data.Source) (*data.Source, error)
	Find(id string) (*data.Source, error)
	Delete(id string) (bool, error)
}
