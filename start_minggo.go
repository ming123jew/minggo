package main

import (

	//"runtime"
	"lib/server"
	"cms/controller/admin"
	"reflect"
)

//用于保存实例化的结构体对象
var regStruct map[string]interface{}

//注册项目结构体
func init()  {
	regStruct = make(map[string]interface{})
	regStruct["admin"] = &admin.AdminController{}
	regStruct["home"] = &admin.AdminController{}
}



//import "./src/server"
func main()  {
	initHttp()
	//println(runtime.NumCPU())
	//runtime.GOMAXPROCS(1)
/*	for {
		go fmt.Print(0)
		fmt.Print(1)
		//time.Sleep(1*time.Second)
	}*/

}

func initHttp()  {
	http_server := server.Http_Server{}


	//mut := reflect.ValueOf(http_server).Elem()
	//mut.FieldByName("Object").SetMapIndex(reflect.ValueOf("admin"), reflect.ValueOf(admin.AdminController{}))
	params :=  []reflect.Value{reflect.ValueOf(regStruct)}
	f := reflect.ValueOf(&http_server).MethodByName("SetObject")
	if f.IsValid() {
		f.Call(params)
	}

	http_server.Run()


}
