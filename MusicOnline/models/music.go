package models

import (
	"github.com/astaxie/beego/orm"
)

type Music struct {
	Id        int64  `orm:"auto"`
	Name      string `orm:"size(128)"`
	Singer    string `orm:"size(128)"`
	SingerId  int64
	Special   string `orm:"size(128)"`
	SpecialId int64
	Category  string `orm:"size(128)"`
	Image     string `orm:"size(128)"`
	Uid       int64
}

func init() {
	orm.RegisterModel(new(Music))
}
