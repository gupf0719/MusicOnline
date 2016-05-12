package controllers

import (
	"MusicOnline/MusicOnline/models"

	"github.com/astaxie/beego"
)

type NewsController struct {
	beego.Controller
}

func (this *NewsController) Get() {
	this.Data["IsNews"] = true
	this.TplName = "news.html"

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

func (this *NewsController) Post() {
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
	this.Redirect("/news", 302)
}

func (this *NewsController) View() {
	this.TplName = "shownews.html"

	news, err := models.GetNews(this.Ctx.Input.Param("0"))

	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["News"] = news
	this.Data["Tid"] = this.Ctx.Input.Param("0")
}
