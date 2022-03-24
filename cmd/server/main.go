package main

import (
	router "github.com/renatodalitba/books-solid-go-lang/http"
	createbook "github.com/renatodalitba/books-solid-go-lang/modules/books/usecases/createBook"
	usecases "github.com/renatodalitba/books-solid-go-lang/modules/books/usecases/listBooks"
)

var (
	bookListController               = usecases.NewBookController()
	createBook                       = createbook.NewCreateBookController()
	httpRouter         router.Router = router.NewMuxRouter()
)

func main() {
	const port string = "9000"

	httpRouter.GET("/books", bookListController.Books)
	httpRouter.POST("/books", createBook.CreateBook)

	httpRouter.SERVER(port)
}
