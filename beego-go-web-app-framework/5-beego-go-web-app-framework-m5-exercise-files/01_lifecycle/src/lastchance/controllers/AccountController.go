package controllers

import (
	"fmt"
	"lastchance/models"

	"github.com/astaxie/beego"
)

type AccountController struct {
	beego.Controller
}

func (c *AccountController) Login() {
	if c.Ctx.Input.IsPost() {
		var loginForm models.LoginForm
		c.ParseForm(&loginForm)
		fmt.Println(loginForm)
	}
	c.TplName = "login.tpl"
}

func (c *AccountController) Create() {
	c.TplName = "createAccount.tpl"
}
