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

func InitialSetup() {
	fmt.Println("Starting inital migration ..")

	GetDb().AutoMigrate(&Source{})
	GetDb().AutoMigrate(&SourceType{})
	// dataTypes.GetDb().AutoMigrate(&dataTypes.DataSet{})

	// Create
	// dataTypes.GetDb().Create(&dataTypes.SourceType{Name: "SQL"})
	// dataTypes.GetDb().Create(&dataTypes.SourceType{Name: "CSV"})
	// dataTypes.GetDb().Create(&dataTypes.Source{Name: "Some_db", Desc: "some db description", Endpoint: "asdad.asdasd.asdsad.asdasd", SourceTypeID: 1})

	// Read
	// var product dataTypes.Product
	// db.First(&product, 1) // find product with integer primary key
	// db.First(&product, "code = ?", "D42") // find product with code D42

	// Migrate the schema

	fmt.Println("Completed inital migration ..")
}

func GetDb() *gorm.DB {
	if db != nil {
		// fmt.Println("Re-using db ..")
		return db
	}

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// defer db.Close()
	// fmt.Println("Connected to db ..")
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
