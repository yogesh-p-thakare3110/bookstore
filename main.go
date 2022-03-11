package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yogesh-p-thakare3110/bookstore/router"
)

func main() {
	r := router.RegisterBookStoreRoutes()
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
