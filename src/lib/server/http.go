package server

import (
	"net/http"
	"fmt"
	"log"
	"time"
	"reflect"
)

var stop chan bool = make(chan bool,1)

type Http_Server struct {

	Routes  map[string]interface{}
	Methods map[string]string
}

type help_handle func(w http.ResponseWriter, r *http.Request)

func init()  {


}

//通过外反射传入
func (self *Http_Server)SetObject(p map[string]interface{})  {
	self.Routes = p
}

//反射调用控制器里面的方法
func (self *Http_Server) RunMethod(handler interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){

		//由于路由全部转发到/  进入r开始处理url

		//如果在方法字典存在 则获取value，此处是{"GET":"Get"}
		params := []reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)}
		method, ok := self.Methods[r.Method]
		if ok {
			f := reflect.ValueOf(handler).MethodByName(method)
			if f.IsValid() {
				f.Call(params)
			}
		}
	}
}
func (self *Http_Server)Run()  {
	self.Methods = make(map[string]string)
	len_methods := len(self.Methods)
	if len_methods<=0{
		self.Methods = map[string]string{
			"GET" : "GET",
			"POST" : "POST",
		}
	}
	//服务器1
	//192.168.14.253:8888
	//目录 /css/  /images/
	//服务器2
	//192.168.14.253:8889
	//目录 /css2/  /images2/

	//http.HandleFunc("/admin/", adminHandler)
	//http.HandleFunc("/login/",loginHandler)
	//http.HandleFunc("/ajax/",ajaxHandler)
	ports := []string{
		":8888",
		":8889",
	}
	for _,v:=range ports{
		mux := http.NewServeMux()

		for k,v:=range self.Routes{
			mux.HandleFunc(k, self.RunMethod(v))
			//http.HandleFunc(k, self.RunMethod(v))
		}
		server := &http.Server{
			Addr: v,
			ReadTimeout: 60 * time.Second,
			WriteTimeout: 60 * time.Second,
			Handler: mux,
		}

		stop <- true
		go func(server *http.Server) {
			if _,ok := <-stop;!ok {
				fmt.Println("Http Server Stop.")
			}else{
				fmt.Println(server)
				err := server.ListenAndServe()
				if err != nil {
					log.Fatal("ListenAndServe: ", err)
				}
			}
		}(server)

		//发送关闭操作
		//close(stop)
	}

	fmt.Println(self.Routes)
	//http.HandleFunc("/", cbc.Test)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.ListenAndServe(":8001", nil)
}

func (h Http_Server)Stop()  {
	stop<-true
	close(stop)
}

func initContrller()  {
	
}


//存入多个静态路径
func set_static_dir(s []string)  {
	l:=len(s)
	for i:=0; i<=l; i++ {
		http.Handle(s[i], http.FileServer(http.Dir("template")))
	}
}

// w表示response对象，返回给客户端的内容都在对象里处理
// r表示客户端请求对象，包含了请求头，请求参数等等
func index(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("User-Agent", "myClient")
	// 往w里写入内容，就会在浏览器里输出
	fmt.Fprintf(w, "Hello golang http!")

	fmt.Println(r.RequestURI)
}