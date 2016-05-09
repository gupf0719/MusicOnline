package models

import (
	"strconv"
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

func DelMusicMessage(mid string) error {
	midNum, err := strconv.ParseInt(mid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	_, err = o.Delete(&Message{Id: midNum})
	if err != nil {
		return err
	}
	return nil
}

func GetAllMusicMessages() (messages []*Message, err error) {
	messages = make([]*Message, 0)

	o := orm.NewOrm()
	qs := o.QueryTable("message")
	_, err = qs.All(&messages)
	return messages, err
}
