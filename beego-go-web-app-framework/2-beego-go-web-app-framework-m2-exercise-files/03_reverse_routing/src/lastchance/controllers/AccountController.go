package controllers

import (
	"github.com/astaxie/beego"
)

type AccountController struct {
	beego.Controller
}

func (c *AccountController) Login() {
	c.TplName = "login.tpl"
}

func (c *AccountController) Create() {
	c.TplName = "createAccount.tpl"
}
