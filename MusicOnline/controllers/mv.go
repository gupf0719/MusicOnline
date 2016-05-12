package controllers

import (
	"MusicOnline/MusicOnline/models"

	"github.com/astaxie/beego"
)

type MvController struct {
	beego.Controller
}

func (this *MvController) Get() {
	this.TplName = "mv.html"

	mv, err := models.GetMv(this.Input().Get("name"))

	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Mv"] = mv
	}
}
