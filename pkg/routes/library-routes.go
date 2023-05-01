package routes

import (
	"go-library/pkg/controllers"

	"github.com/gorilla/mux"
	// "127.0.0.1/controllers/controllers"
	// "github.com/wpcodevo/golang-gorm-postgres/models"
)

var RegisterBookStoreRoutes = func (router *mux.Router){
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}",controllers.DeleteBook).Methods("DELETE")

}