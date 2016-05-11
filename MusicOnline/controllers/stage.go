package controllers

import (
	"MusicOnline/MusicOnline/models"

	"github.com/astaxie/beego"
	"log"
)

type StageController struct {
	beego.Controller
}

func (this *StageController) Get() {
	this.Data["IsStage"] = true
	this.TplName = "stage.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	news, err := models.GetAllStage(false)
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["News"] = news
	}
}

func (this *StageController) Post() {
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

func (this *StageController) View() {
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
