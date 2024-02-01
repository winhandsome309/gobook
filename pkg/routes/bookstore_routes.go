package routes

import (
	"github.com/gorilla/mux"
	"github.com/winhandsome309/gobook/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.MainPage).Methods("GET")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PATCH")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
