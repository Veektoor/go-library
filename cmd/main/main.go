package main

import (
	"go-library/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/jinzhu/gorm/dialect/mysql"
	// "github.com/stacie/go-library/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhsot: 8080", r))
}
