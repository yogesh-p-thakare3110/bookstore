package router

import (
	"github.com/gorilla/mux"
	"github.com/yogesh-p-thakare3110/bookstore/controller/book"
	"github.com/yogesh-p-thakare3110/bookstore/database/sql"
)

func RegisterBookStoreRoutes() *mux.Router {
	db := sql.NewDB()

	b := book.NewBook(db)

	router := mux.NewRouter()
	router.HandleFunc("/getBook/", b.AllBook()).Methods("GET")
	router.HandleFunc("/insertBook/", b.InsertBook()).Methods("POST")
	router.HandleFunc("/updateBook/", b.UpdateBook()).Methods("PUT")
	router.HandleFunc("/deleteBook/", b.DeleteBook()).Methods("DELETE")

	return router
}
