package admin

import "fmt"

type RbacI interface {
	CheckPermission()
	SetPermission()
}

type Rbac struct {

}

func (own *Rbac)test() {
	fmt.Println("test rbaac")
}


