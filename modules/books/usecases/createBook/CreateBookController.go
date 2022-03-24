package usecases

import (
	"encoding/json"
	"net/http"

	"github.com/renatodalitba/books-solid-go-lang/modules/books/entities"
)

type controller struct{}

type CreateBookController interface {
	CreateBook(w http.ResponseWriter, r *http.Request)
}

func NewCreateBookController() CreateBookController {
	return &controller{}
}

var (
	createBookUseCase = NewCreateBookUseCase()
)

func (*controller) CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set((`Content-Type`), `application/json`)

	book := &entities.Book{}
	err := json.NewDecoder(r.Body).Decode(book)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = createBookUseCase.CreateBook(book)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
