package database

import "github.com/yogesh-p-thakare3110/model"

type IDB interface {
	CreateBook(book model.Book) (int64 error)
	GetAllBooks() ([]model.Book, error)
	GetBookById(Id int64) error
	DeleteBook(Id int64) error
}
