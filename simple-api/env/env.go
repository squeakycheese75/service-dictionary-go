package env

import (
	"gorm.io/gorm"
)

type Env struct {
	DB *gorm.DB
	// Logger *log.Logger
}
