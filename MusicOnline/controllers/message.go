package controllers

import (
	"github.com/astaxie/beego"
)

type MessageController struct {
	beego.Controller
}

func (this *MessageController) Get() {
	this.Data["IsMessage"] = true
	this.TplName = "message.html"
}
