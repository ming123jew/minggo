package admin

import (
	"fmt"
	"net/http"

)

type AdminController struct {
	Is_Login bool
}

func (ac *AdminController)GET(w http.ResponseWriter, r *http.Request)  {
	// 往w里写入内容，就会在浏览器里输出
	Session.Get(r,"test")
	//fmt.Println(session)
	fmt.Fprintf(w, "路由成功333333333")
}

