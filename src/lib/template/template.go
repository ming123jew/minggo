package template

import (
	"fmt"
	"net/http"
	"html/template"
	"sync"
)

type MingGoTemplate struct {
	mutex  sync.RWMutex
	Options *Options
	TemplateData *TemplateData
}
type Options struct {
	Directory string
	Charset string
	HTMLContentType string
	Extension string
}

type TemplateData struct {
	SYSDATA		interface{}
	PAGEDATA	interface{}
}

type SYSDATA struct {
	STATIC string
	POWER_BY string
	POSTURL string
}

func NewTemplate(arg ...string) *MingGoTemplate  {
	return &MingGoTemplate{
		Options:&Options{
			Directory: "./src/cms/views/",
			Charset: "UTF-8", // Appends the given charset to the Content-Type header. Default is UTF-8
			// Allows changing of output to XHTML instead of HTML. Default is "text/html"
			HTMLContentType: "text/html",
			Extension: ".html",
		},
	}
}

func (self *MingGoTemplate)Html(w http.ResponseWriter,r *http.Request,name string,templateData interface{})  {

	t,error:=template.ParseFiles(self.Options.Directory+name+self.Options.Extension)//New("login.html")
	if error!=nil{
		fmt.Fprintf(w,error.Error())
		return
	}
	//error = t.Execute(w, TemplateData)
	//fmt.Println(self.Options.TemplateData)
	//for k ,v := range self.Options.TemplateData { templateData[k]=v }
	//for k ,v := range self.Options.TemplateData { fmt.Println(k,v)}
	//fmt.Println(templateData)
	w.Header().Set("Content-Type", self.Options.HTMLContentType+"; "+self.Options.Charset)
	error = t.Execute(w, templateData)
	if error != nil{
		fmt.Fprintf(w,error.Error())
		return
	}
}

func (self *MingGoTemplate)SetTemplateData(i interface{})  {
	self.TemplateData.PAGEDATA = i
}