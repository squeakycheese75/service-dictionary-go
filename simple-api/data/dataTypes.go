package dataTypes

import "gorm.io/gorm"

type Source struct{
	Id string `json:"id"`
	Name string `json:"name"`
    Desc string `json:"desc"`
}


type Product struct {
	gorm.Model
	Code  string `json:"code"`
	Price uint `json:"price"`
  }
  

var Sources []Source
var Products []Product
