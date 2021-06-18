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
	UpdateSource(source *data.Source) (*data.Source, error)
	Find(id string) (*data.Source, error)
	Delete(id string) (bool, error)
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
		err := errors.New("Name can't be empty")
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

func (*service) UpdateSource(source *data.Source) (*data.Source, error) {
	return repo.Update(source)
}

func (*service) Find(id string) (*data.Source, error) {
	return repo.Find(id)
}

func (*service) Delete(id string) (bool, error) {
	return repo.Delete(id)
}
