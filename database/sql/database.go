package sql

import (
	"database/sql"

	"github.com/yogesh-p-thakare3110/config"
	"github.com/yogesh-p-thakare3110/database"
)

type Book struct {
	book *sql.DB
}

func NewDB() database.IDB {
	return &Book{
		Book: config.Connect(),
	}
}

func (b *Book) CreateBook() *Book {
	Book.NewRecord(b)
	Book.Create(&b)
	return b
}

func (b *Book) GetAllBooks() []Book {
	var Books []Book
	b.book.Find(&Books)
	return Books
}

func (b *Book) GetBookById(Id int64) *Book {
	var getBook Book
	db.Where("ID=?", Id).Find((&getBook))
	return &getBook, db
}

func (b *Book) DeleteBook(ID int64) Book {
	var book Book
	Book.Where("ID=?", ID).Delete(book)
	return book
}
