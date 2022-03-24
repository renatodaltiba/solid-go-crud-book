package main

import (
	router "github.com/renatodalitba/books-solid-go-lang/http"
	usecases "github.com/renatodalitba/books-solid-go-lang/modules/books/usecases/listBooks"
)

var (
	bookListController               = usecases.NewBookController()
	httpRouter         router.Router = router.NewMuxRouter()
)

func main() {
	const port string = "9000"

	httpRouter.GET("/books", bookListController.Books)

	httpRouter.SERVER(port)
}
