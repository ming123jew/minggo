package main

import (

	//"runtime"
	"lib/server"
	"cms/controller/admin"
)

//用于保存实例化的结构体对象
var regStruct map[string]interface{}

//注册项目结构体
func init()  {
	regStruct["admin"] = admin.AdminController{}
	regStruct["home"] = admin.AdminController{}
}

//import "./src/server"
func main()  {
	http_server := server.Http_Server{}
	http_server.Run()
	//println(runtime.NumCPU())
	//runtime.GOMAXPROCS(1)
/*	for {
		go fmt.Print(0)
		fmt.Print(1)
		//time.Sleep(1*time.Second)
	}*/

}
