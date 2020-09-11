package controllers

import (
	"encoding/json"
	"fmt"
	"lastchance/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
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

// @router /api/transfer [post]
func (c *BankingController) Transfer() {

	var transfer models.Transfer

	json.Unmarshal(c.Ctx.Input.RequestBody, &transfer)

	valid := validation.Validation{}
	isValid, _ := valid.Valid(&transfer)
	fmt.Println(transfer)
	fmt.Println(valid.ErrorMap())

	response := models.TransferResponse{
		Transfer: transfer,
	}

	if isValid {
		response.Status = "success"
	} else {
		response.Status = "failure"
	}
	c.Data["json"] = response
	c.ServeJSON()
}

// @router /lifecycle [get]
func (c *BankingController) ShowLifecycle() {
	fmt.Println("Action Execution")
}

func (c *BankingController) Init(ctx *context.Context, controllerName,
	actionName string, app interface{}) {
	fmt.Printf("Initialization: %s.%s\n", controllerName, actionName)
	c.Controller.Init(ctx, controllerName, actionName, app)
}

func (c *BankingController) Prepare() {
	fmt.Println("Prepare controller")
	c.Controller.Prepare()
}

func (c *BankingController) Render() error {
	fmt.Println("Render result")
	// c.Ctx.WriteString("result")

	return nil
}

func (c *BankingController) Finish() {
	fmt.Println("Finish controller")
	c.Controller.Finish()
}
