package controllers

import (
	"MusicOnline/MusicOnline/models"

	"github.com/astaxie/beego"
)

type MessageController struct {
	beego.Controller
}

func (this *MessageController) Get() {
	this.Data["IsMessage"] = true
	this.TplName = "message.html"

	messages, err := models.GetAllMusicMessages()
	if err != nil {
		beego.Error(err)
		return
	}

	this.Data["Messages"] = messages
	this.Data["IsLogin"] = checkAccount(this.Ctx)
}

//func (this *MessageController) Post() {

//}

func (this *MessageController) Add() {
	err := models.AddMusicMessage(this.Input().Get("nickname"), this.Input().Get("content"))
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/message", 302)
}
