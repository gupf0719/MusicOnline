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
