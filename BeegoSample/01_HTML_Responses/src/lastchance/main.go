package main

import (
	_ "BeegoSample/01_HTML_Responses/src/lastchance/routers"

	"fmt"

	"github.com/astaxie/beego"
)

func main() {
	//beego.SetStaticPath("/public", "static")
	//beego.DelStaticPath("/static")
	beego.BConfig.Listen.HTTPPort = 8082
	beego.Run()
	f := map[string]string{"foo": "bar"}
	fmt.Println(f)
}
