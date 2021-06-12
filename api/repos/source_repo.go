package repos

import (
	"fmt"

	"github.com/squeakycheese75/service-dictionary-go/api/data"
	"github.com/squeakycheese75/service-dictionary-go/api/env"
)

var Environment *env.Env

type SourceRepository struct {
	Sources map[string]*data.Source
}

func (u *SourceRepository) GetSources() []data.Source {
	var sources []data.Source
	if result := Environment.DB.Find(&sources); result.Error != nil {
		return nil
	}
	return sources
}

func GetSources2() []data.Source {
	var sources []data.Source
	if result := Environment.DB.Find(&sources); result.Error != nil {
		return nil
	}
	return sources
}

func Hi() bool {
	fmt.Println("Hello")
	return true
}

// ByUsername finds a user by their username.
// func (u *SourceRepository) ByUsername(username string) (*model.User, error) {
// 	if user, ok := u.Users[username]; ok {
// 		return user, nil
// 	}
// 	return nil, model.ErrNotFound
// }

// ByID finds a user by their ID
// func (u *SourceRepository) ByID(ID int) (*data.Source, error) {
// 	for _, source := range u.Sources {
// 		if source.ID == ID {
// 			return source, nil
// 		}
// 	}
// 	return nil, model.ErrNotFound
// }
