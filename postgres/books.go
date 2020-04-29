package postgres

import "fmt"

// Book declares the schema for the table "books"
type Book struct {
	ID     int
	Name   string
	Author string
	Pages  int
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
