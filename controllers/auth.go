package controllers

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	models "hello/models"
	"strconv"
	"time"

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
	flash := beego.ReadFromRequest(&c.Controller)
	if n, ok := flash.Data["error"]; ok {
		errors := make(map[string]interface{})
		err := json.Unmarshal([]byte(n), &errors)
		if err == nil {
			c.Data["errors"] = errors
		}
	}
	c.TplName = "auth/login.tpl"
}

func (c *AuthController) PostLogin() {
	flash := beego.NewFlash()

	name := c.Input().Get("name")
	password := c.Input().Get("password")

	sha1 := sha1.New()
	sha1.Write([]byte(password))
	password = hex.EncodeToString(sha1.Sum([]byte("")))

	o := orm.NewOrm()
	user := models.User{Name: name, Password: password}

	if err := o.Read(&user, "Name", "Password"); err == orm.ErrNoRows {
		flash.Error(`{"name":"name and password not match!"}`)
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/login")
		return
	}

	id := strconv.FormatInt(user.Id, 10)
	data := map[string]string{"id": id, "name": user.Name}
	if auth, err := json.Marshal(data); err != nil {
		c.Ctx.Redirect(302, "/login")
	} else {
		c.SetSession("auth", string(auth))
		c.Ctx.Redirect(302, "/")
	}
}

func (c *AuthController) GetRegister() {
	c.Layout = ""
	c.Data["xsrf_token"] = c.XSRFToken()
	flash := beego.ReadFromRequest(&c.Controller)
	if n, ok := flash.Data["error"]; ok {
		errors := make(map[string]interface{})
		err := json.Unmarshal([]byte(n), &errors)
		if err == nil {
			c.Data["errors"] = errors
		}
	}

	c.TplName = "auth/register.tpl"
}

func (c *AuthController) PostRegister() {
	flash := beego.NewFlash()

	name := c.Input().Get("name")
	email := c.Input().Get("email")
	confirmPassword := c.Input().Get("confirm_password")

	o := orm.NewOrm()
	user := models.User{Name: name}

	if err := o.Read(&user, "Name"); err != orm.ErrNoRows {
		flash.Error(`{"name":"user name exist!"}`)
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/register")
		return
	}

	user = models.User{Email: email}
	if err := o.Read(&user, "Email"); err != orm.ErrNoRows {
		flash.Error(`{"email":"user email exist!"}`)
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/register")
		return
	}

	valid := validation.Validation{}
	user = models.User{}
	if err := c.ParseForm(&user); err == nil {
		valid.Required(user.Name, "name")
		valid.MinSize(user.Name, 5, "name")
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
				flash.Error(`{"` + err.Key + `":"` + err.Message + `"}`)
				flash.Store(&c.Controller)
				c.Ctx.Redirect(302, "/register")
				return
			}
		}

		if user.Password != confirmPassword {
			flash.Error(`{"confirm_password":"password not eq confirm_password!"}`)
			flash.Store(&c.Controller)
			c.Ctx.Redirect(302, "/register")
			return
		}
	}

	sha1 := sha1.New()
	sha1.Write([]byte(user.Password))
	password := hex.EncodeToString(sha1.Sum([]byte("")))

	user.Password = password
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	if _, err := o.Insert(&user); err != nil {
		flash.Error(`{"name":"register fail!"}`)
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/register")
		return
	}

	id := strconv.FormatInt(user.Id, 10)
	data := map[string]string{"id": id, "name": user.Name}
	if auth, err := json.Marshal(data); err != nil {
		flash.Error("Settings invalid!")
		flash.Store(&c.Controller)
		c.Redirect("/setting", 302)
		return
	} else {
		c.SetSession("auth", string(auth))
		c.Ctx.Redirect(302, "/")
	}

}

func (c *AuthController) Logout() {
	c.DelSession("auth")
	c.DestroySession()
	c.Ctx.Redirect(302, "/login")
}
