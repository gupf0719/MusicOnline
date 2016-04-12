package models

import (
	"github.com/astaxie/beego/orm"
)

type Message struct {
	Id       int64  `orm:"auto"`
	Nickname string `orm:"size(128)"`
	Content  string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Message))
}
