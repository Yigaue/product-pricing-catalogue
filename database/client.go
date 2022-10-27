package database

import (
	"crud-rest-api/entities"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error
func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("can not connect to DB")
	}
	log.Println("Connect to Database...")
}

func Migrate() {
	Instance.AutoMigrate(&entities.Product{})
	log.Println("Database migration completed...")
}