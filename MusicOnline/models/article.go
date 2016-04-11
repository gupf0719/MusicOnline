package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id      int64     `orm:"auto"`
	Name    string    `orm:"size(128)"`
	Content string    `orm:"size(128)"`
	Time    time.Time `orm:"type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Article))

}
