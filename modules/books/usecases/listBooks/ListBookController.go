package usecases

import (
	"encoding/json"
	"net/http"
)

type controller struct{}

type ListBookController interface {
	Books(w http.ResponseWriter, r *http.Request)
}

func NewBookController() ListBookController {
	return &controller{}
}

var (
	listBookUseCase = NewListBooksUseCase()
)

func (*controller) Books(w http.ResponseWriter, r *http.Request) {
	w.Header().Set((`Content-Type`), `application/json`)

	books, err := listBookUseCase.Books()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)

}
