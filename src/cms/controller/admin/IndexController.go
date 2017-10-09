package admin

import (
	"net/http"
	"fmt"
	//"lib/mustache"
)
var TemplateData  = make(map[string]interface{})
//登录
type Login struct {
	username string
	passwrod string
	is_login bool
}



func  (own *Login)GET(w http.ResponseWriter, r *http.Request)  {
	session, _ := Session.Get(r,"test")
	// Set some session values.
	session.Values["foo"] = "bar"
	session.Values[42] = 43
	// Save it before we write to the response/return from the handler.
	session.Save(r, w)

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
	TemplateData["title"] ="OK"
	Template.Html(w,r,"login",TemplateData)
	//s := mustache.RenderFileInLayout("./src/cms/views/admin/login.html", "./src/cms/views/admin/layout.html.mustache", nil)
	//fmt.Fprintf(w,s)
}

func  (own *Login)POST(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "admin_post_login")

}

