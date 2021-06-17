package data

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// Extend the gorm base class to include a UUID for all tables.
type Base struct {
	gorm.Model
	UUID uuid.UUID `gorm:"type:uuid;"`
}

// Extend the gorm base class to include a UUID for all tables.
func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.UUID = uuid.NewV4()
	return
}

type Source struct {
	Base
	Name         string `json:"name"`
	Desc         string `json:"desc"`
	SourceTypeID int    `json:"source_type_id"`
	Endpoint     string `json:"endpoint"`
}

// type Post struct {
// 	ID    int64  `json:"id"`
// 	Title string `json:"Title"`
// 	Text  string `json:"Text"`
// }


type SourceType struct {
	Base
	Name string `json:"name"`
}

type DataSet struct {
	Base
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	SourceId int    `json:"source_id"`
	Body     string `json:"body"`
}
