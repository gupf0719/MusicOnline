package models

import (
	"github.com/astaxie/beego/orm"
)

type Mv struct {
	Id           int64  `orm:"auto"`
	Name         string `orm:"size(128)"`
	Introduction string `orm:"size(128)"`
	Image        string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Mv))
}

func GetMv(name string)(*Mv, error){

	o := orm.NewOrm()
	mv := new(Mv)

	qs := o.QueryTable("mv")
	err := qs.Filter("name", name).One(mv)
	if err != nil {
		return nil, err
	}

	return mv, err
}

func AddMv(m *Mv) error {
	o := orm.NewOrm()
	_, err := o.Insert(m)
	if err != nil {
		return err
	}
	return nil
}

func ModifyMv(m *Singer) error {
	o := orm.NewOrm()

	mv := new(Singer)

	qs := o.QueryTable("mv")
	err := qs.Filter("id", m.Id).One(mv)
	if err != nil {
		return err
	}

	mv.Name = m.Name
	mv.Image = m.Image

	_, err = o.Update(mv)
	if err != nil {
		return err
	}
	return nil
}

func DelMv(id int64) error {
	o := orm.NewOrm()

	_, err := o.Delete(&Mv{Id: id})
	if err != nil {
		return err
	}
	return nil
}
