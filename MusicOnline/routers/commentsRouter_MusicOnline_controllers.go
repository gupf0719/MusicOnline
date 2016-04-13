package routers

import (
	"github.com/astaxie/beego"
)

func init(){
	beego.GlobalControllerRouter["MusicOnline/MusicOnline/controllers:RegisterController"] = append(beego.GlobalControllerRouter["MusicOnline/MusicOnline/controllers:RegisterController"],
	beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["MusicOnline/MusicOnline/controllers:RegisterController"] = append(beego.GlobalControllerRouter["MusicOnline/MusicOnline/controllers:RegisterController"],
	beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})
}
