package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["lastchance/controllers:BankingController"] = append(beego.GlobalControllerRouter["lastchance/controllers:BankingController"],
		beego.ControllerComments{
			"ShowAccounts",
			`/banking`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["lastchance/controllers:BankingController"] = append(beego.GlobalControllerRouter["lastchance/controllers:BankingController"],
		beego.ControllerComments{
			"Transfer",
			`/api/transfer`,
			[]string{"post"},
			nil})

}
