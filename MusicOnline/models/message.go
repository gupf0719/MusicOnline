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

func AddMusicMessage(m *Message) error {
	o := orm.NewOrm()
	_, err := o.Insert(m)
	if err != nil {
		return err
	}
	return nil
}

func DelMusicMessage(mid int64) error {
	o := orm.NewOrm()

	_, err := o.Delete(&Message{Id: mid})
	if err != nil {
		return err
	}
	return nil
}
