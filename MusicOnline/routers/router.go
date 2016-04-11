package routers

import (
	"MusicOnline/MusicOnline/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
