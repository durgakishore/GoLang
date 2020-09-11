package main

import (
	models "BeegoORM/models"
	_ "BeegoORM/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// This is a dummy change to test Hound
	//var Go_Ing_Go int
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "rapiscan:Anubis2830@tcp(127.0.0.1:1433)/SysDb")
	orm.RegisterModel(new(models.Article))
}

func main() {
	beego.Run()
}
