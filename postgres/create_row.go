package postgres

import "database/sql"

// SQLInsert interface handles inserting elements into the respective tables
type SQLInsert interface {
	insertRow(*sql.DB) bool
}

func Insert(si SQLInsert) bool {
	db := openDBConnection()
	ok := si.insertRow(db)
	if !ok {
		return false
	}
	return true
}
