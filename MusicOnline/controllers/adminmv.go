package controllers

import (
	"MusicOnline/MusicOnline/models"

	"github.com/astaxie/beego"
)

type AdminmvController struct {
	beego.Controller
}

func (this *AdminmvController) Get() {
	this.Data["IsMv"] = true
	this.TplName = "adminmv.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	music, err := models.GetAllMusic(false)
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Music"] = music
	}
}

func (this *AdminmvController) Post() {
	//	if !checkAccount(this.Ctx) {
	//		this.Redirect("/login", 302)
	//		return
	//	}
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

func (this *AdminmvController) Add() {
	this.TplName = "addmusic.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
}

func (this *AdminmvController) Modify() {
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

func (this *AdminmvController) Delete() {
	err := models.DeleteMusic(this.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/adminmusic", 302)
}
