package service

import (
	"testing"

	"github.com/squeakycheese75/service-dictionary-go/api/data"
	"github.com/stretchr/testify/assert"
)

func TestValidateEmptyPost(t *testing.T) {
	testService := NewSourceService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "The item is empty", err.Error())
}

func TestValidateMissingName(t *testing.T){
	source :=  &data.Source{Desc: "Description"}
	testService := NewSourceService(nil)

	err := testService.Validate(source)

	assert.NotNil(t, err)
	assert.Equal(t, "The Name is empty", err.Error())
}