package routers

import (
	"MusicOnline/MusicOnline/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/message", &controllers.MessageController{})
	beego.Router("/news", &controllers.NewsController{})
	beego.Router("/search", &controllers.SearchController{})
	beego.Router("/stage", &controllers.StageController{})
	beego.Router("/regis", &controllers.RegisterController{})
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/adminnews", &controllers.AdminnewsController{})
	beego.Router("/adminstage", &controllers.AdminstageController{})
	beego.Router("/adminmessage", &controllers.AdminmessageController{})
	beego.Router("/adminmusic", &controllers.AdminmusicController{})
	beego.Router("/adminmv", &controllers.AdminmvController{})
	beego.AutoRouter(&controllers.AdminnewsController{})
	beego.AutoRouter(&controllers.AdminstageController{})
	beego.AutoRouter(&controllers.AdminmessageController{})
	beego.AutoRouter(&controllers.NewsController{})
	beego.AutoRouter(&controllers.StageController{})
	beego.AutoRouter(&controllers.MessageController{})
	beego.AutoRouter(&controllers.AdminmusicController{})
	beego.AutoRouter(&controllers.AdminmvController{})
}

