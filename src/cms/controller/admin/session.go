package admin

import "lib/sesson"

var Session = sessions.NewCookieStore([]byte("minggo-admin"))
func init()  {
	//设置session默认参数
	Session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 1,//1天
		HttpOnly: true,
	}
}
