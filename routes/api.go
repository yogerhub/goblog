package routes

import (
	"github.com/gorilla/mux"
	api2 "goblog/app/http/controllers/api"
)

// RegisterApiRoutes 注册网页相关路由
func RegisterApiRoutes(r *mux.Router) {

	api := r.PathPrefix("/api").Subrouter()

	ac := new(api2.ArticlesApiController)
	api.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("api.articles.show")

	api.HandleFunc("/articles", ac.Index).Methods("GET").Name("api.articles.index")

}
