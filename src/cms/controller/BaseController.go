package controller

import (
	"net/http"
	"fmt"
	"cms/controller/admin"
)

type BaseController struct {
	w http.ResponseWriter
	r *http.Request
}

func (bc *BaseController)Init() map[string]interface{}  {

	return admin.Route
}

func (bc *BaseController)Test(w http.ResponseWriter, r *http.Request){

	// 往w里写入内容，就会在浏览器里输出
	fmt.Fprintf(w, "路由成功")
}

