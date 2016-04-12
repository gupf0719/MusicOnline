package controllers

import (
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.TplName = "register.html"
}

/*
func (this *RegisterController) Post() {
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")

	err := models.AddUser(uname, pwd)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/", 301)
	return
}
*/
