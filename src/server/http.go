package server

import (
	"net/http"
	"fmt"
	"log"
	"time"
)

var stop chan bool

type http_server struct {

}

func (h http_server)Run()  {
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
		mux.HandleFunc("/", index)
		server := &http.Server{
			Addr: v,
			ReadTimeout: 60 * time.Second,
			WriteTimeout: 60 * time.Second,
			Handler: mux,
		}
		stop = make(chan bool)
		go func(server *http.Server) {
			if <-stop{
				log.Fatal("Server stop. ")
			}else{
				err := server.ListenAndServe()
				if err != nil {
					log.Fatal("ListenAndServe: ", err)
				}
			}
		}(server)

		//发送关闭操作
		close(stop)
	}

	http.Handle("/", http.HandlerFunc(index))
	http.HandleFunc("/text", index)
	http.ListenAndServe(":8000", nil)

}

func (h http_server)Stop()  {
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
	// 往w里写入内容，就会在浏览器里输出
	fmt.Fprintf(w, "Hello golang http!")
}