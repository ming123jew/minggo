package admin

import "fmt"


type RbacPermission struct {
	Permission  map[string][]string // article : article/get  article/post  article?m=add
}
type UserPermission struct {
	Username 	string
	Roleid		int
	Permission  map[string][]string // article : article/get  article/post  article?m=add
}

type Rbac struct {

}
type RbacI interface {
	CheckPermission()
	SetPermission()
}
func NewRbac(rbac *Rbac) {
	rbac.test()
}

func (own *Rbac)test() {
	fmt.Println("test rbaac")
}


