package database

import "github.com/yogesh-p-thakare3110/bookstore/model"

type IDB interface {
	GetBook() ([]model.Book, error)
	InsertBook(model.Book) (int64, error)
	UpdateBook(model.Book) (int64, error)
	DeleteBook(number int64) error
}
