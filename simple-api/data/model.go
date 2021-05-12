package data

import (
	"gorm.io/gorm"
)

type Source struct {
	gorm.Model
	Name         string `json:"name"`
	Desc         string `json:"desc"`
	SourceTypeID int    `json:"source_type_id"`
	Endpoint     string `json:"endpoint"`
}

type SourceType struct {
	gorm.Model
	Name string `json:"name"`
}

// type DataSet struct {
// 	gorm.Model
// 	Name string `json:"name"`
// }
