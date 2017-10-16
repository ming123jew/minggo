package admin

import (
	"fmt"
	"net/http"
	"log"
)


type RbacPermission struct {
	Permission  map[string][]string // article : article/get  article/post  article?m=add
}
type UserPermission struct {
	Username 	string
	Roleid		int
	Permission  map[string][]string // article : article/get  article/post  article?m=add
}

type Rbac struct {
	UserPermission *UserPermission
}
type RbacI interface {
	CheckPermission()
	SetPermission()
}

func RbacFunc(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log.Println("run rbac")
		log.Println(r.URL)
		next.ServeHTTP(w, r)
		//log.Println("end rbac")

	})
}
func NewRbac(rbac *Rbac) {
	rbac.UserPermission.CheckPermission()
}

func (self *UserPermission)CheckPermission()  {
	fmt.Println("ok111111")
}


func (own *Rbac)test() {
	fmt.Println("test rbaac")
}


