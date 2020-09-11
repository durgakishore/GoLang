package controllers

import (
	"github.com/astaxie/beego"
)

// MainController struct 
type MainController struct {
	beego.Controller
}

//Get  func
func (main *MainController) Get() {
	//main.Data["Website"] = "beego.me"
	main.Data["Email"] = "astaxie@gmail.com"
	main.TplName = "index.tpl"
}

//HelloSitepoint  func
func (main *MainController) HelloSitepoint() {
	//main.Data["Website"] = "My Website"
	main.Data["Email"] = "your.email.address@example.com"
	main.Data["EmailName"] = "Your Name"
	main.TplName = "default/hello-sitepoint.tpl"
}
