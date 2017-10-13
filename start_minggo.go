package main

import (

	//"runtime"
	"lib/server"
	"cms/controller/admin"
	"cms/model"
	"fmt"
	"cms/initialize"
)


//初始化
func init()  {

}

//import "./src/server"
func main()  {
	initModel()

	initHttp()

	//println(runtime.NumCPU())
	//runtime.GOMAXPROCS(1)
/*	for {
		go fmt.Print(0)
		fmt.Print(1)
		//time.Sleep(1*time.Second)
	}*/

}

//启动http服务
func initHttp()  {
	http_server := server.Http_Server{}
	http_server.SetObject(admin.Route)
	//mut := reflect.ValueOf(http_server).Elem()
	//mut.FieldByName("Object").SetMapIndex(reflect.ValueOf("admin"), reflect.ValueOf(admin.AdminController{}))
	//params :=  []reflect.Value{reflect.ValueOf(regStruct)}
	//f := reflect.ValueOf(&http_server).MethodByName("SetObject")
	//if f.IsValid() {
	//	f.Call(params)
	//}
	http_server.Run()
}

//加载model
func initModel()  {
	//同步数据库表
	err := initialize.Orm.Sync2(new(model.AdminUser),new(model.Group) )
	if err!=nil{
		fmt.Println("Orm:Sync2:error:",err)
	}
	fmt.Println("ok")
}