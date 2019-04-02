package routers

import (
	"github.com/astaxie/beego"
	"hello/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{},"get:Dashboard")

	beego.Router("/login", &controllers.AuthController{},"get:GetLogin")
	beego.Router("/login", &controllers.AuthController{},"post:PostLogin")
	beego.Router("/register", &controllers.AuthController{},"get:GetRegister")
	beego.Router("/register", &controllers.AuthController{},"post:PostRegister")
	beego.Router("/logout", &controllers.AuthController{},"*:Logout")

	beego.Router("/user", &controllers.UserController{})
	//beego.Router("/article", &controllers.ArticleController{})
	beego.Include(&controllers.ArticleController{})
}
