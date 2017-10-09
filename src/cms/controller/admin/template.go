package admin

import (
	"lib/template"
)

var Template = template.NewTemplate()
var DefalutStatic map[string]interface{} =  make(map[string]interface{})
func init()  {
	DefalutStatic["__STATIC__"] = "http://192.168.14.7:8001/"
	Template.Options = &template.Options{
		Directory:"./src/cms/views/admin/",
		Charset: "UTF-8", // Appends the given charset to the Content-Type header. Default is UTF-8
		// Allows changing of output to XHTML instead of HTML. Default is "text/html"
		HTMLContentType: "text/html",
		Extension: ".html",
		TemplateData: DefalutStatic,
	}

}