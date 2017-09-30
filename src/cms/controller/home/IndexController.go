package home
import (

	"net/http"
	"fmt"
)

type IndexController struct {

}

func (ic *IndexController)OK(w http.ResponseWriter, r *http.Request)  {

	//r.ParseForm()
	//fmt.Println(r.Form.Get("v"))
	//fmt.Println(r.Form["v"])
	fmt.Fprintf(w, "home路由成功")
	//fmt.Println("ok")
}