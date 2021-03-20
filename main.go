package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"net/url"
	"strings"
)

var router = mux.NewRouter()

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Home</h1>")
	fmt.Fprint(w, "请求路径："+r.URL.Path)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>About</h1>")
	fmt.Fprint(w, "请求路径："+r.URL.Path)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Not Found")
}

func articlesShowHandler(w http.ResponseWriter, r *http.Request) {
	//将 URL 路径参数解析为键值对应的 Map
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprint(w, "文章ID:"+id)
}

func articlesIndexHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "index Page")
}

func articlesCreateHandler(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <title>创建文章 —— 我的技术博客</title>
</head>
<body>
    <form action="%s?test=test" method="post">
        <p><input type="text" name="title"></p>
        <p><textarea name="body" cols="30" rows="10"></textarea></p>
        <p><button type="submit">提交</button></p>
    </form>
</body>
</html>`
	storeURL, _ := router.Get("articles.store").URL()
	fmt.Fprintf(w, html, storeURL)
}

type ArticlesFormData struct {
	Title, Body string
	URL         *url.URL
	Errors      map[string]string
}

func articlesStoreHandler(w http.ResponseWriter, r *http.Request) {

	title := r.PostFormValue("title")
	body := r.PostFormValue("body")

	errors := make(map[string]string)

	//validate
	if title == "" {
		errors["title"] = "标题不能为空"
	} else if len(title) < 3 || len(title) > 40 {
		errors["title"] = "标题的长度需要在3到40之间"
	}

	if body == "" {
		errors["body"] = "标题不能为空"
	} else if len(body) < 10 {
		errors["body"] = "内容长度必须大于10"
	}

	//检查是否错误
	if len(errors) == 0 {
		fmt.Fprint(w, "验证通过！<br>")
		fmt.Fprintf(w, "title的值为：%v <br>", title)
		fmt.Fprintf(w, "title的长度为：%v <br>", len(title))
		fmt.Fprintf(w, "body的值为：%v <br>", body)
		fmt.Fprintf(w, "body的长度为：%v <br>", len(body))
	} else {
		html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <title>创建文章 —— 我的技术博客</title>
    <style type="text/css">.error {color: red;}</style>
</head>
<body>
    <form action="{{ .URL }}" method="post">
        <p><input type="text" name="title" value="{{ .Title }}"></p>
        {{ with .Errors.title }}
        <p class="error">{{ . }}</p>
        {{ end }}
        <p><textarea name="body" cols="30" rows="10">{{ .Body }}</textarea></p>
        {{ with .Errors.body }}
        <p class="error">{{ . }}</p>
        {{ end }}
        <p><button type="submit">提交</button></p>
    </form>
</body>
</html>
`
		storeURL,_ := router.Get("articles.store").URL()
		data := ArticlesFormData{
			Title: title,
			Body: body,
			URL: storeURL,
			Errors: errors,
		}
		tmpl,err :=template.New("create-form").Parse(html)
		if err !=nil{
			panic(err)
		}

		tmpl.Execute(w,data)
	}

}

func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//设置header头
		w.Header().Set("Content-Type", "text/html;charset=utf-8")

		//继续处理请求
		next.ServeHTTP(w, r)
	})
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	router.HandleFunc("/", HomeHandler).Methods("GET").Name("home")
	router.HandleFunc("/about", AboutHandler).Methods("GET").Name("about")
	router.HandleFunc("/articles/{id:[0-9]+}", articlesShowHandler).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles", articlesIndexHandle).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles/create", articlesCreateHandler).Methods("GET").Name("articles.create")
	router.HandleFunc("/articles", articlesStoreHandler).Methods("POST").Name("articles.store")

	//自定义404页面
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	//中间件：强制内容类型为html
	router.Use(forceHTMLMiddleware)

	homeURL, _ := router.Get("home").URL()
	fmt.Println("homeURL:", homeURL)
	articleURL, _ := router.Get("articles.show").URL("id", "1")
	fmt.Println("articleURL:", articleURL)

	http.ListenAndServe(":3000", removeTrailingSlash(router))
}
