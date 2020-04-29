package postgres

import "database/sql"

// SQLInsert interface handles inserting elements into the respective tables
type SQLInsert interface {
	insertRow(*sql.DB) bool
}

func (e Element) insertRow(db *sql.DB) bool {
	insertStatement := `INSERT INTO elements ( element ) VALUES ($1)`
	_, err := db.Exec(insertStatement, e.Text)
	if err != nil {
		return false
	}
	return true
}

func (b Book) insertRow(db *sql.DB) bool {
	insertStatement := `INSERT INTO books (name, author, pages) VALUES ($1, $2, $3)`
	_, err := db.Exec(insertStatement, b.Name, b.Author, b.Pages)
	if err != nil {
		return false
	}
	return true
}

func Insert(si SQLInsert) bool {
	db := openDBConnection()
	ok := si.insertRow(db)
	if !ok {
		return false
	}
	return true
}
