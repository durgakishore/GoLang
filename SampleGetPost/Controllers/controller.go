package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

//MainController of the controller package
type BankingController struct {
	beego.controller
}

func (c *BankingController) URLMapping() {
	c.Mapping("ShowUsers", c.ShowUsers)
}

func (c *BankingController) ShowUsers() {
	//if c.Ctx.Input.IsGet() {
	fmt.Println("Kishore")
	//}
}
