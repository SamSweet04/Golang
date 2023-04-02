package routers

import (
	"github.com/gorilla/mux"
	"bookstore/controller"
)

// SetBookRoutes defines all the book-related routes
func SetBookRoutes(router *mux.Router) *mux.Router {
	// Create a new book
	router.HandleFunc("/books", controller.CreateBook).Methods("POST")

	// Get all books
	router.HandleFunc("/books", controller.GetAllBooks).Methods("GET")

	// Get a book by ID
	router.HandleFunc("/books/{id}", controller.GetBookByID).Methods("GET")

	// Update a book by ID
	router.HandleFunc("/books/{id}", controller.UpdateBook).Methods("PUT")

	// Delete a book by ID
	router.HandleFunc("/books/{id}", controller.DeleteBook).Methods("DELETE")

	// Search books by title
	router.HandleFunc("/books/search", controller.SearchBooksByTitle).Methods("GET")

	// Get a sorted list of books ordered by cost in descending order
	router.HandleFunc("/books/sorted/desc", controller.GetBooksSortedByCostDescending).Methods("GET")

	// Get a sorted list of books ordered by cost in ascending order
	router.HandleFunc("/books/sorted/asc", controller.GetBooksSortedByCostAscending).Methods("GET")

	return router
}
