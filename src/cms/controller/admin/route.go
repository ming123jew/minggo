package admin

var Route  map[string] map[string]interface{} = map[string] map[string]interface{}{
		"/admin/login": {"struct":&Login{},"rbac":&Rbac{}} ,
		"/admin/index":{"struct":&Index{},"rbac":&Rbac{}} ,
		"/admin/test": {"struct":&AdminController{}} ,
}
