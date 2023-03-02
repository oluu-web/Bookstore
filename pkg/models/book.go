package models

import (
	"log"

	"github.com/oluu-web/go-bookstore/pkg/config"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func (b *Book) CreateBook() *Book {
	db := config.Connect()
	defer db.Close()

	stmt, err := db.Query("INSERT into "+config.DB_NAME+" (name, author, publication) VALUES (?,?,?)", b.Name, b.Author, b.Publication)
	if err != nil {
		log.Fatal(err)
	}

	for stmt.Next() {
		err = stmt.Scan(&b.Id, &b.Name, &b.Author, &b.Publication)
		if err != nil {
			panic(err.Error())
		}
	}

	return b
}

func GetAllBooks() []Book {
	db := config.Connect()
	defer db.Close()

	stmt, err := db.Query("SELECT * FROM " + config.DB_NAME)
	if err != nil {
		panic(err.Error())
	}

	var books = []Book{}
	for stmt.Next() {
		var b Book
		err = stmt.Scan(&b.Id, &b.Name, &b.Author, &b.Publication)
		if err != nil {
			panic(err.Error())
		}
		books = append(books, b)
	}
	return books
}

func GetBookById(id int64) Book {
	db := config.Connect()
	defer db.Close()

	var getBook Book
	err := db.QueryRow("SELECT * FROM "+config.DB_NAME+" WHERE id = ?", id).Scan(&getBook.Id, &getBook.Name, &getBook.Author, &getBook.Publication)
	if err != nil {
		panic(err.Error())
	}
	return getBook
}

func DeleteBook(id int64) Book {
	db := config.Connect()
	defer db.Close()

	var b Book
	stmt, err := db.Query("DELETE FROM "+config.DB_NAME+" WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	for stmt.Next() {

		err = stmt.Scan(&b.Id, &b.Name, &b.Author, &b.Publication)
		if err != nil {
			panic(err.Error())
		}
	}
	return b
}
