package routes

import (
	"github.com/gorilla/mux"
	"github.com/theaveas/gCRUD/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/new", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.DeleteBookById).Methods("DELETE")
	router.HandleFunc("/books/{bookId}", controllers.UpdateBookById).Methods("PUT")
}
