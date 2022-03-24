package repositories

import "github.com/renatodalitba/books-solid-go-lang/modules/books/entities"

type IBookRepository interface {
	Books() ([]entities.Book, error)
	CreateBook(book *entities.Book) error
}
