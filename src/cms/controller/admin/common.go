package admin

import (
	"net/http"
	"encoding/json"
	"lib/o-jwt-go"
	"lib/o-jwt-go/request"
	"fmt"
	"cms/model"
	"log"
)

//jwt key
const (
	ConstJwtSecretKey  = "ming123jew!@#$%^&*()"
	ConstSessionAdminLoginFlag  = "admin_login_info"
	ConstSessionAdminLoginFlagValues  = "userinfo"

	ConstSessionOptionsMaxAge =  86400 * 1 //1天
	ConstSessionOptionsPath =  "/"
	ConstSessionOptionsHttpOnly = true

	ConstTemplateSysdataPowerBy = "MingGo  2017"

	ConstTemplateOptionsDirectory = "./src/cms/views/admin/"  //模板对应目录
	ConstTemplateOptionsCharset = "UTF-8" //页面编码
	ConstTemplateOptionsHTMLContentType = "text/html"
	ConstTemplateOptionsExtension = ".html" //模板扩展名
)
var(
	ConstHttpHost = "http://192.168.14.253:8888/"
	ConstTemplateSysdataStatic = "http://192.168.14.253:8001/static/admin/"
	ConstTemplateSysdataPostUrl = ConstHttpHost
)

type JwtToken struct {
	State bool	 `json:"state"`
	Token string `json:"token"`
	Message string `json:"message"`
}

//基于session认证对象
var AuthorizationSession = &SessionIsLoginer{}
//基于jwt-access_token认证对象
var AuthorizationJwt = &JwtIsLoginer{}
var LoginUserInfo = &model.AdminUser{}

//是否登录接口
type IsLoginer interface {
	IsLogin(*http.Handler)http.Handler
}
//session方式
type SessionIsLoginer []string
//根据session进行验证
func (own SessionIsLoginer)IsLogin(next http.HandlerFunc,gourl ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//验证页面 防止缓存
		w.Header().Set("Cache-Control","no-cache")
		w.Header().Set("Pragma","no-cache")
		w.Header().Set("Expires", "0")
		session, _ := Session.Get(r,ConstSessionAdminLoginFlag)
		userinfo_json := session.Values[ConstSessionAdminLoginFlagValues]
		if userinfo_json!= nil{
			json.Unmarshal(userinfo_json.([]byte),&LoginUserInfo)//json 传到 struct
			//log.Println(LoginUserInfo)
			next(w,r)
		}else{
			if gourl[0]!=""{
				//fmt.Fprintf(w,"Unauthorized access to this resource. Please login system.\n")
				http.Redirect(w,r,gourl[0],301)
			}else{
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Unauthorized access to this resource. \nMessage: login failed.")
			}

		}
	})
}
//jwt方式
type JwtIsLoginer []string
//根据地址参数进行验证access_token
func (own JwtIsLoginer)IsLogin(next http.HandlerFunc) http.HandlerFunc {
	var token *jwt.Token
	var err error
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err = request.ParseFromRequest(r, request.ArgumentExtractor{"access_token"}, func(token *jwt.Token) (interface{}, error) {return []byte(ConstJwtSecretKey), nil})
		if err == nil {
			if token.Valid {
				//fmt.Fprint(w, "Token is true.")
				j,_ := json.Marshal(token.Claims.(jwt.MapClaims)[ConstSessionAdminLoginFlag]) //map 转 json
				json.Unmarshal(j,&LoginUserInfo)//json 传到 struct
				next(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Token is not valid.\nError:",err)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized access to this resource. \nError: ",err)
		}
	})
}

//返回JSON
func ReturnJsonResponse(response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func End()  {
	log.Println("end")
}