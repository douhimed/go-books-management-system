package routes

import (
	"github.com/douhimed/go-books-management-system/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBookById).Methods("DELETE")
}
