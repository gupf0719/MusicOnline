package controllers

import (
	"MusicOnline/MusicOnline/models"

	"github.com/astaxie/beego"
)

type AdminmusicController struct {
	beego.Controller
}

func (this *AdminmusicController) Get() {
	this.Data["IsMusic"] = true
	this.TplName = "adminmusic.html"

	music, err := models.GetAllMusic(false)
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Music"] = music
	}
}

func (this *AdminmusicController) Post() {

	tid := this.Input().Get("tid")
	name := this.Input().Get("name")
	singer := this.Input().Get("singer")
	special := this.Input().Get("special")
	category := this.Input().Get("category")

	var err error
	if len(tid) == 0 {
		err = models.AddMusic(name, singer, special, category)
	} else {
		err = models.ModifyMusic(tid, name, singer, special, category)
	}

	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/adminmusic", 302)
}

func (this *AdminmusicController) Add() {
	this.TplName = "addmusic.html"

}

func (this *AdminmusicController) Modify() {
	this.TplName = "modifymusic.html"
	tid := this.Input().Get("tid")

	music, err := models.GetMusic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/adminmusic", 302)
		return
	}

	this.Data["Music"] = music
	this.Data["Tid"] = tid
}

func (this *AdminmusicController) Delete() {
	err := models.DeleteMusic(this.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/adminmusic", 302)
}
