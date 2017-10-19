package admin

import (
	"lib/server"
	"net/http"
	//"log"
	"time"
)
var Route  map[string] interface{} = make(map[string] interface{},32)

func init()  {
	//最外层func先执行，层层执行 | 中间间
	//路由默认访问内置结构体对应的GET|POST方法，如需访问结构体的其他方法，可在地址末尾加上参数m=xxx进行选定
	//server.HttpFunc
	//RbacFunc
	//AuthorizationSession.IsLogin|AuthorizationJwt.IsLogin
	Route = map[string] interface{}{
		"/admin/login": server.HttpFunc(DefaultHttpFunc,&LoginController{}) ,
		"/admin/index": AuthorizationSession.IsLogin( RbacFunc( server.HttpFunc( DefaultHttpFunc,&Index{} ) ),"login" ) ,
		"/admin/test": &AdminController{} ,
	}
}
//next会触发此函数
func DefaultHttpFunc(w http.ResponseWriter, r *http.Request){
	//w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	//w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	//设置模板公共信息
	Template.SetTemplateData(struct {
		Host string
		Proto string
		Protocol string
		Time   int64
	}{
		ReturnProtocol(r.Proto)+r.Host+"/",
		r.Proto,
		ReturnProtocol(r.Proto),
		time.Now().Unix(),
	})
	//log.Println(Template.TemplateData)
}
func ReturnProtocol(s string)string  {
	var r string
	switch s {
	case "HTTP/1.1":
		r="http://"
	}
	return r
}
