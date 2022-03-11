package sql

import (
	"database/sql"

	"github.com/yogesh-p-thakare3110/config"
	"github.com/yogesh-p-thakare3110/database"
	"github.com/yogesh-p-thakare3110/model"
)

type Book struct {
	Book *sql.DB
}

func NewDB() database.IDB {
	return &Book{
		Book: config.Connect(),
	}
}

func (b *Book) GetAllBooks() ([]model.Book, error) {
	var book model.Book
	var arrBook []model.Book
	rows, err := b.Book.Query("SELECT name, author, publication FROM book")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&book.Name, &book.Author, &book.Publication)
		if err != nil {
			return nil, err
		}
		arrBook = append(arrBook, book)
	}
	return arrBook, nil
}

func (b *Book) CreateBook(book model.Book) (int64 error) {
	_, err := b.Book.Exec("INSERT INTO book (name, publication) VALUES(?, ?)", book.Name, book.Publication)
	if err != nil {
		return err
	}
	return err
}

func (b *Book) GetBookById(Id int64) error {
	_, err := b.Book.Exec("SELECT FROM book WHERE Id = ?, Id")
	if err != nil {
		return err
	}
	return err
}

func (b *Book) DeleteBook(Id int64) error {
	_, err := b.Book.Exec("DELETE FROM book WHERE Id = ?, Id")
	if err != nil {
		return err
	}
	return err
}
