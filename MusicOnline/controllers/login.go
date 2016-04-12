package controllers

import (
	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get(){
	this.TplName="login.html"
}

func (this *LoginController) Post(){

}


