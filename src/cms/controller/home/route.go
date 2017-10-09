package admin

var Route  map[string]interface{} = map[string]interface{}{
		"/admin/login":&Login{},
		"/admin/test":&AdminController{},
}
