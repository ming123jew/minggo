package admin

import (
	"net/http"
	"fmt"
	//"lib/mustache"
	//"cms/initialize"
	"encoding/json"
	"cms/model"
	"lib/o-jwt-go"
	"time"
	"lib/o-jwt-go/request"
	"log"
)


//登录
type LoginController struct {

}
//验证
type Authentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ReturnResponser struct {
	State bool `json:"state"`
	Message string `json:"message"`
	Other interface{}
}

//登录页面显示
func  (own *LoginController)GET(w http.ResponseWriter, r *http.Request)  {
	//own.mutex.Lock()
	//defer own.mutex.Unlock()
	//登录页
	//testing
	//session, _ := Session.Get(r,"test")
	// Set some session values.
	//session.Values["foo"] = "bar"
	//session.Values[42] = 43
	// Save it before we write to the response/return from the handler.
	//session.Save(r, w)
	//sql := "select * from m_admin_user"
	//results, err := initialize.Orm.Query(sql)
	//fmt.Println(results,err)
	//fmt.Println(session.Options)
	/*
	t,error:=template.ParseFiles("./src/cms/views/admin/login.html")//New("login.html")
	if error!=nil{
		fmt.Fprintf(w,error.Error())
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	error = t.Execute(w, TemplateData)
	if error != nil{
		fmt.Fprintf(w,error.Error())
		return
	}*/
	//TemplateData["title"] ="OK"
	type Users struct {
		Username string
	}
	users := []Users{
		{Username:"ming123jew"},
		{Username:"ming"},
	}
	type Mo struct {
		Ua string
	}
	mo := []Mo{
		{"mmm"},
	}

	Template.SetTemplateData(
		struct {
		Users []Users
		Mo []Mo
		}{
		users,
		mo,
		},
	)
	Template.Html(w,r,Template.TemplateData,"login",)
	//s := mustache.RenderFileInLayout("./src/cms/views/admin/login.html", "./src/cms/views/admin/layout.html.mustache", nil)
	//fmt.Fprintf(w,s)
}

//登录操作 | session | 返回json格式
func (own *LoginController)POST(w http.ResponseWriter, r *http.Request)  {
	var response ReturnResponser
	//登录前检测
	session, _ := Session.Get(r,ConstSessionAdminLoginFlag)
	//log.Println("login:",session)
	if session.Values[ConstSessionAdminLoginFlagValues]!=nil{
		//log.Println("login success.")
		//http.Redirect(w,r,"index",301)
		response = ReturnResponser{true,"Userinfo from session.",nil}
		ReturnJsonResponse(response,w)
		return
	}
	var user Authentication
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		//w.WriteHeader(http.StatusForbidden)
		//fmt.Fprint(w, "Error in request:",err)
		log.Println("Error in request: error in username or password. \nmessage:",err)
		response = ReturnResponser{false,"Error in request: error in username or password.",nil}
		return
	}

	// Set some session values.
	//session.Values["foo"] = "bar"
	//session.Values[42] = 43
	// Save it before we write to the response/return from the handler.
	//session.Save(r, w)
	model_AdminUser := new(model.AdminUser)
	b,m,err := model_AdminUser.GetByUsernameAndPassword(user.Username,user.Password)
	//fmt.Println(b,m)
	if b==false{
		response =  ReturnResponser{false,"Error in username or password.",nil}
	}else{
		response = ReturnResponser{true,"success.",nil}
		//验证成功
		m_json,err :=  json.Marshal(m)
		if err != nil {
			log.Println("error in struct to json.  \nmessage:",err)
		}
		session.Values[ConstSessionAdminLoginFlagValues] = m_json
		session.Save(r, w)
	}
	ReturnJsonResponse(response,w)
}

//登录操作 | jwt | 返回json格式
func  (own *LoginController)Jwt(w http.ResponseWriter, r *http.Request)  {
	//登录前检测
	token,ok := own.ValidateLoginToken(w,r)
	switch ok {
	case true:
		//fmt.Println("auth:",ok,"\n token:",token,"\n token claims:",token.Claims.(jwt.MapClaims)[AdminLoginFlag],"\n",token.Signature)
		j,err := json.Marshal(token.Claims.(jwt.MapClaims)[ConstSessionAdminLoginFlag]) //map 转 json
		var user model.AdminUser
		json.Unmarshal(j,&user)//json 传到 struct
		if err!=nil{
			log.Println(err)
		}
	case false:
		var user Authentication
		var response JwtToken
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			//w.WriteHeader(http.StatusForbidden)
			//fmt.Fprint(w, "Error in request:",err)
			response = JwtToken{false,"","Error in request:"}
			return
		}
		//查询数据库
		//fmt.Println(r.Body)
		//fmt.Println(user.Username)
		//fmt.Println(user.Password)
		model_AdminUser := new(model.AdminUser)
		b,m,err := model_AdminUser.GetByUsernameAndPassword(user.Username,user.Password)
		if b==false{
			//fmt.Fprintf(w,"Error in username or password")
			response = JwtToken{false,"","Error in username or password"}

		}else{
			//jwt token 加密操作
		//iss: jwt签发者
		//sub: jwt所面向的用户
		//aud: 接收jwt的一方
		//exp: jwt的过期时间，这个过期时间必须要大于签发时间
		//nbf: 定义在什么时间之前，该jwt都是不可用的.
		//iat: jwt的签发时间
		//jti: jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。
			token := jwt.New(jwt.SigningMethodHS256)
			claims := make(jwt.MapClaims)
			claims["exp"] = time.Now().Add(time.Second * time.Duration(1800)).Unix()//time.Now().Add(time.Hour * time.Duration(1)).Unix()
			claims["iat"] = time.Now().Unix()
			claims[ConstSessionAdminLoginFlag] = m
			token.Claims = claims
			tokenString, err := token.SignedString([]byte(ConstJwtSecretKey))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				//fmt.Fprintln(w, "Error while signing the token")
				response = JwtToken{false,tokenString,"Error while signing the token"}
			}else{
				response = JwtToken{true,tokenString,"success."}
			}
		}
		ReturnJsonResponse(response, w)
	}
}

func (own *LoginController)ValidateLoginToken(w http.ResponseWriter, r *http.Request)(*jwt.Token,bool){
	var b bool
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {return []byte(ConstJwtSecretKey), nil})
	if err == nil {
		if token.Valid {
			//fmt.Fprint(w, "Token is true.")
			b = true
		} else {
			//w.WriteHeader(http.StatusUnauthorized)
			//fmt.Fprint(w, "Token is not valid")
			b =false
		}
	} else {
		//w.WriteHeader(http.StatusUnauthorized)
		//fmt.Fprint(w, "Unauthorized access to this resource")
		b =false
	}
	return token,b
}

func (own *LoginController)Main(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("ok")
	fmt.Fprintf(w,"hello")
}


func (own *LoginController)Out(w http.ResponseWriter, r *http.Request){
	//验证页面 防止缓存
	w.Header().Set("Cache-Control","no-cache")
	w.Header().Set("Pragma","no-cache")
	w.Header().Set("Expires", "0")
	session, _ := Session.Get(r,ConstSessionAdminLoginFlag)
	session.Values[ConstSessionAdminLoginFlagValues]=nil
	session.Save(r,w)
	http.Redirect(w,r,"login",301)

}


