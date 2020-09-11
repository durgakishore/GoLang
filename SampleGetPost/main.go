package main

import (
	"github.com/astaxie/beego"
	//"samplegetpost/routers"
)

func main() {

	//	beego.SetStaticPath("/public", "static")
	//beego.DelStaticPath("/static")
	beego.BConfig.Listen.HTTPPort = 8081
	beego.Run()
}
