package routers

import (
	"github.com/astaxie/beego"
	"samplegetpost/controllers"
)

func init() {

	beego.Router("/ShowUsers", &controller.BankingController{})
}
