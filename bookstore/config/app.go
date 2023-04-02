package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func DBConnectionString() string {
	dbname := "book-store"
	dbhost := "localhost"
	dbport := "5432"
	dbuser := "postgres"
	dbpass := "Haker15987"

	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbhost, dbport, dbuser, dbname, dbpass)
}

func NewDBConnection() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", DBConnectionString())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func SetupDB() {
	db, err := NewDBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Enable verbose logging mode for GORM
	db.LogMode(true)

	// Automatically create the table based on the Book model
	db.AutoMigrate(&models.Book{})
}
