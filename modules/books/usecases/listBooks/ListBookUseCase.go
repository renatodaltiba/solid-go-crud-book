package usecases

import (
	"github.com/renatodalitba/books-solid-go-lang/modules/books/entities"
	"github.com/renatodalitba/books-solid-go-lang/modules/books/repositories"
)

type usecase struct{}

type ListBooksUseCase interface {
	Books() ([]entities.Book, error)
}

func NewListBooksUseCase() ListBooksUseCase {
	return &usecase{}
}

var (
	repo repositories.IBookRepository = repositories.NewPostgresRepository()
)

func (u *usecase) Books() ([]entities.Book, error) {
	cc, err := repo.Books()
	if err != nil {
		return nil, err
	}

	return cc, nil
}
