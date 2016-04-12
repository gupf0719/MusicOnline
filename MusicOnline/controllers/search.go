package controllers

import (
	"github.com/astaxie/beego"
)

type SearchController struct {
	beego.Controller
}

func (this *SearchController) Get() {
	this.TplName = "search.html"
}
