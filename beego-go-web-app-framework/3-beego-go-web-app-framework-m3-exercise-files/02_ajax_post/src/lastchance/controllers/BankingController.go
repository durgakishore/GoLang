package controllers

import (
	"encoding/json"
	"fmt"
	"lastchance/models"

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

// @router /api/transfer [post]
func (c *BankingController) Transfer() {
	var transfer models.Transfer
	json.Unmarshal(c.Ctx.Input.RequestBody, &transfer)
	fmt.Println(transfer)

	c.Ctx.WriteString("success")
}
