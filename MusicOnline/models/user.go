package models

import (
	"strings"

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

	user := new(User)
	qs := o.QueryTable("user")
	err := qs.Filter("name", m.Name).One(user)
	if err != nil {
		_, err := o.Insert(m)
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}

func IsExist(uname string, password string) bool {
	o := orm.NewOrm()

	user := new(User)
	qs := o.QueryTable("user")
	err := qs.Filter("name", uname).One(user)
	if err != nil {
		error.Error(err)
	}

	var exist bool
	if strings.EqualFold(password, user.Password) {
		exist = true
		return exist
	} else {
		exist = false
		return exist
	}
}
