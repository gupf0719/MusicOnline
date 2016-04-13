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

func AddSinger(m *Singer) error {
	o := orm.NewOrm()
	_, err := o.Insert(m)
	if err != nil {
		return err
	}
	return nil
}

func ModifySinger(m *Singer) error {
	o := orm.NewOrm()

	singer := new(Singer)

	qs :=o.QueryTable("singer")
	err :=qs.Filter("id",m.Id).One(singer)
	if err != nil {
			return err
	}

	singer.Name=m.Name
	singer.Category=m.Category
	singer.Image=m.Image
	singer.Location=m.Location
	singer.Sex=singer.Sex

	_, err = o.Update(singer)
	if err != nil {
			return err
	}
	return nil
}

func DelSinger(id int64) error {
	o := orm.NewOrm()

	_, err := o.Delete(&Singer{Id: id})
	if err != nil {
		return err
	}
	return nil
}
