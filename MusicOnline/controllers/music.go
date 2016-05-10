package controllers

import (
	"MusicOnline/MusicOnline/models"

	"github.com/astaxie/beego"
)

type MusicController struct {
	beego.Controller
}

func (this *MusicController) Get() {

	this.TplName = "music.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

}

func (this *MusicController) Post() {
	this.TplName = "music.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	name := this.Input().Get("musicname")

	music, err := models.SearchMusic(name)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Music"] = music
}
