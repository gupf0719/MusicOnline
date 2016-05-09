package models

import (
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id   int64
	Title string
}

func init() {
	orm.RegisterModel(new(Category))
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	cates := make([]*Category, 0)

	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}
