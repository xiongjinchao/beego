package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
	_ "hello/routers"
)

func init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	db := beego.AppConfig.String("db")

	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	conn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + db + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", conn)
}

func main() {
	beego.Run()
}
