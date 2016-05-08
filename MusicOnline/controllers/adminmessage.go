package controllers

import (
	"github.com/astaxie/beego"
)

type AdminmessageController struct {
	beego.Controller
}

func (this *AdminmessageController) Get(){
	this.Data["IsMessage"] = true
	this.TplName = "adminmessage.html"
}
