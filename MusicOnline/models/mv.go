package models

import (
	"github.com/astaxie/beego/orm"
)

type Mv struct {
	Id    int64  `orm:"auto"`
	Name  string `orm:"size(128)"`
	Image string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Mv))
}
