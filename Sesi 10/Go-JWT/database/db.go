package database

import (
	"Go-JWT/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "ilham121"
	dbport   = "5432"
	dbname   = "go-jwt"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host = %s port = %s user = %s password = %s dbname = %s sslmode = disable", host, dbport, user, password, dbname)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error Connecting To Database : ", err)
	}

	fmt.Println("Successfully Connected To Database")
	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
