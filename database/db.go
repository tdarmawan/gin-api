package database

import (
	"fmt"
	"log"
	"os"
	"swaggo-gin-api-basic/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db		*gorm.DB
)

func StartDb() {
	err := godotenv.Load()
	if err != nil {
			log.Fatal("Error loading .env file")
	}
	
	host     := os.Getenv("PGHOST")
	user     := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	dbPort   := os.Getenv("PGPORT")
	dbname   := os.Getenv("PGDATABASE")
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting database : ", err)
	}

	db.Debug().AutoMigrate(models.Car{})
}

func GetDB() *gorm.DB {

	return db
}