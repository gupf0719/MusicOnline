package controllers

import (
	"MusicOnline/MusicOnline/models"

	"github.com/astaxie/beego"
)

type StageController struct {
	beego.Controller
}

func (this *StageController) Get() {
	this.Data["IsStage"] = true
	this.TplName = "stage.html"

	news, err := models.GetAllStage(false)
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["News"] = news
	}
}

func (this *StageController) Post() {

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
	this.Redirect("/stage", 302)
}

func (this *StageController) View() {
	this.TplName = "showstage.html"

	stage, err := models.GetStage(this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	// 解析 markdown 格式为 html
	contentMark := []byte(stage.Content)
	htmlBytes := ParseDoc(contentMark)
	stage.Content = string(htmlBytes)

	this.Data["Stage"] = stage
	this.Data["Tid"] = this.Ctx.Input.Param("0")
}
