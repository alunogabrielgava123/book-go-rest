package infra

import (
	sql "database/sql"
	"log"
	_ "github.com/lib/pq"
)

const connStr = "user=postgres  host=localhost port=5432 dbname=gabrielgavapinheiro sslmode=disable"

func ConfigPgSql() *sql.DB {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
    
	return db

}
