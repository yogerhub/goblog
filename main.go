package main

import (
	"fmt"
	"net/http"
	"strings"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html;charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>hello</h1>")
		fmt.Fprint(w, "请求路径："+r.URL.Path)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>not fund</h1>")
		fmt.Fprint(w, "请求路径："+r.URL.Path)
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html;charset=utf-8")
	fmt.Fprint(w, "<h1>About</h1>")
	fmt.Fprint(w, "请求路径："+r.URL.Path)
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", DefaultHandler)
	router.HandleFunc("/about", AboutHandler)
	router.HandleFunc("/articles/", func(writer http.ResponseWriter, request *http.Request) {
		id := strings.SplitN(request.URL.Path, "/", 3)[2]
		fmt.Fprint(writer, "文章 ID："+id)
	})
	router.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprint(w,"GET访问文章列表")
		case "POST":
			fmt.Fprint(w,"POST创建新文章")
		}
	})
	http.ListenAndServe(":3000", router)
}
