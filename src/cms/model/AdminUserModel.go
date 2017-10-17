package model

import (
	"cms/initialize"
)

type AdminUser struct {
	Id 			int			`xorm:"int(11) pk autoincr"`
	Username 	string		`xorm:"char(30) notnull"`
	Password 	string		`xorm:"varchar(40) notnull"`
	Roleid   	int			`xorm:"int(10) notnull"`
	Regtime     int			`xorm:"int(10) notnull created"`
	Email    string 		`xorm:"varchar(60) notnull"`
	Status   int    		`xorm:"bool notnull"`
}
type Group struct {
	Id    int    `xorm:"int(11) pk autoincr"`
	Name  string `xorm:"varchar(25) unique  notnull"`
	Num   int    `xorm:"int(11) notnull"`
	Ctime int    `xorm:"created notnull"`
}


func (self *AdminUser)GetByUsernameAndPassword(username string,password string) (bool,*AdminUser,error) {
	u := &AdminUser{Username:username,Password:password}
	//log.Println(u)
	r,err := initialize.Orm.Get(u)
	u.Password=""
	return r,u,err
}

func (self *AdminUser)IsExist()  {

}
