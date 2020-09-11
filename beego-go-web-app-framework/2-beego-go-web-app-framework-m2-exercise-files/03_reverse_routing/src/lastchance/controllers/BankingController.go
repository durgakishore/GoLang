package controllers

import (
	"github.com/astaxie/beego"
)

type BankingController struct {
	beego.Controller
}

func (c *BankingController) URLMapping() {
	c.Mapping("ShowAccounts", c.ShowAccounts)
}

// @router /banking [get]
func (c *BankingController) ShowAccounts() {
	c.TplName = "banking.tpl"
}
