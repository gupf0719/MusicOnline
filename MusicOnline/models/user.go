package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int64  `orm:"auto"`
	Name     string `orm:"size(128)"`
	Password string `orm:"size(128)"`
	NickName string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func AddUser(m *User) error {

	o := orm.NewOrm()

	_, err := o.Insert(m)
	if err != nil {
		return err
	}

	return err
}
