package middlewares

import (
	"net/http"
)

// ForceHTML 强制标头返回 HTML 内容类型
func ForceHTML(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		next.ServeHTTP(writer,request)
	})
}
