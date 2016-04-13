package models

import (
	"github.com/astaxie/beego/orm"
)

type Special struct {
	Id     int64  `orm:"auto"`
	Name   string `orm:"size(128)"`
	Singer string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Special))
}

func AddSpecial(m *Special) error {
	o := orm.NewOrm()
	_, err := o.Insert(m)
	if err != nil {
		return err
	}
	return nil
}

func ModifySpecial(m *Special) error {
	o := orm.NewOrm()

	special := new(Special)

	qs := o.QueryTable("special")
	err := qs.Filter("id", m.Id).One(special)
	if err != nil {
		return err
	}

	special.Name = m.Name
	special.Singer = m.Singer

	_, err = o.Update(special)
	if err != nil {
		return err
	}
	return nil
}

func DelSpecial(id int64) error {
	o := orm.NewOrm()

	_, err := o.Delete(&Special{Id: id})
	if err != nil {
		return err
	}
	return nil
}
