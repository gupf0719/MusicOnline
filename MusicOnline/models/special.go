package models

import (
	"github.com/astaxie/beego/orm"
)

type Special struct {
	Id       int64  `orm:"auto"`
	Name     string `orm:"size(128)"`
	SingerId int64
}

func init() {
	orm.RegisterModel(new(Special))
}
