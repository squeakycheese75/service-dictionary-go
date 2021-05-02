package dataTypes

import "gorm.io/gorm"

type Source struct{
	gorm.Model
	Name string `json:"name"`
    Desc string `json:"desc"`
	// SourceTypeId int `json:"source_type_id"`
	Endpoint string `json:"endpoint"`
}


type Product struct {
	gorm.Model
	Code  string `json:"code"`
	Price uint `json:"price"`
  }
  

var Sources []Source
var Products []Product
