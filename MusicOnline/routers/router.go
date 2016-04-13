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
}

