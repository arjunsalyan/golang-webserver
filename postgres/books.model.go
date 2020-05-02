package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"learning_project/redis"
	"log"
	"strconv"
)

// Book declares the schema for the table "books"
type Book struct {
	redis.RedisBook
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
	var books []Book
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
	// First check in redis
	val, err := redis.GetBookCache("book_" + strconv.Itoa(id))
	book := Book{val}
	if err != nil {
		log.Println(err.Error())
	} else {
		return book, nil
	}

	// If nothing is found in cache then, get from the database
	sqlStatement := `SELECT * FROM books WHERE id=$1`
	db := openDBConnection()
	row := db.QueryRow(sqlStatement, id)
	switch err := row.Scan(&id, &name, &author, &pages); err {
	case sql.ErrNoRows:
		return Book{}, errors.New("book does not exist")
	case nil:
		book.ID = id
		book.Name = name
		book.Author = author
		book.Pages = pages

		// Since the book was not found in the cache, so set it
		err = book.SetBookCache("book_" + strconv.Itoa(id))
		if err != nil {
			log.Println(err.Error())
		}
		return book, nil
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
