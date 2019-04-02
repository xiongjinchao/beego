package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	c.Layout = "layouts/admin-lte.tpl"

	v := c.GetSession("user")
	if v == nil{
		c.Ctx.Redirect(302, "/login")
	}
	user := make(map[string]interface{})
    err := json.Unmarshal([]byte(v.(string)), &user)
	if err != nil {
        c.Ctx.Redirect(302, "/login")
    }

	c.Data["user"] = user
	c.Data["xsrf_token"] = c.XSRFToken()

	c.Data["Website"] = "beego.user.me"
	c.Data["Email"] = "astaxie@gmail.com"
}
