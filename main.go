package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>hello</h1>")
		fmt.Fprint(w, "请求路径："+r.URL.Path)
	}else if r.URL.Path == "/about" {
		fmt.Fprint(w, "<h1>about</h1>")
		fmt.Fprint(w, "请求路径："+r.URL.Path)
	}else {
		fmt.Fprint(w, "<h1>not fund</h1>")
		fmt.Fprint(w, "请求路径："+r.URL.Path)
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
