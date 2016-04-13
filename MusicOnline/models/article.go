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

func AddArticle(m *Article) error {
	o := orm.NewOrm()
	_, err := o.Insert(m)
	if err != nil {
		return err
	}
	return nil
}

func ModifyArticle(m *Article) error {
	o := orm.NewOrm()

	article := new(Article)

	qs :=o.QueryTable("article")
	err :=qs.Filter("id",m.Id).One(article)
	if err != nil {
			return err
	}

	article.Name=m.Name
	article.Content=m.Content
	article.Time=m.Time

	_, err = o.Update(article)
	if err != nil {
			return err
	}
	return nil
}

func DelArticle(id int64) error {
	o := orm.NewOrm()

	_, err := o.Delete(&Article{Id: id})
	if err != nil {
		return err
	}
	return nil
}


