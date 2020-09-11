package routers

import (
	"BeegoSample/01_HTML_Responses/src/lastchance/controllers"

	"github.com/astaxie/beego"
)

func init() {
	/*beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.AccountController{})
	beego.Include(&controllers.BankingController{})*/

	beego.Router("/banking", &controllers.BankingController{}, "get:ShowAccounts")
}
