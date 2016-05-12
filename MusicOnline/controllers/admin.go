package controllers

import (
	"MusicOnline/MusicOnline/models"

	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get() {
	this.Data["IsHome"] = true
	this.TplName = "admin.html"

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
