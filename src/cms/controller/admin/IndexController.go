package admin

import (
	"net/http"
	"fmt"

)

type Index struct {

}

func (own *Index)GET(w http.ResponseWriter, r *http.Request )  {
	//fmt.Fprintf(w,"hello.<a href='admin/login?m=out'>退出登录</a>",[]byte(LoginUserInfo.Username))

	Template.Html(w,r,Template.TemplateData,"index","public_menu")
}

func (own *Index)Add(w http.ResponseWriter, r *http.Request)  {

	fmt.Fprintf(w,"add hello")
}



