package admin

import (
	"net/http"
	"fmt"
)

var Rbac = &RbacPermission{}

type RbacPermission struct {
	Permission  map[string][]string // article : article/get  article/post  article?m=add
	UserPermission *UserPermission
}
type UserPermission struct {
	Username 	string
	Roleid		int
	Permission  map[string][]string // article : article/get  article/post  article?m=add
}

type RbacI interface {
	CheckPermission()
	SetPermission()
}

func RbacFunc(next http.HandlerFunc) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log.Println("run rbac")
		//log.Println(r.URL)
		next(w, r)
		//log.Println("end rbac")

	})
}

func (self *UserPermission)CheckPermission()  {
	fmt.Println("ok111111")
}


func (own *RbacPermission)test() {
	fmt.Println("test rbaac")
}


