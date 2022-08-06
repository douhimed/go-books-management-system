package main

import (
	"log"
	"net/http"

	"github.com/douhimed/go-books-management-system/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":91", r))
}
