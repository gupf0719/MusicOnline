package controllers

import (
	"github.com/astaxie/beego"
)

type StageController struct{
	beego.Controller
}

func (this *StageController) Get(){
	this.Data["IsStage"] = true
	this.TplName="stage.html"
}
