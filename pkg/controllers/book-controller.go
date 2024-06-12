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
	// Parse request body to create a new book
	createBook := &models.Book{}
	if err := utils.ParseBody(r, createBook); err != nil {
		// If there's an error parsing the request body, return a bad request response
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Attempt to create the book
	newBook, err := createBook.CreateBook()
	if err != nil {
		// If there's an error creating the book, return an internal server error response
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}

	// Marshal the created book to JSON
	res, err := json.Marshal(newBook)
	if err != nil {
		// If there's an error marshaling the book to JSON, return an internal server error response
		http.Error(w, "Failed to marshal book to JSON", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Write response with the created book data
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	// Attempt to fetch all books
	newBooks, err := models.GetBookAllBooks()
	if err != nil {
		// If there's an error fetching books, return an error response
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}

	// Marshal books to JSON
	res, err := json.Marshal(newBooks)
	if err != nil {
		// If there's an error marshaling books to JSON, return an error response
		http.Error(w, "Failed to marshal books to JSON", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Write response
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	bookId := chi.URLParam(r, "bookId")
	id, err := strconv.ParseInt(bookId, 10, 64) // Use base 10 and 64-bit size
	if err != nil {
		// If there's an error parsing the book ID, return a bad request response
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	bookDetails, err := models.GetBookById(id)
	if err != nil {
		// If there's an error fetching the book details, return a not found response
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Marshal book details to JSON
	res, err := json.Marshal(bookDetails)
	if err != nil {
		// If there's an error marshaling book details to JSON, return an internal server error response
		http.Error(w, "Failed to marshal book details to JSON", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Write response with book details
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Parse request body to get updated book details
	updateBook := &models.Book{}
	if err := utils.ParseBody(r, updateBook); err != nil {
		// If there's an error parsing the request body, return a bad request response
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Parse book ID from URL parameter
	bookId := chi.URLParam(r, "bookId")
	id, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		// If there's an error parsing the book ID, return a bad request response
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	// Check if the book exists
	_, err = models.GetBookById(id)
	if err != nil {
		// Book does not exist, return a 404 Not Found response
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Update the book record
	updatedBook, err := updateBook.UpdateBook(id)
	if err != nil {
		// Handle updating error
		http.Error(w, "Error updating book", http.StatusInternalServerError)
		return
	}

	// Marshal the updated book to JSON
	res, err := json.Marshal(updatedBook)
	if err != nil {
		// If there's an error marshaling book details to JSON, return an internal server error response
		http.Error(w, "Failed to marshal book details to JSON", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Write response with updated book details
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Parse book ID from URL parameter
	bookId := chi.URLParam(r, "bookId")
	id, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		// If there's an error parsing the book ID, return a bad request response
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	// Check if the book exists
	_, err = models.GetBookById(id)
	if err != nil {
		// Book does not exist, return a 404 Not Found response
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Attempt to delete the record
	err = models.DeleteBook(id)
	if err != nil {
		// Handle deletion error
		http.Error(w, "Error deleting book", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	res, _ := json.Marshal(fmt.Sprintf("Book with ID %s deleted successfully", bookId))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	w.Write(res)
}
