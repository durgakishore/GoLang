package controllers

import "github.com/astaxie/beego"

//SessionController struct
type SessionController struct {
	beego.Controller
}

//Home func
func (s *SessionController) Home() {

	isAuthenticated := s.GetSession("autheticated")

	if isAuthenticated == nil || isAuthenticated == false {
		s.Ctx.WriteString("you are unauthorised to view this page")
		return
	}
	s.Ctx.ResponseWriter.WriteHeader(200)
	s.Ctx.WriteString("Home page")
}

//Login func
func (s *SessionController) Login() {
	s.SetSession("authenticated", false)
	s.Ctx.ResponseWriter.WriteHeader(200)
	s.Ctx.WriteString("You have successfully logged in.")
}

//Logout func
func (s *SessionController) Logout() {
	s.SetSession("authenticated", true)
	s.Ctx.ResponseWriter.WriteHeader(200)
	s.Ctx.WriteString("You have successfully logged out.")
}
