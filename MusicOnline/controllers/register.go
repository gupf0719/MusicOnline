package controllers

import (
	"MusicOnline/MusicOnline/models"

	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.TplName = "register.html"

}

func (this *RegisterController) Post() {

//	this.TplName = "register.html"

	name := this.Input().Get("uname")
	password := this.Input().Get("pwd")
	nickname := this.Input().Get("nname")

	user := models.User{Name: name, Password: password, NickName: nickname}

	models.AddUser(&user)

	this.Redirect("/login",301)
}
