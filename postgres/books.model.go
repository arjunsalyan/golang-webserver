package postgres

import (
	"database/sql"
	"errors"
	"fmt"
)

// Book declares the schema for the table "books"
type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

// GetAllBooks fetches all rows from the "books" table and returns a
// slice []Book
func GetAllBooks() ([]Book, error) {
	db := openDBConnection()
	selectStatement := `SELECT * FROM books ORDER BY id DESC`
	rows, err := db.Query(selectStatement)
	if err != nil {
		fmt.Println(err.Error())
		return []Book{}, err
	}
	books := []Book{}
	bookObj := Book{}

	for rows.Next() {
		var (
			id     int
			name   string
			author string
			pages  int
		)
		err = rows.Scan(&id, &name, &author, &pages)
		if err != nil {
			fmt.Println(err.Error())
		}
		bookObj.ID = id
		bookObj.Name = name
		bookObj.Author = author
		bookObj.Pages = pages
		books = append(books, bookObj)
	}
	defer db.Close()
	return books, err
}

func GetOneBook(id int) (Book, error) {
	var (
		name   string
		author string
		pages  int
	)
	sqlStatement := `SELECT * FROM books WHERE id=$1`
	db := openDBConnection()
	row := db.QueryRow(sqlStatement, id)
	switch err := row.Scan(&id, &name, &author, &pages); err {
	case sql.ErrNoRows:
		return Book{}, errors.New("book does not exist")
	case nil:
		obj := Book{
			ID:     id,
			Name:   name,
			Author: author,
			Pages:  pages,
		}
		return obj, nil
	default:
		return Book{}, errors.New("something went wrong")
	}
}

// Add a single book object to the database
func (b Book) insertRow(db *sql.DB) bool {
	insertStatement := `INSERT INTO books (name, author, pages) VALUES ($1, $2, $3)`
	_, err := db.Exec(insertStatement, b.Name, b.Author, b.Pages)
	if err != nil {
		return false
	}
	return true
}
