package controllers

import (
	"encoding/json"
	"fmt"
	"lastchance/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
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

	valid := validation.Validation{}
	isValid, _ := valid.Valid(&transfer)
	fmt.Println(transfer)
	fmt.Println(valid.ErrorMap())

	var message string
	if isValid {
		message = "success"
	} else {
		message = "failure"
	}
	c.Ctx.WriteString(message)
}
