package main

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

// Book is a placeholder for book
type Book struct {
	id     int
	name   string
	author string
}

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Println(err)
	}
	// create a table
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS book (id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64), name VARCHAR(64) NULL)")
	if err != nil {
		log.Println("Error in creating table")
	} else {
		log.Println("Successfully created table books!")
	}
	statement.Exec()
	dbOperations(db)
}

func dbOperations(db *sql.DB) {
	// create
	statement, _ := db.Prepare("INSERT INTO book (name, author, isbn) VALUES (?, ?, ?)")
	statement.Exec("A table of two cities", "Charles Dickens", 140430547)
	log.Println("Inserted the book into database!")

	// read
	rows, _ := db.Query("SELECT id, name, author FROM book")
	var tempBook Book
	for rows.Next() {
		rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
		log.Printf("ID:%d, Book:%s, Author:%s\n", tempBook.id, tempBook.name, tempBook.author)
	}

	// update
	statement, _ = db.Prepare("UPDATE book SET name=? WHERE id=?")
	statement.Exec("The Tale of Two Cities", 1)
	log.Println("Successfully updated the book in database!")

	// delete
	statement, _ = db.Prepare("DELETE FROM book WHERE id=?")
	statement.Exec(1)
	log.Println("Successfully deleted the book in database!")
}
