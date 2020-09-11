package controllers

import (
	"github.com/astaxie/beego"
)

//FirstController struct
type FirstController struct {
	beego.Controller
}

//Employee struct
type Employee struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

//Employees type is declared
type Employees []Employee

var employees []Employee

func init() {

	employees = Employees{

		Employee{ID: 1, FirstName: "kishore", LastName: "yerubandi"},
		Employee{ID: 2, FirstName: "Rishika", LastName: "yerubandi"},
	}
}

//GetEmployees func
func (f *FirstController) GetEmployees() {

	f.Ctx.ResponseWriter.WriteHeader(200)
	f.Data["json"] = employees
	f.ServeJSON()
}

//Dashboard func
func (f *FirstController) Dashboard() {

	f.Data["employees"] = employees
	f.TplName = "dashboard.tpl"
}

//GetEmployee gghh
func (f *FirstController) GetEmployee() {
	var id int
	f.Ctx.Input.Bind(&id, "id")
	var isEmployeeExist bool
	var emps []Employee
	for _, employee := range employees {
		if employee.ID == id {
			emps = append(emps, Employee{ID: employee.ID,
				FirstName: employee.FirstName, LastName: employee.LastName})
			isEmployeeExist = true
			break
		}
	}
	if !isEmployeeExist {
		f.Abort("Generic")
	} else {
		f.Data["employees"] = emps
		f.TplName = "dashboard.tpl"
	}
}
