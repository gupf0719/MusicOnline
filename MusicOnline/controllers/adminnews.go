package controllers

import (
	"MusicOnline/MusicOnline/models"

	"github.com/astaxie/beego"
)

type AdminnewsController struct {
	beego.Controller
}

func (this *AdminnewsController) Get() {
	this.Data["IsNews"] = true
	this.TplName = "adminnews.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	news, err := models.GetAllNews(this.Input().Get("cate"), false)
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["News"] = news
	}

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Categories"] = categories
	}
}

func (this *AdminnewsController) Post() {
	//	if !checkAccount(this.Ctx) {
	//		this.Redirect("/login", 302)
	//		return
	//	}
	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	category := this.Input().Get("category")

	var err error
	if len(tid) == 0 {
		err = models.AddNews(title, content, category)
	} else {
		err = models.ModifyNews(tid, title, content, category)
	}

	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/adminnews", 302)
}

func (this *AdminnewsController) Add() {
	this.TplName = "addnews.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
}

func (this *AdminnewsController) View() {
	this.TplName = "viewnews.html"

	news, err := models.GetNews(this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		this.Redirect("/admin", 302)
		return
	}

	this.Data["News"] = news
	this.Data["Tid"] = this.Ctx.Input.Param("0")
}

func (this *AdminnewsController) Modify() {
	this.TplName = "modifynews.html"
	tid := this.Input().Get("tid")

	news, err := models.GetNews2(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/admin", 302)
		return
	}

	this.Data["News"] = news
	this.Data["Tid"] = tid
}

func (this *AdminnewsController) Delete() {
	//	if !checkAccount(this.Ctx) {
	//		this.Redirect("/login", 302)
	//		return
	//	}

	err := models.DeleteNews(this.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/adminnews", 302)
}
