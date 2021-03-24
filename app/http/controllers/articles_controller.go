package controllers

import (
	"database/sql"
	"fmt"
	"goblog/logger"
	"goblog/pkg/route"
	"goblog/types"
	"net/http"
)

// PagesController 处理静态页面
type ArticlesController struct {
}

// Show 文章详情页面
func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	//将 URL 路径参数解析为键值对应的 Map
	// 1. 获取 URL 参数
	id := getRouteVariable("id", r)

	// 2. 读取对应的文章数据
	article, err := getArticleByID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章不存在")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		tmpl, err := template.New("show.tmpl").Funcs(template.FuncMap{
			"RouteName2URL":route.Name2URL,
			"Int64ToString":types.Int64ToString,
		}).ParseFiles("resources/views/articles/show.tmpl")
		logger.LogError(err)
		tmpl.Execute(w, article)
	}
}
