package admin

import (
	"net/http"
	"fmt"
)

//登录
type Login struct {
	username string
	passwrod string
}

func  (own *Login)GET(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "admin_get_login")
}

func  (own *Login)POST(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "admin_post_login")

}

