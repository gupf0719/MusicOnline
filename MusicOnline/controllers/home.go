package controllers

import (
	"MusicOnline/MusicOnline/models"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["IsHome"] = true
	this.TplName = "home.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	news, err := models.GetAllNews("", true)
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["News"] = news
	}

	stage, err := models.GetAllStage(true)
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Stage"] = stage
	}
}
