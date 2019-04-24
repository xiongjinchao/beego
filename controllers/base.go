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

	v := c.GetSession("auth")
	if v == nil {
		c.Ctx.Redirect(302, "/login")
	}
	auth := make(map[string]interface{})
	err := json.Unmarshal([]byte(v.(string)), &auth)
	if err != nil {
		c.Ctx.Redirect(302, "/login")
	}

	c.Data["auth"] = auth
	c.Data["xsrf_token"] = c.XSRFToken()

}
