package main

import (
	_ "hello/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/astaxie/beego/session/redis"
)

func init(){
    dbhost := beego.AppConfig.String("dbhost")
    dbport := beego.AppConfig.String("dbport")
    dbuser := beego.AppConfig.String("dbuser")
    dbpassword := beego.AppConfig.String("dbpassword")
    db := beego.AppConfig.String("db")
 
    orm.RegisterDriver("mysql", orm.DRMySQL)
    conn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + db + "?charset=utf8"
    orm.RegisterDataBase("default", "mysql", conn)
}

func main() {
	beego.Run()
}