package initialize

import (
	_ "lib/o-xorm/drivers/mysql"
	"lib/o-xorm/xorm"
	"time"
	"lib/o-xorm/core"
	"fmt"
	"cms/config"
)

var Orm  *xorm.Engine

func init()  {
	//username := config.DB_USER
	//password := config.DB_PASS
	//server := config.DB_HOST
	//port := config.DB_PORT
	//db_name := config.DB_NAME
	//charset := config.DB_CHARSET
	Orm = NewOrm(config.DB_HOST , config.DB_USER , config.DB_PASS , config.DB_NAME , config.DB_CHARSET , config.DB_PORT )
}

func NewOrm(server string, username string, password string, db_name string, charset string, port string) *xorm.Engine {
	fmt.Println("db initializing...")
	var err error
	orm, err := xorm.NewEngine(config.DB_TYPE, username+":"+password+"@tcp("+server+":"+port+")/"+db_name+"?charset="+charset+"&parseTime=true")
	//Orm, err = xorm.NewEngine("mysql", "root:@/chat?charset=utf8&parseTime=true")
	//fmt.Print(username+":"+password+"@tcp("+server+")/"+db_name+"?charset="+charset+"&parseTime=true")
	if err != nil {
		fmt.Println("mysql:error:",err)
	}
	orm.TZLocation = time.Local
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, config.DB_PREFIX)
	orm.SetTableMapper(tbMapper)
	return orm
}