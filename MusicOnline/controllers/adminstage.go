package controllers

import (
	"github.com/astaxie/beego"
)

type AdminstageController struct {
	beego.Controller
}

func (this *AdminstageController) Get(){
	this.Data["IsStage"] = true
	this.TplName="adminstage.html"
}
