package routes

import (
	"github.com/gorilla/mux"
	"github.com/yogesh-p-thakare3110/controller/book"
	"github.com/yogesh-p-thakare3110/database/sql"
)

func RegisterBookStoreRoutes(router *mux.Router) {
	db := sql.NewDB()

	b := book.NewBook(db)

	router.HandleFunc("/book/", b.CreateBook).Methods("POST")
	router.HandleFunc("/book/", b.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", b.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", b.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", b.DeleteBook).Methods("DELETE")
}
