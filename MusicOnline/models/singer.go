package models

import (
	"github.com/astaxie/beego/orm"
)

type Singer struct {
	Id       int64  `orm:"auto"`
	Name     string `orm:"size(128)"`
	Category string `orm:"size(128)"`
	Image    string `orm:"size(128)"`
	Location string `orm:"size(128)"`
	Sex      string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Singer))
}
