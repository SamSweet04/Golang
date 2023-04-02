package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/SamSweet04/bookstore/models"
	"github.com/SamSweet04/bookstore/utils"
)

// CreateBook creates a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	err := json.NewDecoder(r.Body).Decode(book)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	err = book.Create()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, book)
}

// GetAllBooks gets all books
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetAllBooks()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, books)
}

// GetBookByID gets a book by ID
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	book, err := models.GetBookByID(uint(id))
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Book not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, book)
}

// UpdateBook updates a book by ID
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	book := &models.Book{}
	err = json.NewDecoder(r.Body).Decode(book)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	book.ID = uint(id)
	err = book.Update()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, book)
}

// DeleteBook deletes a book by ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	err = models.DeleteBook(uint(id))
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// SearchBooksByTitle searches for books by title
func SearchBooksByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	if title == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Missing title parameter")
		return
	}

	books, err := models.SearchBooksByTitle(title)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, books)
}
// GetBooksSortedByCostDescending retrieves a sorted list of books ordered by cost in descending order
func GetBooksSortedByCostDescending(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetBooksSortedByCost("desc")
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, books)
}

// GetBooksSortedByCostAscending retrieves a sorted list of books ordered by cost in ascending order
func GetBooksSortedByCostAscending(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetBooksSortedByCost("asc")
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, books)
}
