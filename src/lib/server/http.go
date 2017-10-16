package server

import (
	"net/http"
	"fmt"
	"log"
	"time"
	"reflect"
	"lib/config"
	"sync"
	"net/url"
)
var stop chan bool = make(chan bool,1)

type Http_Server struct {
	mutex sync.RWMutex
	Routes  map[string]interface{}
	Methods map[string]string
}

type Http_Server_Interface interface {
	 NewServer()
}

func (own *Http_Server)NewServer()  *Http_Server{
	return own
}

type Http_Server_Help struct {

}

func NewHttpServerHelp() *Http_Server_Help { return new(Http_Server_Help) }
func (own *Http_Server_Help)test()  {
	fmt.Println("ok")
}

func NewHttpServer()*Http_Server {
	return new(Http_Server)
}

func HttpFunc(next http.Handler,i interface{}) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log.Println("before hook1")
		a:=NewHttpServer()
		a.DealHttpMethod(i,w,r)
		next.ServeHTTP(w, r)
		//log.Println("after hook1")
	})
}

//传入
func (self *Http_Server)SetObject(p map[string]interface{})  {
	self.Routes = p
}

//反射调用控制器里面的方法
func (self *Http_Server) RunMethod(handler interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		//fmt.Println(handler)
		//处理HTTP METHOD
		self.DealHttpMethod(handler,w,r)
	}
}

func  (self *Http_Server)DealHttpMethod(handler_method interface{},w http.ResponseWriter, r *http.Request)  {

	//如果在方法字典存在 则获取value，此处是{"GET":"Get"}
	params := []reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)}
	//通过m来改变，尝试获取地址参数m,如m不为空，反射到此方法
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		fmt.Fprintf(w,err.Error())
	}
	var method  string
	if  len(queryForm["m"]) >0{
		method  = queryForm["m"][0]
	}else{
		method = r.Method
	}
	f := reflect.ValueOf(handler_method).MethodByName(method)
	if f.IsValid() {
		f.Call(params)
	}
}



func (self *Http_Server)Run()  {
	//self.Methods = make(map[string]string)
	//len_methods := len(self.Methods)
	//if len_methods<=0{
	//	self.Methods = map[string]string{
	//		"GET" : "GET",
	//		"POST" : "POST",
	//	}
	//}
	//http服务器配置
	//fmt.Println(config.HTTP_SERVERS)
	for _,v:=range config.HTTP_SERVERS{
		mux := http.NewServeMux()
		//路由映射
		for k,v:=range self.Routes{
			//fmt.Println(v)
			_,ok :=  v.(http.Handler)
			if ok{
				mux.HandleFunc(k, v.(http.HandlerFunc))
			}else{
				mux.HandleFunc(k, self.RunMethod(v))  //
			}
			//mux.Handle(k,v["struct"].(http.Handler))
			//http.HandleFunc(k, self.RunMethod(v))
		}
		//如需作为文件服务器 config.HTTP_SERVERS 对应的Static 不为nil
		if v["Static"]!=nil{
			mux.Handle(v["Static"].(string), http.StripPrefix(v["Static"].(string), http.FileServer(http.Dir("."+v["Static"].(string)))))
		}
		server := &http.Server{
			Addr:v["Addr"].(string) ,
			ReadTimeout: v["ReadTimeout"].(time.Duration),
			WriteTimeout:  v["WriteTimeout"].(time.Duration),
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