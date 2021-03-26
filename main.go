package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"net/http"
)

var router  *mux.Router
var db *sql.DB

func main() {
	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	homeURL, _ := router.Get("home").URL()
	fmt.Println("homeURL:", homeURL)
	articleURL, _ := router.Get("articles.show").URL("id", "1")
	fmt.Println("articleURL:", articleURL)

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
