package routes


import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/SamSweet04/bookstore/pkg/handler"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func getAllBooks(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var books []Book
		result := db.Find(&books)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(books)
	}
}

func addBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		result := db.Create(&book)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(book)
	}
}

func getBookByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "invalid book ID", http.StatusBadRequest)
			return
		}
		var book Book
		result := db.First(&book, id)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(book)
	}
}

func updateBookByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "invalid book ID", http.StatusBadRequest)
			return
		}
		var book Book
		result := db.First(&book, id)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		err = json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		book.ID = uint(id)
		result = db.Save(&book)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(book)
	}
}

func deleteBookByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
		http.Error(w, "invalid book ID", http.StatusBadRequest)
		return
		}
		var book Book
		result := db.Delete(&Book{}, id)
		if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
		}
		json.NewEncoder(w).Encode(map[string]string{"message": "book deleted"})
		}
		}
		
		func searchBookByTitle(db *gorm.DB) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Query().Get("title")
		if title == "" {
		http.Error(w, "missing title parameter", http.StatusBadRequest)
		return
		}
		var books []Book
		result := db.Where("title LIKE ?", "%"+title+"%").Find(&books)
		if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
		}
		json.NewEncoder(w).Encode(books)
		}
		}
		
		func getBooksSortedByCost(db *gorm.DB) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
		order := r.URL.Query().Get("order")
		var books []Book
		result := db.Order("cost " + order).Find(&books)
		if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
		}
		json.NewEncoder(w).Encode(books)
		}
		}