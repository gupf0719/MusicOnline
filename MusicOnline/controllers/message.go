package controllers

import (
	"fmt"
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

	for _, message := range messages {
		fmt.Println("msgcon"+message.Content)
		// 解析 markdown 格式为 html
		contentMark := []byte(message.Content)
		htmlBytes := ParseDoc(contentMark)
		message.Content = string(htmlBytes)
		fmt.Println("msgcon"+message.Content)
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
