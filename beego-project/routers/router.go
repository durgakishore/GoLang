package routers

import (
	"beego-project/controllers"
	"beego-project/filters"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/employees", &controllers.FirstController{}, "get:GetEmployees")
	beego.Router("/dashboard", &controllers.FirstController{}, "get:Dashboard")
	beego.InsertFilter("/*", beego.BeforeRouter,filters.LogManager)
	beego.Router("/home", &controllers.SessionController{}, "get:Home")
	beego.Router("/login", &controllers.SessionController{}, "get:Login")
	beego.Router("/logout", &controllers.SessionController{}, "get:Logout")
}
