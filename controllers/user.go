package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	//c.Ctx.WriteString("Get User")
	c.Layout = "layouts/admin-lte.tpl"
	
	c.Data["Website"] = "beego.user.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "user/index.tpl"
}
