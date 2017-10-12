package initialize

import (
	_ "lib/xorm/drivers/mysql"
	"lib/xorm/xorm"
	"log"
	"time"
	"lib/xorm/core"
	"fmt"
)

var Orm  *xorm.Engine

func init()  {
	username := "root"
	password := "123456"
	server :="127.0.0.1"
	port := "3306"
	db_name := "minggo"
	charset := "utf8"
	Orm = NewOrm(server , username , password , db_name , charset , port )
}

func NewOrm(server string, username string, password string, db_name string, charset string, port string) *xorm.Engine {
	log.Printf("db initializing...")
	var err error
	orm, err := xorm.NewEngine("mysql", username+":"+password+"@tcp("+server+":"+port+")/"+db_name+"?charset="+charset+"&parseTime=true")
	//Orm, err = xorm.NewEngine("mysql", "root:@/chat?charset=utf8&parseTime=true")
	//fmt.Print(username+":"+password+"@tcp("+server+")/"+db_name+"?charset="+charset+"&parseTime=true")
	if err != nil {
		fmt.Println("mysql:error:",err)
	}
	orm.TZLocation = time.Local
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "m_")
	orm.SetTableMapper(tbMapper)
	return orm
}