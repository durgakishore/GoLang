package controllers

import "github.com/astaxie/beego"

//ErrorController rtr
type ErrorController struct {
	beego.Controller
}

//Error404 rge
func (c *ErrorController) Error404() {
	c.Data["content"] = "Page Not Found"
	c.TplName = "404.tpl"
}

//Error500 er
func (c *ErrorController) Error500() {
	c.Data["content"] = "Internal Server Error"
	c.TplName = "500.tpl"
}

//ErrorGeneric we
func (c *ErrorController) ErrorGeneric() {
	c.Data["content"] = "Some Error Occurred"
	c.TplName = "genericerror.tpl"
}
