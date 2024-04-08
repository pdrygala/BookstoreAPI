package models

import (
	"database/sql"
	"fmt"

	"pdrygala.com/go-bookstore-api/pkg/config"
)

var db *sql.DB

type Book struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func (b *Book) CreateBook() (*Book, error) {
	stmt, err := db.Prepare("INSERT INTO books(name, author, publication) VALUES(?,?,?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(b.Name, b.Author, b.Publication)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	b.ID = id

	return b, nil
}

func GetBookAllBooks() ([]Book, error) {
	rows, err := db.Query("SELECT id, name, author, publication FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Publication); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func GetBookById(ID int64) (*Book, error) {
	var book Book
	err := db.QueryRow("SELECT id, name, author, publication FROM books WHERE id=?", ID).Scan(&book.ID, &book.Name, &book.Author, &book.Publication)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (b *Book) UpdateBook(ID int64) (*Book, error) {
	stmt, err := db.Prepare("UPDATE books SET name=?, author=?, publication=? WHERE id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(b.Name, b.Author, b.Publication, ID)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func DeleteBook(ID int64) error {
	_, err := db.Exec("DELETE FROM books WHERE id=?", ID)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
