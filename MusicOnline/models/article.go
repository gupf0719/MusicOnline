package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id      int64     `orm:"auto"`
	Title   string    `orm:"size(128)"`
	Content string    `orm:"size(128)"`
	Time    time.Time `orm:"type(datetime)"`
	Views   int64     `orm:"index"`
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

	qs := o.QueryTable("article")
	err := qs.Filter("id", m.Id).One(article)
	if err != nil {
		return err
	}

	article.Title = m.Title
	article.Content = m.Content
	article.Time = m.Time

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
