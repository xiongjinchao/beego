package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	models "hello/models"
)

type UserController struct {
	BaseController
}

func (c *UserController) Index() {
	users := []*models.User{}
	o := orm.NewOrm()
	_, err := o.QueryTable("user").All(&users)
	if err != nil {
		fmt.Println("查询出错了")
	}
	c.Data["users"] = &users
	c.TplName = "user/index.tpl"
}

func (c *UserController) Create() {
	c.TplName = "user/create.tpl"
}

func (c *UserController) Store() {
	o, valid, flash := orm.NewOrm(), validation.Validation{}, beego.NewFlash()
	user := models.User{}
	if err := c.ParseForm(&user); err != nil {
		fmt.Println("表单赋值出错了")
		return
	}
	if _, err := valid.Valid(user); err != nil {
		for _, err := range valid.Errors {
			flash.Error(`{"` + err.Key + `":"` + err.Message + `"}`)
			flash.Store(&c.Controller)
			c.Ctx.Redirect(302, "/user/create")
			return
		}
	}
	o.Insert(user)
	c.Ctx.Redirect(302, "/user")
}

func (c *UserController) Edit() {
	o := orm.NewOrm()
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	user := models.User{Id: id}
	o.Read(&user, "Id")
	c.Data["user"] = &user
	c.TplName = "user/edit.tpl"
}

func (c *UserController) Update() {
	o := orm.NewOrm()
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	user := models.User{Id: id}
	if err := o.Read(&user, "Id"); err != nil {
		fmt.Println("查询出错了")
		return
	}
	if err := c.ParseForm(&user); err != nil {
		fmt.Println("表单赋值出错了")
		return
	}
	if _, err := o.Update(&user); err != nil {
		fmt.Println("更新出错了")
	}
	c.Ctx.Redirect(302, "/user")
}

func (c *UserController) Delete() {
	c.Ctx.WriteString("Delete")
}
