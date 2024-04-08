package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"pdrygala.com/go-bookstore-api/pkg/models"
	"pdrygala.com/go-bookstore-api/pkg/utils"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	b, _ := createBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks, _ := models.GetBookAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//	r.Get("/book/{bookId}", controllers.GetBookById)

func GetBookById(w http.ResponseWriter, r *http.Request) {
	bookId := chi.URLParam(r, "bookId")
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, _ := models.GetBookById(id)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)

	bookId := chi.URLParam(r, "bookId")
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	updateBook.ID = id
	b, _ := updateBook.UpdateBook(id)
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := chi.URLParam(r, "bookId")
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	//Check if book exists
	_, err = models.GetBookById(id)
	if err != nil {
		// Book does not exist, return a 404 Not Found response
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	// Attempt to delete the book
	err = models.DeleteBook(id)
	if err != nil {
		// Handle deletion error
		http.Error(w, "Error deleting book", http.StatusInternalServerError)
	}

	res, _ := json.Marshal(fmt.Sprintf("Book with ID %s deleted successfully", bookId))
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
