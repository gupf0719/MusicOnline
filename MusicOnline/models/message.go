package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Message struct {
	Id       int64     `orm:"auto"`
	Nickname string    `orm:"size(128)"`
	Content  string    `orm:"size(1024)"`
	Time     time.Time `orm:"index"`
}

func init() {
	orm.RegisterModel(new(Message))
}

func AddMusicMessage(nickname string, content string) error {

	message := &Message{

		Nickname: nickname,
		Content:  content,
		Time:     time.Now(),
	}

	o := orm.NewOrm()
	_, err := o.Insert(message)

	return err
}

func DelMusicMessage(mid int64) error {
	o := orm.NewOrm()

	_, err := o.Delete(&Message{Id: mid})
	if err != nil {
		return err
	}
	return nil
}
