package config

import (
	"database/sql"
)

func connect() *sql.DB {
	dbUser := "root"
	dbPass := "India@2000"
	dbName := "yogeshdb"
	dbDriver := "mysql"

	Book, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return Book
}
