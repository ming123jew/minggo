package admin

import "lib/server"
var Route  map[string] interface{} = make(map[string] interface{},32)

func init()  {
	Route = map[string] interface{}{
		"/admin/login": server.HttpFunc(&Login{},&Login{}) ,
		"/admin/index": RbacFunc(server.HttpFunc(&Index{},&Index{})) ,
		"/admin/test": &AdminController{} ,
	}
}
