package service

import (
	"errors"

	"github.com/squeakycheese75/service-dictionary-go/api/data"
	"github.com/squeakycheese75/service-dictionary-go/api/repository"
)


type SourceService interface {
	Validate(source *data.Source) error
	Create(source *data.Source) (*data.Source, error)
	FindAll() ([]data.Source, error)
}

type service struct{}

var (
	repo repository.SourceRepository
)

func NewSourceService(repository repository.SourceRepository) SourceService {
	repo = repository
	return &service{}
}

func (*service) Validate(source *data.Source) error {
	if source == nil {
		err := errors.New("The item is empty")
		return err
	}
	if source.Name == "" {
		err := errors.New("The Name is empty")
		return err
	}
	return nil
}

func (*service) Create(source *data.Source) (*data.Source, error) {
	return repo.Save(source)
}

func (*service) FindAll() ([]data.Source, error) {
	return repo.FindAll()
}
