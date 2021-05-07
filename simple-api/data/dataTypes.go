package dataTypes

import "gorm.io/gorm"

type Source struct{
	gorm.Model
	Name string `json:"name"`
    Desc string `json:"desc"`
	SourceTypeID int `json:"source_type_id"`
	Endpoint string `json:"endpoint"`
}

type SourceType struct{
	ID   int
  	Name string
}

// type Product struct {
// 	gorm.Model
// 	Code  string `json:"code"`
// 	Price uint `json:"price"`
//   }
  

// var Sources []Source
// var Products []Product
