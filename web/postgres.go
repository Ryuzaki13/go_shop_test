package web

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var connection *sql.DB

func Connect() bool {
	var e error
	connection, e = sql.Open("postgres",
		`host=127.0.0.1
		port=5432
		user=postgres
		password=1234
		dbname=shop
		sslmode=disable`)
	if e != nil {
		fmt.Println("ERROR:", e)
		return false
	}
	return true
}
