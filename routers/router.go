package routers

import (
	"beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "get:Dashboard")

	beego.Router("/login", &controllers.AuthController{}, "get:GetLogin;post:PostLogin")
	beego.Router("/register", &controllers.AuthController{}, "get:GetRegister;post:PostRegister")
	beego.Router("/logout", &controllers.AuthController{}, "*:Logout")

	beego.Router("/user", &controllers.UserController{}, "get:Index")
	beego.Router("/user/create", &controllers.UserController{}, "get:Create;post:Store")
	beego.Router("/user/:id", &controllers.UserController{}, "get:Edit;post:Update;delete:Delete")

	beego.Router("/article", &controllers.ArticleController{})
}
