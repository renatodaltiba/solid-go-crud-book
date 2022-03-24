package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	routerDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	routerDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) SERVER(port string) {
	fmt.Println("Serving on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, routerDispatcher))
}
