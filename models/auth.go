package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Auth struct {
	Id int64 `orm:"auto"`
	Name string `orm:"size(128)" form:"name" valid:"Required;MaxSize(20);MinSize(6)"`
	Email string `orm:"unique;size(128)" form:"email" valid:"Email"`
	Password string `orm:"size(128)" form:"password" valid:"Required;MaxSize(20);MinSize(6)"`
	RememberToken string `orm:"size(128)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	tbprefix := beego.AppConfig.String("tbprefix")
	orm.RegisterModelWithPrefix(tbprefix, new(Auth))
	//orm.RegisterModel(new(Auth))
}

// func (m *Auth) TableName() string {
// 	tbprefix := beego.AppConfig.String("tbprefix")
//     return tbprefix + "user"
// }
