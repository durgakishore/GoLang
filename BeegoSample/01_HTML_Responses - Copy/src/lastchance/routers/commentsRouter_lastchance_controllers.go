package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["lastchance/controllers:BankingController"] = append(beego.GlobalControllerRouter["lastchance/controllers:BankingController"],
		beego.ControllerComments{
			Method:           "ShowAccounts",
			Router:           `/banking`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["lastchance/controllers:BankingController"] = append(beego.GlobalControllerRouter["lastchance/controllers:BankingController"],
		beego.ControllerComments{
			Method:           "Transfer",
			Router:           `/api/transfer`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

}
