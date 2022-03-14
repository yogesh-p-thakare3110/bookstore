package sql

import (
	"database/sql"

	"github.com/yogesh-p-thakare3110/bookstore/config"
	"github.com/yogesh-p-thakare3110/bookstore/database"
	"github.com/yogesh-p-thakare3110/bookstore/model"
)

type db struct {
	db *sql.DB
}

func NewDB() database.IDB {
	return &db{
		db: config.Connect(),
	}
}

func (b *db) GetBook() ([]model.Book, error) {
	var book model.Book
	var arrBook []model.Book
	rows, err := b.db.Query("SELECT name, number, publication FROM book")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&book.Name, &book.Number, &book.Publication)
		if err != nil {
			return nil, err
		}
		arrBook = append(arrBook, book)
	}
	return arrBook, nil
}

func (b *db) InsertBook(book model.Book) (int64, error) {

	res, err := b.db.Exec("INSERT INTO book(name, publication) VALUES(?, ?)", book.Name, book.Publication)
	if err != nil {
		return 0, err
	}
	number, err := res.LastInsertId()
	return number, err
}

func (b *db) UpdateBook(book model.Book) (int64, error) {
	res, err := b.db.Exec("Update book SET name=?, publication=?, WHERE number=?", book.Number, book.Name, book.Publication)
	if err != nil {
		return 0, err
	}
	number, err := res.LastInsertId()
	return number, err
}

func (b *db) DeleteBook(number int64) error {
	_, err := b.db.Exec("DELETE FROM book WHERE number = ?", number)
	if err != nil {
		return err
	}
	return err
}
