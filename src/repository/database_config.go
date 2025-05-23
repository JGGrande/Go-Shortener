package repository

import (
	"log"
	"shortener/src/model"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB

func Connect() {
	if dbConnection != nil {
		log.Println("Database connection already established")
		return
	}

	dsn := "host=postgres user=postgres password=12345678 dbname=url_shortener port=5432"
	time.Sleep(10 * time.Second)
	// db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	dbConnection = db
}

func RunMigrations() {
	if dbConnection == nil {
		log.Fatal("Database connection not established")
	}

	err := dbConnection.AutoMigrate(&model.Shortener{})

	if err != nil {
		log.Fatal("Failed to run migrations")
	}
}
