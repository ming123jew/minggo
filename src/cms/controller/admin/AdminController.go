package admin

import (
	"fmt"
	"net/http"
)

//集中所有controller
type AdminController struct {
	Is_Login bool

}

func (ac *AdminController)Test(w http.ResponseWriter, r *http.Request)  {
	// 往w里写入内容，就会在浏览器里输出
	fmt.Fprintf(w, "路由成功333333333")
}