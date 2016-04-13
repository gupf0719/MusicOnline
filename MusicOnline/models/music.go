package models

import (
	"github.com/astaxie/beego/orm"
)

type Music struct {
	Id        int64  `orm:"auto"`
	Name      string `orm:"size(128)"`
	Singer    string `orm:"size(128)"`
	Special   string `orm:"size(128)"`
	Category  string `orm:"size(128)"`
	Image     string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Music))
}

func AddMusic(m *Music) error {
	o := orm.NewOrm()
	_, err := o.Insert(m)
	if err != nil {
		return err
	}
	return nil
}

func ModifyMusic(m *Music) error {
	o := orm.NewOrm()

	music := new(Music)

	qs := o.QueryTable("music")
	err := qs.Filter("id", m.Id).One(music)
	if err != nil {
		return err
	}

	music.Name = m.Name
	music.Singer = m.Singer
	music.Special = m.Special
	music.Category = m.Category
	music.Image = m.Image

	_, err = o.Update(music)
	if err != nil {
		return err
	}
	return nil
}

func DelMusic(id int64) error {
	o := orm.NewOrm()

	_, err := o.Delete(&Music{Id: id})
	if err != nil {
		return err
	}
	return nil
}
