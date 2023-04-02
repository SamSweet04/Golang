package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/SamSweet04/bookstore/config"
	"github.com/SamSweet04/bookstore/routers"
	"github.com/SamSweet04/bookstore/models"
)

func main() {
	db, err := gorm.Open("postgres", config.DBConnectionString())
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Automatically create the table based on the Book model
	db.AutoMigrate(&models.Book{})

	router := mux.NewRouter()
	routers.RegisterBookRoutes(router)

	// Start the server
	port := ":8000"
	fmt.Println("Server is listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
