package db

import (
	"fmt"
	"log"
	"os"

	"github.com/fabioalvarez/bookStore/internal/book/infra/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// ConnectDB connect to Postgres DB
func Connect() *gorm.DB {
	if db != nil {
		return db
	}

	var DSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	conn, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	//Check for Errors in DB
	if err != nil {
		log.Fatalf("Error in connect the DB %v", err)
		return nil
	}

	db = conn
	db.AutoMigrate(repository.BookSchema{})
	return db
}

func DB() *gorm.DB {
	return db
}
