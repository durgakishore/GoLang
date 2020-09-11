package main

import (
	_ "lastchance/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/public", "static")
	beego.DelStaticPath("/static")
	beego.Run()
	f := map[string]string{"foo": "bar"}
	fmt.Println(f)
}
