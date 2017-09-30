package admin

import (

	"net/http"
	"fmt"
)


type IndexController struct {

}

func (ic *IndexController)OK(w http.ResponseWriter, r *http.Request)  {

	fmt.Fprintf(w, "admin路由成功")
	fmt.Println("ok")
}