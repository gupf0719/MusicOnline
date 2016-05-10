package controllers

import (
	"MusicOnline/MusicOnline/models"

	"github.com/astaxie/beego"
)

type AdminstageController struct {
	beego.Controller
}

func (this *AdminstageController) Get() {
	this.Data["IsStage"] = true
	this.TplName = "adminstage.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	stage, err := models.GetAllStage(false)
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Stage"] = stage
	}
}

func (this *AdminstageController) Post() {
	//	if !checkAccount(this.Ctx) {
	//		this.Redirect("/login", 302)
	//		return
	//	}
	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")

	var err error
	if len(tid) == 0 {
		err = models.AddStage(title, content)
	} else {
		err = models.ModifyStage(tid, title, content)
	}

	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/adminstage", 302)
}

func (this *AdminstageController) Add() {
	this.TplName = "addstage.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
}

func (this *AdminstageController) View() {
	this.TplName = "viewstage.html"

	stage, err := models.GetStage(this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		this.Redirect("/admin", 302)
		return
	}

	this.Data["Stage"] = stage
	this.Data["Tid"] = this.Ctx.Input.Param("0")
}

func (this *AdminstageController) Modify() {
	this.TplName = "modifystage.html"
	tid := this.Input().Get("tid")

	stage, err := models.GetStage2(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/admin", 302)
		return
	}

	this.Data["Stage"] = stage
	this.Data["Tid"] = tid
}

func (this *AdminstageController) Delete() {
	//	if !checkAccount(this.Ctx) {
	//		this.Redirect("/login", 302)
	//		return
	//	}

	err := models.DeleteStage(this.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/adminstage", 302)
}
