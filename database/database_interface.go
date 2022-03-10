package database

type IDB interface {
	CreateBook() *Book
	GetAllBooks() []Book
	GetBookById(Id int64) *Book
	DeleteBook(ID int64) Book
}
