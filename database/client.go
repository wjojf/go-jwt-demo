package database

import (
	"golang-jwt/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instace *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instace, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}

func Migrate() {
	Instace.AutoMigrate((&models.User{}))
	log.Println("Database Migration Completed!")
}
