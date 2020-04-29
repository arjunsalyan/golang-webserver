package postgres

// Book declares the schema for the table "books"
type Book struct {
	ID     int
	Name   string
	Author string
	Pages  int
}
