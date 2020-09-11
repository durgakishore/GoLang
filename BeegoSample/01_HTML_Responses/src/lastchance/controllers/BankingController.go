package controllers

import (
	"BeegoSample/01_HTML_Responses/src/lastchance/models"

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
	c.Data["accounts"] = []models.Account{
		models.Account{
			ID:     1,
			Name:   "Checking",
			Number: "8888",
			Amount: 642.27,
		},
		models.Account{
			ID:     2,
			Name:   "Savings",
			Number: "3344",
			Amount: 1000,
		},
	}

	c.TplName = "banking.tpl"
}
