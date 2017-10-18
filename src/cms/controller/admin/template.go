package admin

import (
	"lib/template"
	//"lib/o-concurrent-map"
	//"fmt"
)

var Template = template.NewTemplate()

func init()  {

	ds := template.SYSDATA{
		ConstTemplateSysdataStatic,
		ConstTemplateSysdataPowerBy,
		ConstTemplateSysdataPostUrl,
	}

	//Map_DefalutStatic := cmap.New()
	//Map_DefalutStatic.Set("SYS",ds)
	//// Retrieve item from map.
	//DefalutStatic, ok := Map_DefalutStatic.Get("SYS")
	//if !ok{
	//	fmt.Println("Map_Create_Error.")
	//}

	Template.Options = &template.Options{
		Directory:ConstTemplateOptionsDirectory,
		Charset: ConstTemplateOptionsCharset, // Appends the given charset to the Content-Type header. Default is UTF-8
		// Allows changing of output to XHTML instead of HTML. Default is "text/html"
		HTMLContentType: ConstTemplateOptionsHTMLContentType,
		Extension: ConstTemplateOptionsExtension,
	}
	Template.TemplateData = &template.TemplateData{ds,nil}

}