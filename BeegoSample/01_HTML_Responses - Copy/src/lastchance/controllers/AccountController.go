package controllers

import (
	"BeegoSample/01_HTML_Responses/src/lastchance/models"
	"fmt"

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
