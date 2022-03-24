package usecases

import (
	"github.com/renatodalitba/books-solid-go-lang/modules/books/entities"
	"github.com/renatodalitba/books-solid-go-lang/modules/books/repositories"
)

type usecase struct{}

type CreateBookUseCase interface {
	CreateBook(book *entities.Book) error
}

func NewCreateBookUseCase() CreateBookUseCase {
	return &usecase{}
}

var (
	repo repositories.IBookRepository = repositories.NewPostgresRepository()
)

func (*usecase) CreateBook(book *entities.Book) error {
	err := repo.CreateBook(book)
	if err != nil {
		return err
	}

	return nil
}
