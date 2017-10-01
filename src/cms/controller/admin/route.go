package admin

var Route  map[string]interface{} = map[string]interface{}{
		"/admin/login":"&admin.AdminController{}",
		"/home/index":"&home.HomeController{}",
}
