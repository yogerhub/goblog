package api

import (
	"encoding/json"
	"fmt"
	"goblog/app/models/category"
	"goblog/pkg/route"
	"net/http"
)

// CategoriesApiController 处理静态页面
type CategoriesApiController struct {
	BaseApiController
}

// Show 文章详情页面
func (ac *CategoriesApiController) Show(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数

	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	_category, err := category.Get(id)

	// 3. 如果出现错误
	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {
		// 4. 读取成功，显示文章
		js,_ := json.Marshal(_category)
		fmt.Fprintf(w,"%s", js)
	}
}

func (ac *CategoriesApiController) Index(w http.ResponseWriter, r *http.Request) {

	//获取结果集
	categories, err := category.All()

	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {
		// ---  2. 加载模板 ---
		js,_ := json.Marshal(categories)
		fmt.Fprintf(w,"%s", js)
	}

}
