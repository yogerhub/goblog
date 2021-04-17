package routes

import (
	"github.com/gorilla/mux"
	"goblog/app/http/controllers/api"
)

// RegisterApiRoutes 注册网页相关路由
func RegisterApiRoutes(r *mux.Router) {

	apiRoute := r.PathPrefix("/api").Subrouter()

	ac := new(api.ArticlesApiController)
	apiRoute.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("api.articles.show")
	apiRoute.HandleFunc("/articles", ac.Index).Methods("GET").Name("api.articles.index")

	cg := new(api.CategoriesApiController)
	apiRoute.HandleFunc("/categories/{id:[0-9]+}", cg.Show).Methods("GET").Name("api.categories.show")
	apiRoute.HandleFunc("/categories", cg.Index).Methods("GET").Name("api.categories.index")

}
