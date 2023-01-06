package main
import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/smkatash/golang_projects/bookstore_mysql/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.BookstoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
