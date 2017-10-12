package controller

import (
	"net/http"
	"fmt"
	"lib/sesson"
	"lib/core"
	"cms/initialize"
	"cms/model"
)
var (
	SessionStore = sessions.NewCookieStore([]byte("minggo"))
	VIEW_PATH =  core.GetCurrentDirectory()
)

func init()  {
	//设置session默认参数
	SessionStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 1,//1天
		HttpOnly: true,
	}
}

type BaseController struct {
	w http.ResponseWriter
	r *http.Request
}

func (bc *BaseController)Test(w http.ResponseWriter, r *http.Request){
	// 往w里写入内容，就会在浏览器里输出
	fmt.Fprintf(w, "路由成功")
}

