package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//ManageController strct
type ManageController struct {
	beego.Controller
}

//Details strt
type Details struct {
}

//Select func
func (manage *ManageController) Select() {

	flash := beego.ReadFromRequest(&manage.Controller)

	if ok := flash.Data["error"]; ok != "" {
		// Display error messages
		manage.Data["errors"] = ok
		manage.Ctx.WriteString("errors")
	}

	if ok := flash.Data["notice"]; ok != "" {
		// Display error messages
		manage.Data["notices"] = ok
		manage.Ctx.WriteString("notices")
	}

	o := orm.NewOrm()
	o.Using("company")

	var details []Details
	num, err := o.QueryTable("T_Bag_Info").All(&details)

	if err != orm.ErrNoRows && num > 0 {
		manage.Data["records"] = details
	}
}
