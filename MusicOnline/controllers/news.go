package controllers

import (
	"github.com/astaxie/beego"
)

type NewsController struct {
	beego.Controller
}

func (this *NewsController) Get() {
	this.Data["IsNews"] = true
	this.TplName = "news.html"
}
