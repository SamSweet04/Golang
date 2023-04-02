package main

import (
	"fmt"
	"log"
	"net/http"
    "github.com/SamSweet04/bookstore/pkg/models"
    "github.com/SamSweet04/bookstore/pkg/utils"
    "github.com/SamSweet04/bookstore/pkg/handler"
    "github.com/SamSweet04/bookstore/pkg/config"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config := getConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = db.AutoMigrate(&Book{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/books", getAllBooks(db)).Methods(http.MethodGet)
	r.HandleFunc("/books", addBook(db)).Methods(http.MethodPost)
	r.HandleFunc("/books/{id}", getBookByID(db)).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", updateBookByID(db)).Methods(http.MethodPut)
	r.HandleFunc("/books/{id}", deleteBookByID(db)).Methods(http.MethodDelete)
	r.HandleFunc("/books/search", searchBooksByTitle(db)).Methods(http.MethodGet)
	r.HandleFunc("/books/sort", getAllBooksSortedByCost(db)).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}