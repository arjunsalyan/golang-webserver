package postgres

import (
	"database/sql"
	"fmt"

	// github.com/lib/pq is needed to connect to a postgresql db
	_ "github.com/lib/pq"
)

const (
	dbHost = "localhost"
	dbPort = 5432
	dbUser = "postgres"
	dbName = "learning_project"
)

func openDBConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection to the database successfull.")
	return db
}
