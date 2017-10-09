package main

import (

	//"runtime"
	"lib/server"
	"cms/controller/admin"
)


//初始化
func init()  {
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
