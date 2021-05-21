package data

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	dsn = "host=localhost user=postgres password=changeme dbname=postgres port=5432 sslmode=disable"
)

var db *gorm.DB

func GetDb() *gorm.DB {
	if db != nil {
		fmt.Println("Re-using db ..")
		return db
	}

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// defer db.Close()
	fmt.Println("Connected to db ..")
	return db
}

func getDbPostgres() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// fmt.Println("Connected to db ..")
	return db
}

func getDbSqlLite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// fmt.Println("Connected to db ..")
	return db
}
