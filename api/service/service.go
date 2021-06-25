package service

import (
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
	ValidateSourceType(sourceTyoe *data.SourceType) error
	FindAllSourceTypes() ([]data.SourceType, error)
}

type service struct{}

var (
	repo repository.SourceRepository
)

func NewSourceService(repository repository.SourceRepository) SourceService {
	repo = repository
	return &service{}
}
