package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["BeegoSample/01_HTML_Responses/src/lastchance/controllers:BankingController"] = append(beego.GlobalControllerRouter["BeegoSample/01_HTML_Responses/src/lastchance/controllers:BankingController"],
        beego.ControllerComments{
            Method: "Transfer",
            Router: `/api/transfer`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BeegoSample/01_HTML_Responses/src/lastchance/controllers:BankingController"] = append(beego.GlobalControllerRouter["BeegoSample/01_HTML_Responses/src/lastchance/controllers:BankingController"],
        beego.ControllerComments{
            Method: "ShowAccounts",
            Router: `/banking`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
