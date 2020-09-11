package routers

import (
	"ORM/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/select", &controllers.ManageController{}, "get:Select")
}
