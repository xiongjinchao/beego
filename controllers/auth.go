package controllers

import (
	//"fmt"
	//"errors"
	models "hello/models"
	"encoding/json"
	//"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) GetLogin() {
	c.Layout = ""
	c.Data["xsrf_token"] = c.XSRFToken()
	c.TplName = "auth/login.tpl"
}

func (c *AuthController) PostLogin() {
	data := map[string]string{"id":"1","name":"Teddy"}
	user, err := json.Marshal(data)
    if err != nil {
    	//flash data
        c.Ctx.Redirect(302, "/login")
    }
	c.SetSession("user",string(user))
	c.Ctx.Redirect(302, "/")
}

func (c *AuthController) GetRegister() {
	c.Layout = ""
	c.Data["xsrf_token"] = c.XSRFToken()
	flash := beego.ReadFromRequest(&c.Controller)
    if n, ok := flash.Data["error"]; ok{
		errors := make(map[string]interface{})
	    err := json.Unmarshal([]byte(n), &errors)
		if err == nil {
	        c.Data["errors"] = errors
	    }
    }

    c.TplName = "auth/register.tpl"
}

func (c *AuthController) PostRegister() {
	flash:=beego.NewFlash()

	name := c.Input().Get("name")
	email := c.Input().Get("email")
	//password := c.Input().Get("password")
	confirmPassword := c.Input().Get("confirm_password")

	orm.Debug = true;
	o := orm.NewOrm()
	user := models.User{}
	user.Name = name
	err := o.Read(&user, "Name")

	if err != orm.ErrNoRows {
		flash.Error(`{"name":"user name exist!"}`)
        flash.Store(&c.Controller)
        c.Ctx.Redirect(302, "/register")
		return
	}

	o = orm.NewOrm()
	user = models.User{Email: email}
	err = o.Read(&user, "Email")

	if err != orm.ErrNoRows {
		flash.Error(`{"email":"user email exist!"}`)
        flash.Store(&c.Controller)
        c.Ctx.Redirect(302, "/register")
		return
	}

	valid := validation.Validation{}
	user = models.User{}
	if err = c.ParseForm(&user); err == nil {
		valid.Required(user.Name, "name")
		valid.MinSize(user.Name, 6, "name")
	    valid.MaxSize(user.Name, 20, "name")

	    valid.Required(user.Email, "email")
	    valid.Email(user.Email, "email")

	    valid.Required(user.Password, "password")
	    valid.MinSize(user.Password, 6, "password")
	    valid.MaxSize(user.Password, 20, "password")

	    valid.Required(confirmPassword, "confirm_password")
	    valid.MaxSize(confirmPassword, 20, "confirm_password")
	    valid.MinSize(confirmPassword, 6, "confirm_password")

	    if valid.HasErrors() {
	        for _, err := range valid.Errors {
	            //log.Println(err.Key, err.Message)
	            flash.Error(`{"`+err.Key+`":"`+err.Message+`"}`)
		        flash.Store(&c.Controller)
		        c.Ctx.Redirect(302, "/register")
		        return
	        }
	    }

	    if(user.Password != confirmPassword){
    		flash.Error(`{"confirm_password":"password not eq confirm_password!"}`)
	        flash.Store(&c.Controller)
	        c.Ctx.Redirect(302, "/register")
	        return
	    }
	}


	/*
	data := map[string]string{"id":"1","name":"Teddy"}
	user, err := json.Marshal(data)
    if err != nil {
    	flash.Error("Settings invalid!")
        flash.Store(&c.Controller)
        c.Redirect("/setting",302)
        return
        //c.Ctx.Redirect(302, "/login")
    }

	c.SetSession("user",string(user))
	c.Ctx.Redirect(302, "/")
	*/
}

func (c *AuthController) Logout() {
	c.DelSession("user")
	c.DestroySession()
	c.Ctx.Redirect(302, "/login")
}