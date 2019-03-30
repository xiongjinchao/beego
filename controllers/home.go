package controllers

type HomeController struct {
	BaseController
}

func (c *HomeController) Dashboard() {
	c.TplName = "home/dashboard.tpl"
}