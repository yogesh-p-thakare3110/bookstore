package config

import (
	"database/sql"
)

func Connect() *sql.DB {
	dbUser := "root"
	dbPass := "India@2000"
	dbName := "yogeshdb"
	dbDriver := "mysql"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
