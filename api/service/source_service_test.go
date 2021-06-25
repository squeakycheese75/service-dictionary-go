package service

import (
	"testing"

	"github.com/squeakycheese75/service-dictionary-go/api/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(source *data.Source) (*data.Source, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*data.Source), args.Error(1)
}
func (mock *MockRepository) FindAll() ([]data.Source, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]data.Source), args.Error(1)
}
func (mock *MockRepository) Update(source *data.Source) (*data.Source, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*data.Source), args.Error(1)
}
func (mock *MockRepository) Find(id string) (*data.Source, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*data.Source), args.Error(1)
}
func (mock *MockRepository) Delete(id string) (bool, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(bool), args.Error(1)
}
func (mock *MockRepository) FindAllSourceTypes() ([]data.SourceType, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]data.SourceType), args.Error(1)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewSourceService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "The item is empty", err.Error())
}

func TestValidateMissingName(t *testing.T) {
	source := &data.Source{Desc: "Description", Endpoint: "SOmeEndpoint"}
	testService := NewSourceService(nil)

	err := testService.Validate(source)

	assert.NotNil(t, err)
	assert.Equal(t, "Name can't be empty", err.Error())
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)
	sourceA := data.Source{Desc: "Description", Name: "A", Endpoint: "Endpoint A"}
	sourceB := data.Source{Desc: "Description", Name: "B", Endpoint: "Endpoint B"}

	mockRepo.On("FindAll").Return([]data.Source{sourceA, sourceB}, nil)

	testService := NewSourceService(mockRepo)
	result, _ := testService.FindAll()

	// Mock Assertion
	mockRepo.AssertExpectations(t)

	// Data Assertions
	assert.Equal(t, result[0].Name, "A")
	assert.Equal(t, result[1].Name, "B")
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)
	sourceA := data.Source{Desc: "Description", Name: "A", Endpoint: "Endpoint A", SourceTypeID: 1}

	mockRepo.On("Save").Return(&sourceA, nil)

	testService := NewSourceService(mockRepo)
	result, err := testService.Create(&sourceA)

	// Mock Assertion
	mockRepo.AssertExpectations(t)

	// Data Assertions
	assert.Equal(t, result.Name, "A")
	assert.Equal(t, result.SourceTypeID, 1)
	assert.Nil(t, err)
}

func TestUpdate(t *testing.T) {
	mockRepo := new(MockRepository)
	sourceA := data.Source{Desc: "Description", Name: "A", Endpoint: "Endpoint A"}

	mockRepo.On("Update").Return(&sourceA, nil)

	testService := NewSourceService(mockRepo)
	result, err := testService.UpdateSource(&sourceA)

	// Mock Assertion
	mockRepo.AssertExpectations(t)

	// Data Assertions
	assert.Equal(t, result.Name, "A")
	assert.Equal(t, result.SourceTypeID, 0)
	assert.Nil(t, err)
}

func TestDelete(t *testing.T) {
	mockRepo := new(MockRepository)

	mockRepo.On("Delete").Return(true, nil)

	testService := NewSourceService(mockRepo)
	result, err := testService.Delete("2")

	// Mock Assertion
	mockRepo.AssertExpectations(t)

	// Data Assertions
	assert.Equal(t, result, true)
	assert.Nil(t, err)
}
