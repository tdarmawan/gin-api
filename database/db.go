package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host			= "localhost"
	user			= "postgres"
	password	= "123456"
	dbPort		= "5433"
	dbName		= "universal"
	db		*gorm.DB
	err		error
)

func StartDb() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting database : ", err)
	}

	db.AutoMigrate()
}

func GetDB() *gorm.DB {

	return db
}