package controllers

import (
	"MusicOnline/MusicOnline/models"

	"github.com/astaxie/beego"
)

type AdminmessageController struct {
	beego.Controller
}

func (this *AdminmessageController) Get() {
	this.Data["IsMessage"] = true
	this.TplName = "adminmessage.html"

	messages, err := models.GetAllMusicMessages()
	if err != nil {
		beego.Error(err)
		return
	}

	this.Data["Messages"] = messages
	this.Data["IsLogin"] = checkAccount(this.Ctx)
}

func (this *AdminmessageController) Delete() {
	err := models.DelMusicMessage(this.Input().Get("mid"))
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/adminmessage", 302)
}
